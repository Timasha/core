package kafka

import (
	"context"
	"errors"

	"github.com/IBM/sarama"
	"github.com/Timasha/core/kafka/domain"
	"github.com/Timasha/core/kafka/errlist"
	"github.com/Timasha/core/log"
)

type Client struct {
	cfg Config

	sp sarama.SyncProducer
	cg sarama.ConsumerGroup

	eventMap domain.EventMap
}

func (c *Client) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Client) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Client) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				return nil
			}

			err := c.MsgHandler(session, msg)
			if err != nil {
				log.Errorf("error handling kafka message: %v", err)
			}
		}
	}
}

func (c *Client) MsgHandler(
	session sarama.ConsumerGroupSession,
	msg *sarama.ConsumerMessage,
) (err error) {
	defer session.MarkMessage(msg, "")

	eventType := ""
	for _, header := range msg.Headers {
		if string(header.Key) == EventTypeKey {
			eventType = string(header.Value)
		}
	}

	if eventType == "" {
		return errlist.ErrNoEventType
	}

	ctx := context.Background()
	event := domain.Event{
		Key:  string(msg.Key),
		Data: msg.Value,
	}

	topicEventMap, ok := c.eventMap[msg.Topic]
	if !ok {
		return errlist.ErrNoTopicHandlers
	}

	handler, ok := topicEventMap[eventType]
	if !ok {
		return errlist.ErrNoEventTypeHandler
	}

	err = handler(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Start(ctx context.Context) (err error) {
	config := sarama.NewConfig()

	config.Version, err = sarama.ParseKafkaVersion(c.cfg.Version)
	if err != nil {
		return err
	}

	config.Producer.Partitioner = func(topic string) sarama.Partitioner {
		return sarama.NewHashPartitioner(topic)
	}

	switch c.cfg.Assignor {
	case StickyAssignor:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case RoundRobinAssignor:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case RangeAssignor:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		return errors.New("invalid rebalance strategy")
	}

	if c.cfg.Oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	c.sp, err = sarama.NewSyncProducer(c.cfg.Brokers, config)
	if err != nil {
		log.Fatalf("Ошибка при создании producer: %v", err)
	}

	consumerGroup, err := sarama.NewConsumerGroup(c.cfg.Brokers, c.cfg.ConsumerGroup, config)
	if err != nil {
		return err
	}

	go c.ConsumeFunc(context.Background(), consumerGroup)

	return nil
}

func (c *Client) ConsumeFunc(ctx context.Context, consumerGroup sarama.ConsumerGroup) {
	if c.cfg.EnableConsumer {
		for {
			err := consumerGroup.Consume(ctx, c.cfg.Topics, c)
			if err != nil {
				log.Fatalf("Ошибка при Consume: %v", err)
			}
		}
	}
}

func (c *Client) Stop(ctx context.Context) (err error) {
	err = c.sp.Close()
	if err != nil {
		return err
	}

	err = c.cg.Close()
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetName() string {
	return "kafka client"
}

func (c *Client) IsEnabled() bool {
	return c.cfg.IsEnabled
}

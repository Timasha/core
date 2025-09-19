package kafka

import (
	"encoding/json"

	"github.com/Timasha/core/components"
	"github.com/Timasha/core/kafka/errlist"
)

type Assignor string

const (
	StickyAssignor     Assignor = "sticky"
	RoundRobinAssignor Assignor = "roundrobin"
	RangeAssignor      Assignor = "range"
)

type Config struct {
	components.Config

	Brokers []string
	Topics  []string

	ConsumerGroup  string
	Version        string
	Assignor       Assignor
	EnableConsumer bool `default:"true"`
	Oldest         bool `default:"true"`
}

func (c *Config) UnmarshalJSON(bytes []byte) (err error) {
	type Alias Config
	aux := struct {
		Alias
	}{}

	err = json.Unmarshal(bytes, &aux)
	if err != nil {
		return err
	}

	*c = Config(aux.Alias)

	if c.Assignor != StickyAssignor &&
		c.Assignor != RoundRobinAssignor &&
		c.Assignor != RangeAssignor {
		return errlist.ErrInvalidAssignor
	}

	return nil
}

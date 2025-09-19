package domain

import "context"

// EventMap - key: topic, value: event handler map
type EventMap map[string]TopicEventMap

// TopicEventMap - key: eventType, value: event handler
type TopicEventMap map[string]func(ctx context.Context, event Event) error

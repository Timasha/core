package errlist

import "errors"

var (
	ErrNoTopicHandlers    = errors.New("no topic handlers")
	ErrNoEventTypeHandler = errors.New("no event type handler")
	ErrNoEventType        = errors.New("no event type")
	ErrInvalidAssignor    = errors.New("invalid assignor")
)

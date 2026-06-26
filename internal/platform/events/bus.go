package events

import (
	"context"
)

type Bus interface {
	Publish(ctx context.Context, event Event) error
	Subscribe(eventName string, handler Handler)
}

type Handler func(context.Context, Event) error

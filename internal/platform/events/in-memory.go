package events

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
)

type InMemoryBus struct {
	mu       sync.Mutex
	logger   *slog.Logger
	handlers map[string][]Handler
}

func NewInMemoryBus(logger *slog.Logger) *InMemoryBus {
	return &InMemoryBus{
		logger:   logger,
		handlers: make(map[string][]Handler),
		mu:       sync.Mutex{},
	}
}

func (b *InMemoryBus) Publish(ctx context.Context, event Event) error {

	b.mu.Lock()
	handlers := b.handlers[event.EventName()]
	b.mu.Unlock()

	for _, h := range handlers {
		err := h(ctx, event)
		if err != nil {
			errorMsg := fmt.Sprintf("Error publishing event %v: %v", event.EventName(), err)
			b.logger.Error(fmt.Sprintf(errorMsg))
			return fmt.Errorf(errorMsg)
		}
	}

	b.logger.Info(fmt.Sprintf("Event published %v", event.EventName()))
	return nil
}

func (b *InMemoryBus) Subscribe(eventName string, handler Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[eventName] = append(b.handlers[eventName], handler)
}

package workerspool

import (
	"context"
	"sync"
)

type Job[T any] struct {
	ID   string
	Data T
}

type WorkerPool[T any] struct {
	maxWorkers int
	Workers    map[string]func()

	ctx    context.Context
	cancel context.CancelFunc

	Queue chan Job[T]
	mutex sync.Mutex
	wg    sync.WaitGroup

	idGen IdGenerator
}

func NewWorkerPool[T any](maxWorkers int) *WorkerPool[T] {

	ctx, cancel := context.WithCancel(context.Background())
	idGen := NewGoogleIdGenerator()

	return &WorkerPool[T]{
		maxWorkers: maxWorkers,
		Workers:    make(map[string]func()),
		ctx:        ctx,
		cancel:     cancel,
		Queue:      make(chan Job[T]),
		mutex:      sync.Mutex{},
		wg:         sync.WaitGroup{},
		idGen:      idGen,
	}
}

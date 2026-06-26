package workerspool

import (
	"context"
	"fmt"
	"sync"
)

type Job[T any] struct {
	ID   string
	Data T
}

type HandleFn[T any] func(
	ctx context.Context,
	workerId string,
	data T,
) error

type WorkerPool[T any] struct {
	maxWorkers int
	Workers    map[string]context.CancelFunc

	ctx    context.Context
	cancel context.CancelFunc

	Queue chan Job[T]
	mutex sync.Mutex
	wg    sync.WaitGroup

	idGen IdGenerator

	HandleFn HandleFn[T]
}

func NewWorkerPool[T any](maxWorkers, queueSize int) *WorkerPool[T] {

	ctx, cancel := context.WithCancel(context.Background())
	idGen := NewGoogleIdGenerator()

	return &WorkerPool[T]{
		maxWorkers: maxWorkers,
		Workers:    make(map[string]context.CancelFunc),
		ctx:        ctx,
		cancel:     cancel,
		Queue:      make(chan Job[T], queueSize),
		mutex:      sync.Mutex{},
		wg:         sync.WaitGroup{},
		idGen:      idGen,
	}
}

func (w *WorkerPool[T]) Start(fn HandleFn[T]) error {
	w.HandleFn = fn
	for i := 0; i < w.maxWorkers; i++ {
		id := w.idGen.Generate()
		w.startWorker(id)
	}

	return nil
}

func (w *WorkerPool[T]) Stop() {
	w.cancel()
	w.wg.Wait()
}

func (w *WorkerPool[T]) Submit(job Job[T]) error {
	select {
	case w.Queue <- job:
		return nil
	case <-w.ctx.Done():
		return fmt.Errorf("worker pool stopped")
	}
}

func (w *WorkerPool[T]) startWorker(id string) {

	ctx, cancel := context.WithCancel(context.Background())

	w.mutex.Lock()
	w.Workers[id] = cancel
	w.mutex.Unlock()

	w.wg.Add(1)

	go func() {

		defer w.wg.Done()

		defer func() {

			if r := recover(); r != nil {
				fmt.Printf("Worker %s crashed. Reason: %v\n", id, r)
			}

			w.mutex.Lock()
			delete(w.Workers, id)
			w.mutex.Unlock()

			if w.ctx.Err() != nil {
				fmt.Printf("Restarting worker %s.\n", id)
				w.startWorker(id)
			}
		}()

		w.runWorker(ctx, id)

	}()
}

func (w *WorkerPool[T]) runWorker(ctx context.Context, id string) {

	for {
		select {
		case <-w.ctx.Done():
			return
		case job := <-w.Queue:
			err := w.HandleFn(ctx, id, job.Data)
			if err != nil {
				return
			}
		}
	}

}

func (w *WorkerPool[T]) stopWorker(id string) {
	w.mutex.Lock()
	cancel, ok := w.Workers[id]
	w.mutex.Unlock()

	if ok {
		cancel()
	}
}

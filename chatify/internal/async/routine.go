package async

import (
	"context"
	"sync"
	"time"
)

func RunEndlessRoutineWithCancel(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc, work func() error, delay time.Duration) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return // Terminate the goroutine if the context is canceled
		default: // Default case prevents blocking the select
			// Wait a second before ending
			<-time.After(delay)
		}

		if err := work(); err != nil {
			cancel()
		}
	}
}

func RunWorker(dataCh chan []byte, wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc, work func([]byte) error) {
	defer wg.Done()

	for {
		select {
		case data := <-dataCh:
			if err := work(data); err != nil {
				cancel()
			}
		case <-ctx.Done():
			return
		}
	}
}

func RunRoutineWithCancel(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc, work func() error) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return // Terminate the goroutine if the context is canceled
	default: // Default case prevents blocking the select
	}

	if err := work(); err != nil {
		cancel()
	}
}

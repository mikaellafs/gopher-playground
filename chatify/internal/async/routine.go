package async

import (
	"context"
	"sync"
)

func RunEndlessRoutineWithCancel(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc, work func() error) {
	defer wg.Done()

	for {
		RunRoutineWithCancel(ctx, cancel, work)
	}
}

func RunRoutineWithCancel(ctx context.Context, cancel context.CancelFunc, work func() error) {
	select {
	case <-ctx.Done():
		return // Terminate the goroutine if the context is canceled
	default: // Default case prevents blocking the select
	}

	if err := work(); err != nil {
		cancel()
	}
}

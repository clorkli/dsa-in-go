package concurrency

import (
	"context"
	"testing"
	"time"
)

func TestProducerConsumer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	dataChan := make(chan int, 5)

	done := make(chan struct{})

	go func() {
		for {
			select {
				case <-ctx.Done():
					close(done)
					return
				case val, ok := <-dataChan:
					if !ok {
						close(done)
						return
					}
					t.Logf("Consumed: %d", val)
			}
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			dataChan <- i
		}
		close(dataChan)
	}()

	select {
		case <-done:
			t.Log("Test completed")
		case <-ctx.Done():
			t.Fatal("Test timed out! Possible deadlock.")
	}
}
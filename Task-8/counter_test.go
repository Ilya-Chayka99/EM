package Task_8

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}
	count := 0
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go counter(mx, &count, i, wg)
	}
	wg.Wait()
	if count != 20 {
		t.Fatalf("expected 20, got %d", count)
	}
}

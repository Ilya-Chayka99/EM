package Task_8

import "sync"

func counter(mx *sync.Mutex, counter *int, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	mx.Lock()
	defer mx.Unlock()
	if num%2 == 0 {
		*counter += num
	}
}

package script

import (
	"fmt"
	"sync"
)

type counter struct {
	lock sync.RWMutex
	val  int
}

func (c *counter) Add(int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.val++
}

func (c *counter) Value() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.val
}
func MutexMain() {
	var wg sync.WaitGroup
	var meter counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}

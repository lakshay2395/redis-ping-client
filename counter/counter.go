package counter

import "sync"

type Counter interface {
	Get() int
	Increment()
}

type counter struct {
	number int
	mutex  sync.Mutex
}

func New() *counter {
	return &counter{number: 0}
}

func (c *counter) Increment() {
	c.mutex.Lock()
	c.number += 1
	c.mutex.Unlock()
}

func (c *counter) Get() int {
	c.mutex.Lock()
	value := c.number
	c.mutex.Unlock()
	return value
}

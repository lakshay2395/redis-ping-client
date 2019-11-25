package counter

import (
	"github.com/lakshay2395/redis-ping-client/testutils"
	"testing"
)

func TestNew(t *testing.T) {
	c := New()
	value := c.Get()
	expectedValue := 0
	testutils.Compare(t, value, expectedValue)
}

func TestIncrement(t *testing.T) {
	c := New()
	c.Increment()
	c.Increment()
	expectedValue := 2
	value := c.Get()
	testutils.Compare(t, value, expectedValue)
}

func TestIncrementWithConcurrency(t *testing.T) {
	c := New()
	ch := make(chan int, 3)
	go concurrentIncrement(c, ch)
	go concurrentIncrement(c, ch)
	go concurrentIncrement(c, ch)
	for i := 0; i < 3; i++ {
		<-ch
	}
	expectedValue := 3
	value := c.Get()
	testutils.Compare(t, value, expectedValue)
}

func concurrentIncrement(c *counter, ch chan int) {
	c.Increment()
	ch <- 1
}

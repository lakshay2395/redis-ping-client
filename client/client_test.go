package client

import (
	"reflect"
	"testing"

	"github.com/lakshay2395/redis-ping-client/param"
	"github.com/lakshay2395/redis-ping-client/testutils"
)

func TestNew(t *testing.T) {
	p := param.New("h", "k", "p", 1, 1)
	c := New(p)
	testutils.CheckForNil(t, c.parameters)
	testutils.CheckForNil(t, c.redisClient)
	testutils.CheckForNil(t, c.counter)
}

func TestGenerateBatches(t *testing.T) {
	requestCount := 100
	batchSize := 10
	batches, err := generateBatches(requestCount, batchSize)
	if err != nil {
		t.Error("Error Occurred : ", err)
	}
	expectedValue := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	if !reflect.DeepEqual(batches, expectedValue) {
		t.Errorf("Expected %v, got %v", batches, expectedValue)
	}
}

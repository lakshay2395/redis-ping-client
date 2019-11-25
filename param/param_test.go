package param

import (
	"testing"

	"github.com/lakshay2395/redis-ping-client/testutils"
)

func TestNew(t *testing.T) {
	p := New("h", "k", "p", 1, 1)
	if p == (params{}) {
		t.Error("Incorrect value returned from new function")
	}
}

func TestHost(t *testing.T) {
	p := New("h", "k", "p", 1, 1)
	expectedValue := "h"
	value := p.Host()
	testutils.Compare(t, value, expectedValue)
}

func TestKey(t *testing.T) {
	p := New("h", "k", "p", 1, 1)
	expectedValue := "k"
	value := p.Key()
	testutils.Compare(t, value, expectedValue)
}

func TestPort(t *testing.T) {
	p := New("h", "k", "p", 1, 1)
	expectedValue := "p"
	value := p.Port()
	testutils.Compare(t, value, expectedValue)
}

func TestRequestCount(t *testing.T) {
	p := New("h", "k", "p", 1, 1)
	expectedValue := 1
	value := p.RequestCount()
	testutils.Compare(t, value, expectedValue)
}

func TestBatchSize(t *testing.T) {
	p := New("h", "k", "p", 1, 1)
	expectedValue := 1
	value := p.BatchSize()
	testutils.Compare(t, value, expectedValue)
}

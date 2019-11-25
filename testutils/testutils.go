package testutils

import "testing"

func Compare(t *testing.T, value interface{}, expectedValue interface{}) {
	if value != expectedValue {
		t.Errorf(t.Name()+" : Expected value is %v, got %v", expectedValue, value)
	}
}

func CheckForNil(t *testing.T, value interface{}) {
	if value == nil {
		t.Errorf(t.Name() + ": Found nil")
	}
}

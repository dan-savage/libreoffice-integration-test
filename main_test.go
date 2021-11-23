package main

import (
	"testing"
)

func TestReturnOne(t *testing.T) {
	res := ReturnOne()
	if res != 1 {
		t.Errorf("failed to return 1: %d", res)
	} else {
		t.Logf("test succeeded")
	}
}

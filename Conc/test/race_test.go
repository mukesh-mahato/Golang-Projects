package main

import (
	"sync/atomic"
	"testing"
)

func TestDataRaceCondition(t *testing.T) {
	var state int32
	// var mu sync.RWMutex
	for i := 0; i < 10; i++ {
		go func(i int) {
			// mu.Lock()
			// state += int32(i)
			atomic.AddInt32(&state, int32(i))
			// mu.Unlock()
		}(i)
	}
}

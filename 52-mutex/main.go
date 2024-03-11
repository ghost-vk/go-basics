// ASSIGNMENT
//
// We send emails across many different goroutines at Mailio. To
// keep track of how many we've sent to a given email address, we use an
// in-memory map.
//
// Our safeCounter struct is unsafe! Update the inc() and val() methods so
// that they utilize the safeCounter's mutex to ensure that the map is not
// accessed by multiple goroutines at the same time.

package main

import (
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	// pointer to mutex to sharing across goroutines
	mu *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.slowVal(key)
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}

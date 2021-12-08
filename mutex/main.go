package main

import (
	"strconv"
	"sync"
)

var myBalance = &balance{amount: 50.00, currency: "USD"}


// Mutex is only being used in the context of our balance object, we can embed it
// This will keep locking rules nice and clear if our codebase grows large
type balance struct {
	amount 		float64
	currency 	string
	// A reader/writer mutual exclusion lock which allows any number of readers
	// to hold the lock or one writer
	// more efficient than using a full mutex in situations where we have a 
	// high ratios of reads to writes
	mu  		sync.RWMutex 
}

func (b *balance) Add(i float64) {
	// Possible Race Condition without mutex
	b.mu.Lock()
	b.amount += i
	b.mu.Unlock()
}

func (b *balance) Display() string {
	// Possible Race Condition without mutex
	b.mu.RLock()
	// ensure the mutex gets unlocked before executing the return statement
	defer b.mu.RUnlock()
	return strconv.FormatFloat(b.amount, 'f', 2, 64) + " " + b.currency
}
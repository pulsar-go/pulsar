package queue

import "github.com/panjf2000/ants"

// Pool defines the default queue pool
var Pool *ants.Pool

// NewPool creates the default queue pool
func NewPool(number int) {
	// Create that pool
	pool, _ := ants.NewPool(number)
	Pool = pool
}

// Dispatch runs the handler in a goroutine
func Dispatch(handler func()) error {
	return Pool.Submit(handler)
}

package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, threshold int) Circuit {
	var failures int
	var last = time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		m.RLock() // Establish a "read lock"

		d := failures - threshold

		if d >= 0 {
			shouldRetryAt := last.Add((2 << d) * time.Second)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return "", errors.New("service unreachable")
			}
		}

		m.RUnlock() // Release read lock

		response, err := circuit(ctx) // Issue the request proper

		m.Lock() // Lock around shared resources
		defer m.Unlock()

		last = time.Now() // Record time of attempt

		if err != nil { // Circuit returned an error,
			failures++           // so we count the failure
			return response, err // and return
		}

		failures = 0 // Reset failures counter

		return response, nil
	}
}

func main() {
	// Function with error
	faultyService := func(ctx context.Context) (string, error) {
		return "", errors.New("service error")
	}

	// Handle function
	protectedService := Breaker(faultyService, 3)

	// Run
	for i := 0; i < 10; i++ {
		result, err := protectedService(context.Background())
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	}
}

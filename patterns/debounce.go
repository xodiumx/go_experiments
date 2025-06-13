package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

func DebounceFirst(circuit Circuit, d time.Duration) Circuit {
	var threshold time.Time
	var result string
	var err error
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()
		defer m.Unlock()

		if time.Now().Before(threshold) {
			return result, err
		}

		result, err = circuit(ctx)
		threshold = time.Now().Add(d)

		return result, err
	}
}

func main() {
	// An example of a function that we want to restrict.
	expensiveOperation := func(ctx context.Context) (string, error) {
		fmt.Println("Heavy operation processing...")
		return "function result", nil
	}

	// Restricted function
	debounced := DebounceFirst(expensiveOperation, 2*time.Second)

	// Run
	for i := 0; i < 5; i++ {
		res, err := debounced(context.Background())
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", res)
		}
		time.Sleep(500 * time.Millisecond) // timeout imitation
	}

	// Run again
	time.Sleep(3 * time.Second)
	res, err := debounced(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}
}

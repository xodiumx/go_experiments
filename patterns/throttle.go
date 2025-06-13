package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type Effector func(context.Context) (string, error)

func Throttle(e Effector, max uint, refill uint, d time.Duration) Effector {
	var tokens = max
	var once sync.Once
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		if ctx.Err() != nil {
			return "", ctx.Err()
		}

		once.Do(func() {
			ticker := time.NewTicker(d)

			go func() {
				defer ticker.Stop()

				for {
					fmt.Println("Tokens: ", tokens)
					select {
					case <-ctx.Done():
						return

					case <-ticker.C:
						m.Lock()
						t := tokens + refill
						if t > max {
							t = max
						}
						tokens = t
						m.Unlock()
					}
				}
			}()
		})

		m.Lock()
		defer m.Unlock()

		if tokens <= 0 {
			return "", fmt.Errorf("too many calls")
		}

		tokens--

		return e(ctx)
	}
}

func main() {
	// Restricted function
	myOperation := func(ctx context.Context) (string, error) {
		return "Successfully!", nil
	}

	// Throttled funcition, max 2 requests, update 1 token/sec
	throttledOperation := Throttle(myOperation, 2, 1, time.Second)

	// Run
	for i := 0; i < 15; i++ {
		ctx := context.Background()
		result, err := throttledOperation(ctx)
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			log.Printf("Result: %s", result)
		}
		time.Sleep(400 * time.Millisecond) // Imitate timeout
	}
}

package main

import (
	"context"
	"errors"
	"log"
	"time"
)

type Effector func(context.Context) (string, error)

func Retry(effector Effector, retries int, delay time.Duration) Effector {
	return func(ctx context.Context) (string, error) {
		for r := 0; ; r++ {
			response, err := effector(ctx)
			if err == nil || r >= retries {
				return response, err
			}

			log.Printf("Attempt %d failed; retrying in %v", r+1, delay)

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
	}
}

func main() {
	// counter
	attempt := 0

	// retries
	retry := 5

	// Function with error
	unreliableOperation := func(ctx context.Context) (string, error) {
		attempt++
		// Imitate erorr
		if attempt <= retry {
			return "", errors.New("temporary failure")
		}
		return "Successfully!", nil
	}

	// Function with retry
	retryOperation := Retry(unreliableOperation, retry, 2*time.Second)

	// Run
	ctx := context.Background()
	result, err := retryOperation(ctx)
	if err != nil {
		log.Fatalf("Error after all retries: %v", err)
	}
	log.Printf("Result: %s", result)
}

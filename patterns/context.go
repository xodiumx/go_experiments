package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var ErrCustomTimeout = errors.New("custom timeout error: operation took too long")

func slowOperation(ctx context.Context) (int, error) {
	select {
	case <-time.After(5 * time.Second):
		return 1, nil
	case <-ctx.Done():
		return 0, ErrCustomTimeout
	}
}

func Stream(ctx context.Context, out chan<- int) error {
	dctx, dcancel := context.WithTimeout(ctx, time.Second*3) // timeout for slow operation
	defer dcancel()

	res, err := slowOperation(dctx)
	if err != nil {
		return err
	}

	select {
	case out <- res:
		return nil
	case <-dctx.Done():
		return dctx.Err()
	}
}

func main() {
	ch1 := make(chan int, 1)
	ctx := context.Background()

	err := Stream(ctx, ch1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success:", <-ch1)
	}
}

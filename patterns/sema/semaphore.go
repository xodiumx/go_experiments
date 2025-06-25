package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

func main() {
	sem := semaphore.NewWeighted(3) // максимум 3 горутины одновременно
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		go func(n int) {
			if err := sem.Acquire(ctx, 1); err != nil {
				fmt.Println("acquire error:", err)
				return
			}
			defer sem.Release(1)

			// критическая секция
			time.Sleep(1 * time.Second)
			fmt.Println("running", n)
		}(i)
	}

	select {} // ждём
}

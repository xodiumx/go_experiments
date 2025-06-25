package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Worker get numbers from input and processing it
func worker(id int, input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range input {
		fmt.Printf("Worker %d processing number %d\n", id, num)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // Imitate work
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	taskCount := 101
	workerCount := 3

	// Distributes data from one channel to multiple channels
	tasks := make(chan int, taskCount)
	var wg sync.WaitGroup

	// Run 3 workers (gorutines)
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Send tasks in garbage
	for i := 1; i <= taskCount; i++ {
		tasks <- i
	}
	close(tasks) // Closing the channel to let workers know that the tasks are over

	// Wait all workers
	wg.Wait()

	fmt.Println("All tasks completed.")
}

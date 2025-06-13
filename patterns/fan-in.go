package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func generator(name string, wg *sync.WaitGroup) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		defer wg.Done()
		for i := 0; i < 5; i++ {
			value := name + "  " + strconv.Itoa(rand.Intn(100))
			out <- value
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		}
	}()
	return out
}

func fanIn(channels ...<-chan string) <-chan string {
	// merge many channels into 1 and reads from it
	out := make(chan string)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	wg.Add(2)
	ch1 := generator("Source 1", &wg)
	ch2 := generator("Source 2", &wg)

	merged := fanIn(ch1, ch2)

	for val := range merged {
		fmt.Println(val)
	}

	wg.Wait()
}

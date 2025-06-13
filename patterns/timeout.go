package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func get_req(ctx context.Context) (int, error) {
	// Use NewRequestWithContext for send context in request
	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.ya.ru", nil)
	if err != nil {
		return 0, fmt.Errorf("Error in request creation: %v", err)
	}

	// Imitate delay
	time.Sleep(3 * time.Second)

	// Make request
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error in request operation: %v", err)
	}
	defer response.Body.Close()

	return response.StatusCode, nil
}

func main() {
	ctx := context.Background()
	dctx, dcancel := context.WithTimeout(ctx, time.Second*2)
	defer dcancel()

	result, error_message := get_req(dctx)
	fmt.Println(result, error_message)
}

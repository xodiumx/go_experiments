package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	pb "grpc-keyvalue/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewKeyValueClient(conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		args := strings.Fields(input)

		log.Printf("Args: %v", args)

		if len(args) == 0 {
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		switch args[0] {
		case "get":
			if len(args) < 2 {
				fmt.Println("Usage: get KEY")
				continue
			}
			r, err := client.Get(ctx, &pb.GetRequest{Key: args[1]})
			if err != nil {
				log.Printf("Get error: %v", err)
			} else {
				fmt.Printf("Value: %s\n", r.Value)
			}

		case "put":
			if len(args) < 3 {
				fmt.Println("Usage: put KEY VALUE")
				continue
			}
			_, err := client.Put(ctx, &pb.PutRequest{Key: args[1], Value: strings.Join(args[2:], " ")})
			if err != nil {
				log.Printf("Put error: %v", err)
			} else {
				fmt.Println("OK")
			}

		case "exit":
			return

		default:
			fmt.Println("Commands: get KEY | put KEY VALUE | exit")
		}
	}
}

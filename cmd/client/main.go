package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Ayupov-Ayaz/grpc/cmd/client/internal"
)

const port = 50_000

func main() {
	if err := run(); err != nil {
		log.Fatalf("run: %v", err)
	}
}

func run() error {
	client, err := createClient()
	if err != nil {
		return fmt.Errorf("create client: %w", err)
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("close rpc client: %v", err)
		}
	}()

	return nil
}

func createClient() (*internal.Client, error) {
	fmt.Println("Creating client...")

	client, err := internal.New(makeAddress())
	if err != nil {
		return nil, fmt.Errorf("new client: %w", err)
	}

	return client, nil
}

func makeAddress() string {
	return "localhost:" + strconv.Itoa(port)
}

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Ayupov-Ayaz/grpc/cmd/server/internal"
	"github.com/Ayupov-Ayaz/grpc/cmd/server/internal/transaction"
	"github.com/Ayupov-Ayaz/grpc/cmd/server/internal/wallet"
	"github.com/Ayupov-Ayaz/grpc/gen/go/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = 50_000

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	app := internal.NewApplication(
		wallet.NewInMemoryStore(),
		transaction.NewInMemoryStore(),
	)

	return startGRPServer(app)
}

func startGRPServer(app *internal.Application) error {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return fmt.Errorf("listen: %v", err)
	}

	server := grpc.NewServer()
	api.RegisterTransactionServiceServer(server, app)
	api.RegisterWalletServiceServer(server, app)

	// Register reflection service on gRPC server.
	// This is useful for debugging and testing.
	reflection.Register(server)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("gRPC server started on port %d", port)

		if err := server.Serve(lis); err != nil {
			fmt.Printf("gRPC server error: %v\n", err)
		}
	}()

	<-stop

	log.Println("Stopping gRPC server...")

	return nil
}

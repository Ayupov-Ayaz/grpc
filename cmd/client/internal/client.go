package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/Ayupov-Ayaz/grpc/cmd/client/internal/actions"
	api "github.com/Ayupov-Ayaz/grpc/gen/go/aayupov/wallet/v1alpha1"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	close func() error
	src   api.TransactionServiceClient
}

func New(address string) (*Client, error) {
	const timeout = 5 * time.Second

	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(timeout),
		grpc.WithBlock())

	if err != nil {
		return nil, fmt.Errorf("rpc connect [%s]: %w", address, err)
	}

	return &Client{
		src:   api.NewTransactionServiceClient(conn),
		close: closeConnectionCallback(conn),
	}, nil
}

func (c *Client) Close() error {
	return c.close()
}

func (c *Client) Bet(ctx context.Context, bet actions.Bet) (int64, error) {
	req := &api.BetRequest{
		Amount:      bet.GetAmount(),
		OperationId: bet.GetID(),
		UserId:      bet.GetUserID(),
	}

	resp, err := c.src.Bet(ctx, req)
	if err != nil {
		// TODO: handle error
		return 0, fmt.Errorf("rpc bet: %w", err)
	}

	return resp.GetBalance(), nil
}

func closeConnectionCallback(conn *grpc.ClientConn) func() error {
	return func() error {
		if err := conn.Close(); err != nil {
			return fmt.Errorf("rpc close: %w", err)
		}
		return nil
	}
}

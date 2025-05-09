package wallet

import "context"

type Store interface {
	CreateWallet(ctx context.Context, userID string) error
	Add(ctx context.Context, userID string, amount uint64) (int64, error)
	Subtract(ctx context.Context, userID string, amount uint64) (int64, error)
	GetBalance(ctx context.Context, userID string) (int64, error)
	CheckWallet(ctx context.Context, userID string) error
}

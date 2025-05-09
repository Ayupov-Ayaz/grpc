package internal

import (
	"context"

	"github.com/Ayupov-Ayaz/grpc/cmd/server/internal/transaction"
	"github.com/Ayupov-Ayaz/grpc/cmd/server/internal/wallet"
	api "github.com/Ayupov-Ayaz/grpc/gen/go/aayupov/wallet/v1alpha1"
)

type Application struct {
	api.UnimplementedTransactionServiceServer
	api.UnimplementedWalletServiceServer

	wallet      wallet.Store
	transaction transaction.Store
}

func NewApplication(
	wallet wallet.Store,
	transaction transaction.Store,
) *Application {
	return &Application{
		wallet:      wallet,
		transaction: transaction,
	}
}

func (a *Application) CreateWallet(
	ctx context.Context, req *api.CreateWalletRequest,
) (*api.CreateWalletResponse, error) {
	err := a.wallet.CreateWallet(ctx, req.GetUserId())
	if err != nil {
		return nil, wrapCreateWalletError(err)
	}

	balance, err := a.wallet.Add(ctx, req.GetUserId(), uint64(req.GetBalance()))
	if err != nil {
		return nil, wrapChangeWalletError(err)
	}

	return &api.CreateWalletResponse{
		Balance: balance,
	}, nil
}

func (a *Application) Bet(
	ctx context.Context, req *api.BetRequest,
) (*api.BetResponse, error) {
	err := a.wallet.CheckWallet(ctx, req.GetUserId())
	if err != nil {
		return nil, wrapChangeWalletError(err)
	}

	err = a.transaction.SetTransactionID(ctx, req.GetUserId(), req.GetId())
	if err != nil {
		return nil, wrapSetTransactionError(err)
	}

	balance, err := a.wallet.Subtract(ctx, req.GetUserId(), req.GetAmount())
	if err != nil {
		return nil, wrapChangeWalletError(err)
	}

	return &api.BetResponse{
		Balance: balance,
	}, nil
}

func (a *Application) Win(ctx context.Context, req *api.WinRequest) (*api.WinResponse, error) {
	err := a.wallet.CheckWallet(ctx, req.GetUserId())
	if err != nil {
		return nil, wrapChangeWalletError(err)
	}

	err = a.transaction.CheckTransactionID(ctx,
		req.GetUserId(), req.GetBetTransactionId())
	if err != nil {
		return nil, wrapCheckBetTransactionError(err)
	}

	err = a.transaction.SetTransactionID(ctx, req.GetUserId(), req.GetId())
	if err != nil {
		return nil, wrapSetTransactionError(err)
	}

	balance, err := a.wallet.Add(ctx, req.GetUserId(), req.GetAmount())
	if err != nil {
		return nil, wrapChangeWalletError(err)
	}

	return &api.WinResponse{
		Balance: balance,
	}, nil
}

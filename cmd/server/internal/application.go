package internal

import (
	"context"

	"github.com/Ayupov-Ayaz/grpc/cmd/server/internal/transaction"
	"github.com/Ayupov-Ayaz/grpc/cmd/server/internal/wallet"
	"github.com/Ayupov-Ayaz/grpc/gen/go/api/v1"
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

func (a *Application) CreateWallet(ctx context.Context, req *api.CreateWalletRequest) (*api.CreateWalletResponse, error) {
	err := a.wallet.CreateWallet(req.UserId)
	if err != nil {
		// TODO: gRPC error
		return nil, err
	}

	balance, err := a.wallet.Add(req.GetUserId(), uint64(req.GetBalance()))
	if err != nil {
		// TODO: gRPC error
		return nil, err
	}

	return &api.CreateWalletResponse{
		Balance: balance,
	}, nil
}

func (a *Application) Bet(ctx context.Context, req *api.BetRequest) (*api.BetResponse, error) {
	err := a.transaction.SetTransactionID(req.UserId, req.OperationId)
	if err != nil {
		// TODO: gRPC error
		return nil, err
	}

	balance, err := a.wallet.Subtract(req.UserId, req.Amount)
	if err != nil {
		// TODO: gRPC error
		return nil, err
	}

	return &api.BetResponse{
		Balance: balance,
	}, nil
}

func (a *Application) Win(ctx context.Context, req *api.WinRequest) (*api.WinResponse, error) {
	err := a.transaction.CheckTransactionID(req.UserId, req.BetOperationId)
	if err != nil {
		// TODO: gRPC error
		return nil, err
	}

	err = a.transaction.SetTransactionID(req.UserId, req.OperationId)
	if err != nil {
		// TODO: gRPC error
		return nil, err
	}

	balance, err := a.wallet.Add(req.UserId, req.Amount)
	if err != nil {
		// TODO: gRPC error
		return nil, err
	}

	return &api.WinResponse{
		Balance: balance,
	}, nil
}

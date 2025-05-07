package wallet

import "errors"

var (
	ErrInsufficientFunds   = errors.New("insufficient funds")
	ErrWalletNotFound      = errors.New("wallet not found")
	ErrWalletAlreadyExists = errors.New("wallet already exists")
)

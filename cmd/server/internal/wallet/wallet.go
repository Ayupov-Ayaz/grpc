package wallet

type Store interface {
	CreateWallet(userID string) error
	Add(userID string, amount uint64) (int64, error)
	Subtract(userID string, amount uint64) (int64, error)
	GetBalance(userID string) (int64, error)
}

package wallet

type NotFoundError struct {
	userID string
}

func ErrNotFound(userID string) *NotFoundError {
	return &NotFoundError{
		userID: userID,
	}
}

func (e *NotFoundError) UserID() string { return e.userID }

func (e *NotFoundError) Error() string {
	return "wallet not found"
}

type AlreadyExistsError struct {
	userID string
}

func ErrAlreadyExist(userID string) *AlreadyExistsError {
	return &AlreadyExistsError{
		userID: userID,
	}
}

func (e *AlreadyExistsError) UserID() string { return e.userID }

func (e *AlreadyExistsError) Error() string {
	return "wallet already exists"
}

type InsufficientFundsError struct {
	userID  string
	balance int64
	bet     uint64
}

func ErrInsufficientFunds(
	userID string,
	bet uint64,
	balance int64,
) *InsufficientFundsError {
	return &InsufficientFundsError{
		userID:  userID,
		bet:     bet,
		balance: balance,
	}
}

func (e *InsufficientFundsError) UserID() string { return e.userID }
func (e *InsufficientFundsError) Balance() int64 { return e.balance }
func (e *InsufficientFundsError) Bet() uint64    { return e.bet }

func (e *InsufficientFundsError) Error() string {
	return "insufficient funds"
}

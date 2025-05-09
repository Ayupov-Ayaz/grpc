package transaction

type NotFoundError struct {
	transactionID ID
}

func ErrNotFound(transactionID ID) *NotFoundError {
	return &NotFoundError{
		transactionID: transactionID,
	}
}
func (e *NotFoundError) TransactionID() ID { return e.transactionID }

func (e *NotFoundError) Error() string {
	return "transaction " + e.transactionID.String() + " not found"
}

type AlreadyExistsError struct {
	transactionID ID
}

func ErrAlreadyExists(transactionID ID) *AlreadyExistsError {
	return &AlreadyExistsError{
		transactionID: transactionID,
	}
}

func (e *AlreadyExistsError) TransactionID() ID { return e.transactionID }

func (e *AlreadyExistsError) Error() string {
	return "transaction " + e.transactionID.String() + " already exists"
}

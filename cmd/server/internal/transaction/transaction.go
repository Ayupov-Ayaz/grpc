package transaction

type Store interface {
	SetTransactionID(userID, operationID string) error
	CheckTransactionID(userID, operationID string) error
}

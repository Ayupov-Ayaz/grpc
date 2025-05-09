package transaction

import "context"

type Store interface {
	SetTransactionID(ctx context.Context, userID, operationID string) error
	CheckTransactionID(ctx context.Context, userID, operationID string) error
}

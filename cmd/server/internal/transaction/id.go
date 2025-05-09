package transaction

type id string

func newID(userID, operationID string) id {
	const separator = ":"

	return id(userID + separator + operationID)
}

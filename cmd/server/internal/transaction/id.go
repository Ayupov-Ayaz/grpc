package transaction

import "strings"

type ID string

const (
	separator        = ":"
	userIDIndex      = 0
	operationIDIndex = 1
	idComponentsLen  = 2
)

func newID(userID, operationID string) ID {
	return ID(userID + separator + operationID)
}

func (id ID) GetUserID() string {
	parts := id.split()
	if len(parts) <= userIDIndex {
		return ""
	}

	return parts[userIDIndex]
}

func (id ID) GetOperationID() string {
	parts := id.split()
	if len(parts) <= operationIDIndex {
		return ""
	}

	return parts[operationIDIndex]
}

func (id ID) String() string {
	return string(id)
}

func (id ID) split() []string {
	return strings.SplitN(string(id), separator, idComponentsLen)
}

package transaction

import (
	"sync"
)

type InMemoryStore struct {
	mu          *sync.Mutex
	transaction map[ID]struct{}
}

var _ Store = (*InMemoryStore)(nil)

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		mu:          &sync.Mutex{},
		transaction: make(map[ID]struct{}),
	}
}

func (s *InMemoryStore) SetTransactionID(userID, operationID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	transactionID := newID(userID, operationID)

	exist := s.hasTransactionIDs(transactionID)
	if exist {
		return ErrAlreadyExists(transactionID)
	}

	s.setTransactionID(transactionID)

	return nil
}

func (s *InMemoryStore) CheckTransactionID(userID, operationID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	transactionID := newID(userID, operationID)

	exist := s.hasTransactionIDs(transactionID)
	if !exist {
		return ErrNotFound(transactionID)
	}

	return nil
}

func (s *InMemoryStore) hasTransactionIDs(transactionID ID) bool {
	_, ok := s.transaction[transactionID]
	return ok
}

func (s *InMemoryStore) setTransactionID(transactionID ID) {
	s.transaction[transactionID] = struct{}{}
}

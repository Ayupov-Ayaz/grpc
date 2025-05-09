package wallet

import (
	"context"
	"sync"
)

type InMemoryStore struct {
	mu         *sync.Mutex
	userWallet map[string]int64
}

var _ Store = (*InMemoryStore)(nil)

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		mu:         &sync.Mutex{},
		userWallet: make(map[string]int64),
	}
}

func (s *InMemoryStore) CreateWallet(
	_ context.Context, userID string,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.userWallet[userID]; ok {
		return ErrAlreadyExist(userID)
	}

	s.userWallet[userID] = 0

	return nil
}

func (s *InMemoryStore) Add(
	_ context.Context, userID string, amount uint64,
) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	balance, err := s.getBalance(userID)
	if err != nil {
		return 0, err
	}

	balance += int64(amount)

	s.setBalance(userID, balance)

	return balance, nil
}

func (s *InMemoryStore) Subtract(
	_ context.Context, userID string, amount uint64,
) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	balance, err := s.getBalance(userID)
	if err != nil {
		return 0, err
	}

	if balance < int64(amount) {
		return 0, ErrInsufficientFunds(userID, amount, balance)
	}

	balance -= int64(amount)

	s.setBalance(userID, balance)

	return balance, nil
}

func (s *InMemoryStore) GetBalance(
	_ context.Context, userID string,
) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	balance, err := s.getBalance(userID)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (s *InMemoryStore) CheckWallet(
	_ context.Context, userID string,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.getBalance(userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *InMemoryStore) setBalance(userID string, balance int64) {
	s.userWallet[userID] = balance
}

func (s *InMemoryStore) getBalance(userID string) (int64, error) {
	if balance, ok := s.userWallet[userID]; ok {
		return balance, nil
	}

	return 0, ErrNotFound(userID)
}

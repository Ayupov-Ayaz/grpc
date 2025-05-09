package wallet

import "sync"

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

func (s *InMemoryStore) CreateWallet(userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.userWallet[userID]; ok {
		return ErrWalletAlreadyExists
	}

	s.userWallet[userID] = 0

	return nil
}

func (s *InMemoryStore) Add(userID string, amount uint64) (int64, error) {
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

func (s *InMemoryStore) Subtract(userID string, amount uint64) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	balance, err := s.getBalance(userID)
	if err != nil {
		return 0, err
	}

	if balance < int64(amount) {
		return 0, ErrInsufficientFunds
	}

	balance -= int64(amount)

	s.setBalance(userID, balance)

	return balance, nil
}

func (s *InMemoryStore) GetBalance(userID string) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	balance, err := s.getBalance(userID)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (s *InMemoryStore) setBalance(userID string, balance int64) {
	s.userWallet[userID] = balance
}

func (s *InMemoryStore) getBalance(userID string) (int64, error) {
	if balance, ok := s.userWallet[userID]; ok {
		return balance, nil
	}

	return 0, ErrWalletNotFound
}

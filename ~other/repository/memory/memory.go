package memory

import (
	"errors"
	"rep/repo"
	"sync"
)

type InMemoryUserRepo struct {
	data map[int]*repo.User
	mu   sync.RWMutex
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		data: make(map[int]*repo.User),
	}
}

func (r *InMemoryUserRepo) GetByID(id int) (*repo.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.data[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepo) Create(user *repo.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.data[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.data[user.ID] = user
	return nil
}

func (r *InMemoryUserRepo) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, id)
	return nil
}

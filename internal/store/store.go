package store

import (
	"errors"
	"sync"

	"github.com/davenfinch/followmyjourney-profile-service/internal/model"
)

var ErrNotFound = errors.New("not found")

// Store defines storage operations for profiles
type Store interface {
	Get(id string) (*model.Profile, error)
	Create(p *model.Profile) error
}

// InMemoryStore basic impl for development
type InMemoryStore struct {
	mu sync.RWMutex
	m  map[string]*model.Profile
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{m: make(map[string]*model.Profile)}
}

func (s *InMemoryStore) Get(id string) (*model.Profile, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	p, ok := s.m[id]
	if !ok {
		return nil, ErrNotFound
	}
	return p, nil
}

func (s *InMemoryStore) Create(p *model.Profile) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.m[p.GUID]; exists {
		return errors.New("already exists")
	}
	s.m[p.GUID] = p
	return nil
}

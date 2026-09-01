package service

import (
	"errors"

	"github.com/davenfinch/followmyjourney-profile-service/internal/model"
	"github.com/davenfinch/followmyjourney-profile-service/internal/store"
)

type ProfileService struct {
	store store.Store
}

func NewProfileService(s store.Store) *ProfileService {
	return &ProfileService{store: s}
}

func (s *ProfileService) GetProfile(id string) (*model.Profile, error) {
	return s.store.Get(id)
}

func (s *ProfileService) CreateProfile(p *model.Profile) error {
	if p == nil || p.GUID == "" {
		return errors.New("invalid profile")
	}
	return s.store.Create(p)
}

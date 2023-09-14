package cars

import (
	"context"
	"github.com/iamvbr/vehicle-service/internal/models"
	"github.com/iamvbr/vehicle-service/internal/services"
	"github.com/iamvbr/vehicle-service/internal/stores"
)

type Svc struct {
	store stores.CarStorage
}

func (s *Svc) Find(ctx context.Context, category string) ([]*models.Car, error) {
	return s.store.Find(ctx, category)
}

func (s *Svc) Get(ctx context.Context, id string) (*models.Car, error) {
	return s.store.Get(ctx, id)
}

func (s *Svc) Create(ctx context.Context, m *models.Car) (*models.Car, error) {
	if ok, _ := s.store.Exists(ctx, m.Id); ok {
		return nil, services.CarExists
	}

	return s.store.Create(ctx, m)
}

func (s *Svc) Update(ctx context.Context, id string, m *models.Car) (*models.Car, error) {
	if ok, _ := s.store.Exists(ctx, m.Id); !ok {
		return nil, services.CarNotFound
	}

	return s.store.Update(ctx, id, m)
}

func New(s stores.CarStorage) *Svc {
	return &Svc{store: s}
}

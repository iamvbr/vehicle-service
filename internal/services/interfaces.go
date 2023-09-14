package services

import (
	"context"
	"github.com/iamvbr/vehicle-service/internal/models"
)

type Car interface {
	Find(ctx context.Context, category string) ([]*models.Car, error)
	Get(ctx context.Context, id string) (*models.Car, error)
	Create(ctx context.Context, m *models.Car) (*models.Car, error)
	Update(ctx context.Context, id string, m *models.Car) (*models.Car, error)
}

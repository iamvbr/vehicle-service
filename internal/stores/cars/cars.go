package cars

import (
	"context"
	"errors"
	"github.com/iamvbr/vehicle-service/internal/models"
)

type Store struct {
	// TODO: map is not concurrency safe.
	cache map[string]int
	cars  []*models.Car
}

func (s *Store) Exists(ctx context.Context, id string) (bool, error) {
	_, ok := s.cache[id]
	return ok, nil
}

func New() *Store {
	return &Store{
		cache: map[string]int{},
		cars:  make([]*models.Car, 0),
	}
}

// Find TODO: needs pagination
func (s *Store) Find(ctx context.Context, category string) ([]*models.Car, error) {
	if category == "" {
		return s.cars, nil
	}

	var cars []*models.Car

	for _, car := range s.cars {
		if car.Category == category {
			cars = append(cars, car)
		}
	}

	return cars, nil
}

func (s *Store) Get(ctx context.Context, id string) (*models.Car, error) {
	i, ok := s.cache[id]
	if !ok {
		return nil, errors.New("car not found")
	}

	return s.cars[i], nil
}

func (s *Store) Create(ctx context.Context, m *models.Car) (*models.Car, error) {
	s.cache[m.Id] = len(s.cars)
	s.cars = append(s.cars, m)

	return m, nil
}

func (s *Store) Update(ctx context.Context, id string, m *models.Car) (*models.Car, error) {
	m.Id = id
	s.cars[s.cache[id]] = m

	return m, nil
}

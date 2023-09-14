package cars

import (
	"context"
	"github.com/iamvbr/vehicle-service/internal/models"
	"github.com/iamvbr/vehicle-service/internal/stores"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func setupMock(t *testing.T) *stores.MockCarStorage {
	ctrl := gomock.NewController(t)
	m := stores.NewMockCarStorage(ctrl)
	return m
}

func TestSvc_Create(t *testing.T) {
	mockStore := setupMock(t)

	car1 := &models.Car{
		Id:    "JHK123",
		Make:  "Ford",
		Model: "F10",
	}

	tests := []struct {
		name    string
		m       *models.Car
		want    *models.Car
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			name:    "Create a new car",
			m:       car1,
			want:    car1,
			wantErr: false,
			calls: []*gomock.Call{
				mockStore.EXPECT().Exists(gomock.Any(), car1.Id).Return(false, nil),
				mockStore.EXPECT().Create(gomock.Any(), car1).Return(car1, nil),
			},
		},
		{
			name:    "Trying to create it again",
			m:       car1,
			want:    nil,
			wantErr: true,
			calls: []*gomock.Call{
				mockStore.EXPECT().Exists(gomock.Any(), car1.Id).Return(true, nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(mockStore)

			got, err := s.Create(context.Background(), tt.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSvc_Update(t *testing.T) {
	mockStore := setupMock(t)

	car1 := &models.Car{
		Id:    "JHK123",
		Make:  "Ford",
		Model: "F10",
	}

	tests := []struct {
		name    string
		m       *models.Car
		want    *models.Car
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			name:    "update a car that doesn't exist",
			m:       car1,
			want:    nil,
			wantErr: true,
			calls: []*gomock.Call{
				mockStore.EXPECT().Exists(gomock.Any(), car1.Id).Return(false, nil),
			},
		},
		{
			name:    "update a car",
			m:       car1,
			want:    car1,
			wantErr: false,
			calls: []*gomock.Call{
				mockStore.EXPECT().Exists(gomock.Any(), car1.Id).Return(true, nil),
				mockStore.EXPECT().Update(gomock.Any(), car1.Id, car1).Return(car1, nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(mockStore)

			got, err := s.Update(context.Background(), tt.m.Id, tt.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

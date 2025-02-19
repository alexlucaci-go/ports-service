package ports

import (
	"context"
	"testing"

	"github.com/alexlucaci-go/ports-service/entities"
	"github.com/alexlucaci-go/ports-service/models"

	"github.com/stretchr/testify/require"
)

type testStore struct{}

//nolint:gocritic // test
func (testStore) Create(_ context.Context, _ entities.Port) error {
	return nil
}

//nolint:gocritic // test
func (testStore) Update(_ context.Context, _ entities.Port) error {
	return nil
}

func (testStore) Get(_ context.Context, _ string) (entities.Port, error) {
	return entities.Port{}, nil
}

func (testStore) Delete(_ context.Context, _ string) error {
	return nil
}

func (testStore) List(_ context.Context, _ int) ([]entities.Port, error) {
	return nil, nil
}

func TestCreate_not_all_coordinates(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name          string
		cp            models.CreatePort
		expectedError error
	}{
		{
			name:          "missing coordinates",
			cp:            models.CreatePort{},
			expectedError: ErrNoCoordinates,
		},
		{
			name: "empty coordinates",
			cp: models.CreatePort{
				Coordinates: []float64{},
			},
			expectedError: ErrNoCoordinates,
		},
		{
			name: "missing one coordinate",
			cp: models.CreatePort{
				Coordinates: []float64{15.2},
			},
			expectedError: ErrNoCoordinates,
		},
		{
			name: "longitude out of bounds",
			cp: models.CreatePort{
				Coordinates: []float64{190, 15},
			},
			expectedError: ErrIncorrectLatitudeOrLongitudeValues,
		},
		{
			name: "latitude out of bounds",
			cp: models.CreatePort{
				Coordinates: []float64{180, 95},
			},
			expectedError: ErrIncorrectLatitudeOrLongitudeValues,
		},
		{
			name: "both latitude and longitude out of bounds",
			cp: models.CreatePort{
				Coordinates: []float64{190, 95},
			},
			expectedError: ErrIncorrectLatitudeOrLongitudeValues,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			d := NewService(nil)
			err := d.Create(context.Background(), tc.cp)
			require.EqualError(t, err, tc.expectedError.Error())
		})
	}
}

func TestUpdate_not_all_coordinates(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name          string
		up            models.UpdatePort
		expectedError error
	}{
		{
			name:          "missing coordinates",
			up:            models.UpdatePort{},
			expectedError: nil,
		},
		{
			name: "empty coordinates",
			up: models.UpdatePort{
				Coordinates: &[]float64{},
			},
			expectedError: ErrNoCoordinates,
		},
		{
			name: "missing one coordinate",
			up: models.UpdatePort{
				Coordinates: &[]float64{15.2},
			},
			expectedError: ErrNoCoordinates,
		},
		{
			name: "longitude out of bounds",
			up: models.UpdatePort{
				Coordinates: &[]float64{190, 15},
			},
			expectedError: ErrIncorrectLatitudeOrLongitudeValues,
		},
		{
			name: "latitude out of bounds",
			up: models.UpdatePort{
				Coordinates: &[]float64{180, 95},
			},
			expectedError: ErrIncorrectLatitudeOrLongitudeValues,
		},
		{
			name: "both latitude and longitude out of bounds",
			up: models.UpdatePort{
				Coordinates: &[]float64{190, 95},
			},
			expectedError: ErrIncorrectLatitudeOrLongitudeValues,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			d := NewService(testStore{})
			err := d.Update(context.Background(), "", tc.up)
			if tc.expectedError == nil {
				require.NoError(t, err, "updating")
				return
			}
			require.EqualError(t, err, tc.expectedError.Error(), "updating")
		})
	}
}

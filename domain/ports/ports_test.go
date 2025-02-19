package ports

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

//nolint:gocritic // test
func (TestStore) Create(_ context.Context, _ NewPort) error {
	return nil
}

//nolint:gocritic // test
func (TestStore) Update(_ context.Context, _ string, _ UpdatePort) error {
	return nil
}

func (TestStore) Get(_ context.Context, _ string) (Port, error) {
	return Port{}, nil
}

func (TestStore) Delete(_ context.Context, _ string) error {
	return nil
}

func (TestStore) List(_ context.Context, _ int) ([]Port, error) {
	return nil, nil
}

func TestCreate_not_all_coordinates(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name          string
		np            NewPort
		expectedError error
	}{
		{
			name:          "missing coordinates",
			np:            NewPort{},
			expectedError: ErrNoCoordinates,
		},
		{
			name: "empty coordinates",
			np: NewPort{
				Port: Port{
					Coordinates: []float64{},
				},
			},
			expectedError: ErrNoCoordinates,
		},
		{
			name: "missing one coordinate",
			np: NewPort{
				Port: Port{
					Coordinates: []float64{15.2},
				},
			},
			expectedError: ErrNoCoordinates,
		},
		{
			name: "longitude out of bounds",
			np: NewPort{
				Port: Port{
					Coordinates: []float64{190, 15},
				},
			},
			expectedError: ErrIncorrectLatitudeOrLongitudeValues,
		},
		{
			name: "latitude out of bounds",
			np: NewPort{
				Port: Port{
					Coordinates: []float64{180, 95},
				},
			},
			expectedError: ErrIncorrectLatitudeOrLongitudeValues,
		},
		{
			name: "both latitude and longitude out of bounds",
			np: NewPort{
				Port: Port{
					Coordinates: []float64{190, 95},
				},
			},
			expectedError: ErrIncorrectLatitudeOrLongitudeValues,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDomain(nil)
			err := d.Create(context.Background(), tc.np)
			require.EqualError(t, err, tc.expectedError.Error())
		})
	}
}

func TestUpdate_not_all_coordinates(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name          string
		up            UpdatePort
		expectedError error
	}{
		{
			name:          "missing coordinates",
			up:            UpdatePort{},
			expectedError: nil,
		},
		{
			name: "empty coordinates",
			up: UpdatePort{
				Coordinates: &[]float64{},
			},
			expectedError: ErrNoCoordinates,
		},
		{
			name: "missing one coordinate",
			up: UpdatePort{
				Coordinates: &[]float64{15.2},
			},
			expectedError: ErrNoCoordinates,
		},
		{
			name: "longitude out of bounds",
			up: UpdatePort{
				Coordinates: &[]float64{190, 15},
			},
			expectedError: ErrIncorrectLatitudeOrLongitudeValues,
		},
		{
			name: "latitude out of bounds",
			up: UpdatePort{
				Coordinates: &[]float64{180, 95},
			},
			expectedError: ErrIncorrectLatitudeOrLongitudeValues,
		},
		{
			name: "both latitude and longitude out of bounds",
			up: UpdatePort{
				Coordinates: &[]float64{190, 95},
			},
			expectedError: ErrIncorrectLatitudeOrLongitudeValues,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDomain(TestStore{})
			err := d.Update(context.Background(), "", tc.up)
			if tc.expectedError == nil {
				require.NoError(t, err, "updating")
				return
			}
			require.EqualError(t, err, tc.expectedError.Error(), "updating")
		})
	}
}

package ports

import (
	"github.com/stretchr/testify/require"
	"testing"
)

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
			err := d.Create(nil, tc.np)
			require.EqualError(t, err, tc.expectedError.Error())
		})
	}
}

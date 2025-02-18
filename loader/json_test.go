package loader

import (
	"context"
	"testing"

	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/alexlucaci-go/ports-service/infrastructure/store/inmemorydb"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestJsonLoader(t *testing.T) {
	t.Parallel()
	portDomain := ports.NewDomain(inmemorydb.NewInMemoryDB())
	jsonLoader := NewJSON(portDomain)
	err := jsonLoader.LoadFromFile(context.Background(), "ports_test.json")
	require.NoError(t, err)

	// pick last port id from test file
	got, err := portDomain.Get(context.Background(), "ZWUTA")
	require.NoError(t, err, "getting last port from file")

	expected := ports.Port{
		ID:          "ZWUTA",
		Name:        "Mutare",
		City:        "Mutare",
		Country:     "Zimbabwe",
		Coordinates: []float64{32.670573, -18.9707},
		Province:    "Manicaland",
		Timezone:    "Africa/Harare",
		Unlocs:      []string{"ZWUTA"},
	}

	diff := cmp.Diff(expected, got)
	require.NotEmpty(t, diff, "comparing expected and got ports")
}

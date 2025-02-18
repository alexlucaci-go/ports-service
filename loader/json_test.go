package loader

import (
	"context"
	"testing"

	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/alexlucaci-go/ports-service/infrastructure/store/inmemorydb"

	"github.com/stretchr/testify/require"
)

func TestJsonLoader(t *testing.T) {
	t.Parallel()
	portdomain := ports.NewDomain(inmemorydb.NewInMemoryDB())
	jsonLoader := NewJSON(portdomain)
	err := jsonLoader.LoadFromFile(context.Background(), "ports_test.json")
	require.NoError(t, err)

	id := "ZWUTA"

	port, err := portdomain.Get(context.Background(), id)
	require.NoError(t, err, "getting last port from file")

	require.Equal(t, "ZWUTA", port.Unlocs[0])
}

package inmemorydb

import (
	"context"
	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSavePort(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	id := "AEAJM"
	port := ports.Port{
		Name:        "Ajman",
		City:        "Ajman",
		Country:     "United Arab Emirates",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{55.5136433, 25.4052165},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	}

	err := store.Create(ctx, id, port)
	require.NoError(t, err, "creating port")

	savedPort, err := store.Get(ctx, id)
	require.NoError(t, err, "getting port after creation")
	require.NotNil(t, savedPort)
	require.Equal(t, port, savedPort, "comparing saved port with original port")
}

func TestSavePort_AlreadyExistingID(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	id := "AEAJM"
	port := ports.Port{
		Name:        "Ajman",
		City:        "Ajman",
		Country:     "United Arab Emirates",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{55.5136433, 25.4052165},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	}

	err := store.Create(ctx, id, port)
	require.NoError(t, err, "creating port")

	err = store.Create(ctx, id, port)
	require.Error(t, err, "creating port with already existing id")
	require.Equal(t, ports.ErrAlreadyExists, err, "checking error type")
}

func TestUpdatePort_OneField(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	id := "AEAJM"
	port := ports.Port{
		Name:        "Ajman",
		City:        "Ajman",
		Country:     "United Arab Emirates",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{55.5136433, 25.4052165},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	}

	err := store.Create(ctx, id, port)
	require.NoError(t, err, "creating port")

	up := ports.UpdatePort{
		Name: ports.StringToPointerString("Ajman_test"),
	}

	err = store.Update(ctx, id, up)
	require.NoError(t, err, "updating port")

	updatedPort := port
	updatedPort.Name = "Ajman_test"

	savedPort, err := store.Get(ctx, id)
	require.NoError(t, err, "getting saved port after update")
	require.NotNil(t, savedPort)
	require.Equal(t, updatedPort, savedPort, "comparing updated port with saved port")
}

func TestUpdatePort_NotExistingID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := NewInMemoryDB()
	id := "AEAJM"
	up := ports.UpdatePort{
		Name: ports.StringToPointerString("Ajman_test"),
	}

	err := store.Update(ctx, id, up)
	require.Error(t, err, "updating port with not existing id")
	require.Equal(t, ports.ErrNotFound, err, "checking error type")
}

func TestUpdatePort_EmptyUpdate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	id := "AEAJM"
	port := ports.Port{
		Name:        "Ajman",
		City:        "Ajman",
		Country:     "United Arab Emirates",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{55.5136433, 25.4052165},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	}

	err := store.Create(ctx, id, port)
	require.NoError(t, err, "creating port")

	up := ports.UpdatePort{}

	err = store.Update(ctx, id, up)
	require.NoError(t, err, "updating port with empty update")

	savedPort, err := store.Get(ctx, id)
	require.NoError(t, err, "getting port after update")

	require.NotNil(t, savedPort)
	require.Equal(t, port, savedPort, "comparing saved port with original port")
}

func TestDeletePort(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	id := "AEAJM"
	port := ports.Port{
		Name:        "Ajman",
		City:        "Ajman",
		Country:     "United Arab Emirates",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{55.5136433, 25.4052165},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	}

	err := store.Create(ctx, id, port)
	require.NoError(t, err, "creating port")

	got, err := store.Get(ctx, id)
	require.NoError(t, err, "getting port after creation")
	require.NotEmpty(t, got)

	err = store.Delete(ctx, id)
	require.NoError(t, err, "deleting port")

	_, err = store.Get(ctx, id)
	require.EqualError(t, err, ports.ErrNotFound.Error())
}

func TestListPorts_limit_less_than_length(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	id := "AEAJM"
	port := ports.Port{
		Name:        "Ajman",
		City:        "Ajman",
		Country:     "United Arab Emirates",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{55.5136433, 25.4052165},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	}

	err := store.Create(ctx, id, port)
	require.NoError(t, err, "creating 1st port")

	err = store.Create(ctx, id+"_1", port)
	require.NoError(t, err, "creating 2nd port")

	listedPorts, err := store.List(ctx, 1)
	require.NoError(t, err, "listing ports with limit 1")
	require.Equal(t, 1, len(listedPorts), "checking length of listed ports")
}

func TestListPorts_limit_equal_with_length(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	id := "AEAJM"
	port := ports.Port{
		Name:        "Ajman",
		City:        "Ajman",
		Country:     "United Arab Emirates",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{55.5136433, 25.4052165},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	}

	err := store.Create(ctx, id, port)
	require.NoError(t, err, "creating 1st port")

	err = store.Create(ctx, id+"_1", port)
	require.NoError(t, err, "creating 2nd port")

	listedPorts, err := store.List(ctx, 2)
	require.NoError(t, err, "listing ports with limit 1")
	require.Equal(t, 2, len(listedPorts), "checking length of listed ports")
}

func TestListPorts_limit_greater_than_length(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	id := "AEAJM"
	port := ports.Port{
		Name:        "Ajman",
		City:        "Ajman",
		Country:     "United Arab Emirates",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{55.5136433, 25.4052165},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	}

	err := store.Create(ctx, id, port)
	require.NoError(t, err, "creating 1st port")

	err = store.Create(ctx, id+"_1", port)
	require.NoError(t, err, "creating 2nd port")

	listedPorts, err := store.List(ctx, 3)
	require.NoError(t, err, "listing ports with limit 1")
	require.Equal(t, 2, len(listedPorts), "checking length of listed ports")
}

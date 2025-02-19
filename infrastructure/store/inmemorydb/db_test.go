package inmemorydb

import (
	"context"
	"testing"

	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/alexlucaci-go/ports-service/entities"

	"github.com/stretchr/testify/require"
)

const testID = "AEAJM_1"

func TestSavePort(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	port := entities.Port{
		ID:          "AEAJM",
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

	err := store.Create(ctx, port)
	require.NoError(t, err, "creating port")

	savedPort, err := store.Get(ctx, port.ID)
	require.NoError(t, err, "getting port after creation")
	require.NotNil(t, savedPort)
	require.Equal(t, port, savedPort, "comparing saved port with original port")
}

func TestSavePort_AlreadyExistingID(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	port := entities.Port{
		ID:          "AEAJM",
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

	err := store.Create(ctx, port)
	require.NoError(t, err, "creating port")

	err = store.Create(ctx, port)
	require.Error(t, err, "creating port with already existing testID")
	require.Equal(t, ports.ErrAlreadyExists, err, "checking error type")
}

func TestUpdatePort_OneField(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	port := entities.Port{
		ID:          "AEAJM",
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

	err := store.Create(ctx, port)
	require.NoError(t, err, "creating port")

	port.Name = "Ajman_test"

	err = store.Update(ctx, port)
	require.NoError(t, err, "updating port")

	savedPort, err := store.Get(ctx, port.ID)
	require.NoError(t, err, "getting saved port after update")
	require.NotNil(t, savedPort)
	require.Equal(t, port, savedPort, "comparing updated port with saved port")
}

func TestUpdatePort_NotExistingID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := NewInMemoryDB()
	port := entities.Port{ID: "AEAJM"}

	err := store.Update(ctx, port)
	require.Error(t, err, "updating port with not existing id")
	require.Equal(t, ports.ErrNotFound, err, "checking error type")
}

func TestUpdatePort_EmptyUpdate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	port := entities.Port{
		ID:          "AEAJM",
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

	err := store.Create(ctx, port)
	require.NoError(t, err, "creating port")

	err = store.Update(ctx, port)
	require.NoError(t, err, "updating port with empty update")

	savedPort, err := store.Get(ctx, port.ID)
	require.NoError(t, err, "getting port after update")

	require.NotNil(t, savedPort)
	require.Equal(t, port, savedPort, "comparing saved port with original port")
}

func TestDeletePort(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	port := entities.Port{
		ID:          "AEAJM",
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

	err := store.Create(ctx, port)
	require.NoError(t, err, "creating port")

	got, err := store.Get(ctx, port.ID)
	require.NoError(t, err, "getting port after creation")
	require.NotEmpty(t, got)

	err = store.Delete(ctx, port.ID)
	require.NoError(t, err, "deleting port")

	_, err = store.Get(ctx, port.ID)
	require.EqualError(t, err, ports.ErrNotFound.Error())
}

func TestListPorts_limit_less_than_length(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	port := entities.Port{
		ID:          "AEAJM",
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

	err := store.Create(ctx, port)
	require.NoError(t, err, "creating 1st port")

	port.ID = testID

	err = store.Create(ctx, port)
	require.NoError(t, err, "creating 2nd port")

	listedPorts, err := store.List(ctx, 1)
	require.NoError(t, err, "listing ports with limit 1")
	require.Len(t, listedPorts, 1, "checking length of listed ports")
}

func TestListPorts_limit_equal_with_length(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	port := entities.Port{
		ID:          "AEAJM",
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

	err := store.Create(ctx, port)
	require.NoError(t, err, "creating 1st port")

	port.ID = testID

	err = store.Create(ctx, port)
	require.NoError(t, err, "creating 2nd port")

	listedPorts, err := store.List(ctx, 2)
	require.NoError(t, err, "listing ports with limit 1")
	require.Len(t, listedPorts, 2, "checking length of listed ports")
}

func TestListPorts_limit_greater_than_length(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	store := NewInMemoryDB()
	port := entities.Port{
		ID:          "AEAJM",
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

	err := store.Create(ctx, port)
	require.NoError(t, err, "creating 1st port")

	port.ID = testID

	err = store.Create(ctx, port)
	require.NoError(t, err, "creating 2nd port")

	listedPorts, err := store.List(ctx, 3)
	require.NoError(t, err, "listing ports with limit 1")
	require.Len(t, listedPorts, 2, "checking length of listed ports")
}

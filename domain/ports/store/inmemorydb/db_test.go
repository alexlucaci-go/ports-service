package inmemorydb

import (
	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSavePort(t *testing.T) {
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

	err := store.Create(id, port)
	require.NoError(t, err, "creating port")

	savedPort, err := store.Get(id)
	require.NoError(t, err, "getting port after creation")
	require.NotNil(t, savedPort)
	require.Equal(t, port, savedPort, "comparing saved port with original port")
}

func TestSavePort_AlreadyExistingID(t *testing.T) {
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

	err := store.Create(id, port)
	require.NoError(t, err, "creating port")

	err = store.Create(id, port)
	require.Error(t, err, "creating port with already existing id")
	require.Equal(t, ErrAlreadyExists, err, "checking error type")
}

func TestUpdatePort_OneField(t *testing.T) {
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

	err := store.Create(id, port)
	require.NoError(t, err, "creating port")

	up := ports.UpdatePort{
		Name: ports.StringToPointerString("Ajman_test"),
	}

	err = store.Update(id, up)
	require.NoError(t, err, "updating port")

	updatedPort := port
	updatedPort.Name = "Ajman_test"

	savedPort, err := store.Get(id)
	require.NoError(t, err, "getting saved port after update")
	require.NotNil(t, savedPort)
	require.Equal(t, updatedPort, savedPort, "comparing updated port with saved port")
}

func TestUpdatePort_MultipleFields(t *testing.T) {
	//...
}

func TestUpdatePort_NotExistingID(t *testing.T) {
	store := NewInMemoryDB()
	id := "AEAJM"
	up := ports.UpdatePort{
		Name: ports.StringToPointerString("Ajman_test"),
	}

	err := store.Update(id, up)
	require.Error(t, err, "updating port with not existing id")
	require.Equal(t, ErrNotFound, err, "checking error type")
}

func TestUpdatePort_EmptyUpdate(t *testing.T) {
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

	err := store.Create(id, port)
	require.NoError(t, err, "creating port")

	up := ports.UpdatePort{}

	err = store.Update(id, up)
	require.NoError(t, err, "updating port with empty update")

	savedPort, err := store.Get(id)
	require.NoError(t, err, "getting port after update")

	require.NotNil(t, savedPort)
	require.Equal(t, port, savedPort, "comparing saved port with original port")
}

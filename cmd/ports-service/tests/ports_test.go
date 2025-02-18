package tests

import (
	"bytes"
	"encoding/json"
	"github.com/alexlucaci-go/ports-service/cmd/ports-service/handlers"
	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/alexlucaci-go/ports-service/domain/ports/store/inmemorydb"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreatePort(t *testing.T) {
	t.Parallel()
	api := handlers.API(make(chan os.Signal), inmemorydb.NewInMemoryDB())
	portData := ports.NewPort{
		ID: "AEAJM",
		Port: ports.Port{
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
		},
	}

	body, _ := json.Marshal(portData)
	req := httptest.NewRequest(http.MethodPost, "/v1/ports", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusCreated, res.StatusCode)
}

func TestCreatePort_AlreadyExisting(t *testing.T) {
	t.Parallel()
	api := handlers.API(make(chan os.Signal), inmemorydb.NewInMemoryDB())
	portData := ports.NewPort{
		ID: "AEAJM",
		Port: ports.Port{
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
		},
	}

	body, err := json.Marshal(portData)
	require.NoError(t, err, "marshalling request body")
	req := httptest.NewRequest(http.MethodPost, "/v1/ports", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusCreated, res.StatusCode)

	req = httptest.NewRequest(http.MethodPost, "/v1/ports", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec = httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res = rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusConflict, res.StatusCode)
}

func TestUpdatePort(t *testing.T) {
	t.Parallel()
	api := handlers.API(make(chan os.Signal), inmemorydb.NewInMemoryDB())
	portData := ports.NewPort{
		ID: "AEAJM",
		Port: ports.Port{
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
		},
	}

	body, err := json.Marshal(portData)
	require.NoError(t, err, "marshalling request body")

	req := httptest.NewRequest(http.MethodPost, "/v1/ports", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusCreated, res.StatusCode)

	var up ports.UpdatePort
	up.Name = ports.StringToPointerString("New Ajman")
	body, err = json.Marshal(up)
	require.NoError(t, err, "marshalling request body")

	req = httptest.NewRequest(http.MethodPatch, "/v1/ports/AEAJM", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec = httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res = rec.Result()
	defer res.Body.Close()

	expectedPortData := portData.Port
	expectedPortData.Name = "New Ajman"

	var got ports.Port
	err = json.NewDecoder(res.Body).Decode(&got)
	require.NoError(t, err, "decoding response body")

	diff := cmp.Diff(expectedPortData, got)
	require.Empty(t, diff)

	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestUpdatePort_NotExisting(t *testing.T) {
	t.Parallel()
	api := handlers.API(make(chan os.Signal), inmemorydb.NewInMemoryDB())

	var up ports.UpdatePort
	up.Name = ports.StringToPointerString("New Ajman")
	body, err := json.Marshal(up)
	require.NoError(t, err, "marshalling request body")

	req := httptest.NewRequest(http.MethodPatch, "/v1/ports/AEAJM", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusNotFound, res.StatusCode)
}

func TestGetPort_existing_and_not_existing(t *testing.T) {
	t.Parallel()
	api := handlers.API(make(chan os.Signal), inmemorydb.NewInMemoryDB())

	// not existing
	req := httptest.NewRequest(http.MethodGet, "/v1/ports/AEAJM", nil)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res := rec.Result()
	require.Equal(t, http.StatusNotFound, res.StatusCode, "getting not existing port")

	portData := ports.NewPort{
		ID: "AEAJM",
		Port: ports.Port{
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
		},
	}

	body, _ := json.Marshal(portData)
	req = httptest.NewRequest(http.MethodPost, "/v1/ports", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec = httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res = rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusCreated, res.StatusCode, "creating port")

	req = httptest.NewRequest(http.MethodGet, "/v1/ports/AEAJM", nil)
	req.Header.Set("Content-Type", "application/json")

	rec = httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res = rec.Result()

	require.Equal(t, http.StatusOK, res.StatusCode, "getting existing port")

	var got ports.Port
	err := json.NewDecoder(res.Body).Decode(&got)
	require.NoError(t, err, "decoding response body")

	diff := cmp.Diff(portData.Port, got)
	require.Empty(t, diff)
}

func TestDeletePort_existing_not_existing(t *testing.T) {
	t.Parallel()
	api := handlers.API(make(chan os.Signal), inmemorydb.NewInMemoryDB())

	// not existing
	req := httptest.NewRequest(http.MethodDelete, "/v1/ports/AEAJM", nil)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res := rec.Result()
	require.Equal(t, http.StatusNotFound, res.StatusCode, "deleting not existing port")

	portData := ports.NewPort{
		ID: "AEAJM",
		Port: ports.Port{
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
		},
	}

	body, _ := json.Marshal(portData)
	req = httptest.NewRequest(http.MethodPost, "/v1/ports", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec = httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res = rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusCreated, res.StatusCode, "creating port")

	req = httptest.NewRequest(http.MethodDelete, "/v1/ports/AEAJM", nil)
	req.Header.Set("Content-Type", "application/json")

	rec = httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res = rec.Result()

	require.Equal(t, http.StatusNoContent, res.StatusCode, "getting existing port")
}

func TestList(t *testing.T) {
	t.Parallel()

	api := handlers.API(make(chan os.Signal), inmemorydb.NewInMemoryDB())
	portData := ports.NewPort{
		ID: "AEAJM",
		Port: ports.Port{
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
		},
	}

	body, _ := json.Marshal(portData)
	req := httptest.NewRequest(http.MethodPost, "/v1/ports", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusCreated, res.StatusCode, "creating port")

	// list

	req = httptest.NewRequest(http.MethodGet, "/v1/ports", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec = httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res = rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode, "listing ports")

	var got []ports.Port
	err := json.NewDecoder(res.Body).Decode(&got)
	require.NoError(t, err, "decoding response body")
	diff := cmp.Diff([]ports.Port{portData.Port}, got)
	require.Empty(t, diff)

	// adding another port

	portdata2 := ports.NewPort{
		ID: "AEAJM_2",
		Port: ports.Port{
			Name:        "Ajman_2",
			City:        "Ajman",
			Country:     "United Arab Emirates",
			Alias:       []string{},
			Regions:     []string{},
			Coordinates: []float64{55.5136433, 25.4052165},
			Province:    "Ajman",
			Timezone:    "Asia/Dubai",
			Unlocs:      []string{"AEAJM_2"},
			Code:        "52000",
		},
	}

	body, _ = json.Marshal(portdata2)
	req = httptest.NewRequest(http.MethodPost, "/v1/ports", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec = httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res = rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusCreated, res.StatusCode, "creating port")

	// list

	req = httptest.NewRequest(http.MethodGet, "/v1/ports", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec = httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res = rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode, "listing ports")

	got = []ports.Port{}
	err = json.NewDecoder(res.Body).Decode(&got)
	require.NoError(t, err, "decoding response body")
	diff = cmp.Diff([]ports.Port{portData.Port, portdata2.Port}, got)
	require.Empty(t, diff)
}

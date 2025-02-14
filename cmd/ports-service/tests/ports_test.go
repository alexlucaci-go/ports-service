package tests

import (
	"bytes"
	"encoding/json"
	"github.com/alexlucaci-go/ports-service/cmd/ports-service/handlers"
	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreatePort(t *testing.T) {
	api := handlers.API(make(chan os.Signal))
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
	api := handlers.API(make(chan os.Signal))
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
	api := handlers.API(make(chan os.Signal))
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
	api := handlers.API(make(chan os.Signal))

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

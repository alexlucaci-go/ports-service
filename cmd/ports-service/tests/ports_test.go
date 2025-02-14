package tests

import (
	"bytes"
	"encoding/json"
	"github.com/alexlucaci-go/ports-service/cmd/ports-service/handlers"
	"github.com/alexlucaci-go/ports-service/domain/ports"
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

	body, _ := json.Marshal(portData)
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

	body, _ := json.Marshal(portData)
	req := httptest.NewRequest(http.MethodPost, "/v1/ports", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusCreated, res.StatusCode)

	var up ports.UpdatePort
	up.Name = ports.StringToPointerString("New Ajman")
	body, _ = json.Marshal(portData)
	req = httptest.NewRequest(http.MethodPut, "/v1/ports/AEAJM", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec = httptest.NewRecorder()
	api.ServeHTTP(rec, req)

	res = rec.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}

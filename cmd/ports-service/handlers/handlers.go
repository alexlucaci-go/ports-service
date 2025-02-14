package handlers

import (
	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/alexlucaci-go/ports-service/domain/ports/store/inmemorydb"
	web "github.com/alexlucaci-go/ports-service/http"
	"net/http"
	"os"
)

func API(shutdown chan os.Signal) http.Handler {
	service := web.NewService(shutdown)
	ph := portsHandler{ports.NewDomain(inmemorydb.NewInMemoryDB())}
	service.Handle(http.MethodPost, "/v1/ports", ph.Create)
	service.Handle(http.MethodPatch, "/v1/ports/{id}", ph.Update)

	return service
}

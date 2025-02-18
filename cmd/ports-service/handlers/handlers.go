package handlers

import (
	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/alexlucaci-go/ports-service/domain/ports/store/inmemorydb"
	"github.com/alexlucaci-go/ports-service/web"
	"net/http"
	"os"
)

func API(shutdown chan os.Signal, db *inmemorydb.InMemoryDB) http.Handler {
	service := web.NewService(shutdown)
	ph := portsHandler{ports.NewDomain(db)}
	service.Handle(http.MethodPost, "/v1/ports", ph.Create)
	service.Handle(http.MethodPatch, "/v1/ports/{id}", ph.Update)
	service.Handle(http.MethodGet, "/v1/ports/{id}", ph.Get)
	service.Handle(http.MethodDelete, "/v1/ports/{id}", ph.Delete)
	service.Handle(http.MethodGet, "/v1/ports", ph.List)

	return service
}

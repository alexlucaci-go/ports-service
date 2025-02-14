package handlers

import (
	"context"
	"github.com/alexlucaci-go/ports-service/domain/ports"
	"net/http"
)

type portsHandler struct {
	domain *ports.Domain
}

func (ph *portsHandler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (ph *portsHandler) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

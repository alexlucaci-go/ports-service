package handlers

import (
	"context"
	"github.com/alexlucaci-go/ports-service/domain/ports"
	web "github.com/alexlucaci-go/ports-service/http"
	"github.com/pkg/errors"
	"net/http"
)

type portsHandler struct {
	domain *ports.Domain
}

func (ph *portsHandler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var np ports.NewPort
	if err := web.Decode(r, &np); err != nil {
		return err
	}

	err := ph.domain.Create(ctx, np)
	if err != nil {
		switch {
		case errors.As(err, &ports.ErrAlreadyExists):
			return web.Respond(ctx, w, nil, http.StatusConflict)
		default:
			return web.RespondError(ctx, w, errors.Wrap(err, "creating port"))
		}
	}

	return web.Respond(ctx, w, nil, http.StatusCreated)
}

func (ph *portsHandler) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

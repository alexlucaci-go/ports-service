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
			return web.NewRequestError(errors.New("a resource with provided id already exists"), http.StatusConflict)
		default:
			return web.RespondError(ctx, w, errors.Wrap(err, "creating port"))
		}
	}

	return web.Respond(ctx, w, np.Port, http.StatusCreated)
}

func (ph *portsHandler) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := web.Param(r, "id")
	if id == "" {
		return web.NewRequestError(errors.New("id path param is required"), http.StatusBadRequest)
	}

	var up ports.UpdatePort
	if err := web.Decode(r, &up); err != nil {
		return err
	}

	err := ph.domain.Update(ctx, id, up)
	if err != nil {
		switch {
		case errors.As(err, &ports.ErrNotFound):
			return web.NewRequestError(errors.New("resource with provided id is not found"), http.StatusNotFound)
		default:
			return web.RespondError(ctx, w, errors.Wrap(err, "updating port"))
		}
	}

	updatedPort, err := ph.domain.Get(ctx, id)
	if err != nil {
		return web.RespondError(ctx, w, errors.Wrap(err, "getting updated port"))
	}

	return web.Respond(ctx, w, updatedPort, http.StatusOK)
}

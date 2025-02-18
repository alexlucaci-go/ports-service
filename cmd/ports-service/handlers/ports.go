package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/alexlucaci-go/ports-service/web"
	"net/http"
)

const fixedListLimit = 5

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
		case errors.Is(err, ports.ErrAlreadyExists):
			return web.NewRequestError(errors.New("a resource with provided id already exists"), http.StatusConflict)
		default:
			return web.RespondError(ctx, w, fmt.Errorf("creating port: %w", err))
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
		case errors.Is(err, ports.ErrNotFound):
			return web.NewRequestError(errors.New("resource with provided id is not found"), http.StatusNotFound)
		default:
			return web.RespondError(ctx, w, fmt.Errorf("updating port: %w", err))
		}
	}

	updatedPort, err := ph.domain.Get(ctx, id)
	if err != nil {
		return web.RespondError(ctx, w, fmt.Errorf("getting updated port: %w", err))
	}

	return web.Respond(ctx, w, updatedPort, http.StatusOK)
}

func (ph *portsHandler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := web.Param(r, "id")
	if id == "" {
		return web.NewRequestError(errors.New("id path param is required"), http.StatusBadRequest)
	}

	port, err := ph.domain.Get(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, ports.ErrNotFound):
			return web.NewRequestError(errors.New("resource with provided id is not found"), http.StatusNotFound)
		default:
			return web.RespondError(ctx, w, fmt.Errorf("getting port: %w", err))
		}
	}

	return web.Respond(ctx, w, port, http.StatusOK)
}

func (ph *portsHandler) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := web.Param(r, "id")
	if id == "" {
		return web.NewRequestError(errors.New("id path param is required"), http.StatusBadRequest)
	}

	err := ph.domain.Delete(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, ports.ErrNotFound):
			return web.NewRequestError(errors.New("resource with provided id is not found"), http.StatusNotFound)
		default:
			return web.RespondError(ctx, w, fmt.Errorf("deleting port: %w", err))
		}
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

func (ph *portsHandler) List(ctx context.Context, w http.ResponseWriter, _ *http.Request) error {
	// will use a hardcoded limit for demo purposes
	listedPorts, err := ph.domain.List(ctx, fixedListLimit)
	if err != nil {
		return web.RespondError(ctx, w, fmt.Errorf("listing ports: %w", err))
	}

	return web.Respond(ctx, w, listedPorts, http.StatusOK)
}

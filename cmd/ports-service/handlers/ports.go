package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/alexlucaci-go/ports-service/models"
	"github.com/alexlucaci-go/ports-service/web"
)

const fixedListLimit = 5

type portsHandler struct {
	domain ports.Service
}

func (ph portsHandler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var cp models.CreatePort
	if err := web.Decode(r, &cp); err != nil {
		return err
	}

	err := ph.domain.Create(ctx, cp)
	if err != nil {
		switch {
		case errors.Is(err, ports.ErrAlreadyExists):
			return web.NewRequestError(errors.New("a resource with provided id already exists"), http.StatusConflict)
		default:
			return web.RespondError(w, fmt.Errorf("creating port: %w", err))
		}
	}

	return web.Respond(w, cp, http.StatusCreated)
}

func (ph portsHandler) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := web.Param(r, "id")
	if id == "" {
		return web.NewRequestError(errors.New("id path param is required"), http.StatusBadRequest)
	}

	var up models.UpdatePort
	if err := web.Decode(r, &up); err != nil {
		return err
	}

	err := ph.domain.Update(ctx, id, up)
	if err != nil {
		switch {
		case errors.Is(err, ports.ErrNotFound):
			return web.NewRequestError(errors.New("resource with provided id is not found"), http.StatusNotFound)
		default:
			return web.RespondError(w, fmt.Errorf("updating port: %w", err))
		}
	}

	updatedPort, err := ph.domain.Get(ctx, id)
	if err != nil {
		return web.RespondError(w, fmt.Errorf("getting updated port: %w", err))
	}

	return web.Respond(w, updatedPort, http.StatusOK)
}

func (ph portsHandler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
			return web.RespondError(w, fmt.Errorf("getting port: %w", err))
		}
	}

	return web.Respond(w, port, http.StatusOK)
}

func (ph portsHandler) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
			return web.RespondError(w, fmt.Errorf("deleting port: %w", err))
		}
	}

	return web.Respond(w, nil, http.StatusNoContent)
}

func (ph portsHandler) List(ctx context.Context, w http.ResponseWriter, _ *http.Request) error {
	// will use a hardcoded limit for demo purposes
	listedPorts, err := ph.domain.List(ctx, fixedListLimit)
	if err != nil {
		return web.RespondError(w, fmt.Errorf("listing ports: %w", err))
	}

	return web.Respond(w, listedPorts, http.StatusOK)
}

package ports

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexlucaci-go/ports-service/entities"
	"github.com/alexlucaci-go/ports-service/models"
)

var ErrNotFound = errors.New("store resource not found")
var ErrAlreadyExists = errors.New("store resource already exists")
var ErrNoCoordinates = errors.New("both coordinates are required")
var ErrIncorrectLatitudeOrLongitudeValues = errors.New(
	"incorrect latitude or longitude. latitude should range from" +
		" -90 to 90 and longitude from -180 to 180")

type Storer interface {
	Create(ctx context.Context, port entities.Port) error
	Update(ctx context.Context, port entities.Port) error
	Get(ctx context.Context, id string) (entities.Port, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit int) ([]entities.Port, error)
}

type Service struct {
	store Storer
}

func NewService(store Storer) Service {
	return Service{store: store}
}

//nolint:gocritic // it is intentionally passed as value as I don't want any unnecessary pointers to be passed
func (d Service) Create(ctx context.Context, cp models.CreatePort) error {
	// Do some domain logic here like adding creation date or some other business logic checks
	if len(cp.Coordinates) != 2 {
		return ErrNoCoordinates
	}

	if cp.Coordinates[0] < -180 || cp.Coordinates[0] > 180 || cp.Coordinates[1] < -90 || cp.Coordinates[1] > 90 {
		return ErrIncorrectLatitudeOrLongitudeValues
	}

	err := d.store.Create(ctx, entities.Port(cp))
	if err != nil {
		return fmt.Errorf("calling store create: %w", err)
	}

	return nil
}

//nolint:gocritic // it is intentionally passed as value as I don't want any unnecessary pointers to be passed
func (d Service) Update(ctx context.Context, id string, up models.UpdatePort) error {
	// Do some domain logic here like adding update date or some other business logic checks
	port, err := d.store.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("calling store get: %w", err)
	}

	if up.Coordinates != nil {
		coords := *up.Coordinates
		if len(coords) != 2 {
			return ErrNoCoordinates
		}

		if coords[0] < -180 || coords[0] > 180 || coords[1] < -90 || coords[1] > 90 {
			return ErrIncorrectLatitudeOrLongitudeValues
		}
	}

	if up.Name != nil {
		port.Name = *up.Name
	}
	if up.City != nil {
		port.City = *up.City
	}
	if up.Country != nil {
		port.Country = *up.Country
	}
	if up.Alias != nil {
		port.Alias = *up.Alias
	}
	if up.Regions != nil {
		port.Regions = *up.Regions
	}
	if up.Coordinates != nil {
		port.Coordinates = *up.Coordinates
	}
	if up.Province != nil {
		port.Province = *up.Province
	}
	if up.Timezone != nil {
		port.Timezone = *up.Timezone
	}
	if up.Unlocs != nil {
		port.Unlocs = *up.Unlocs
	}
	if up.Code != nil {
		port.Code = *up.Code
	}

	err = d.store.Update(ctx, port)
	if err != nil {
		return fmt.Errorf("calling store update: %w", err)
	}

	return nil
}

func (d Service) Get(ctx context.Context, id string) (models.Port, error) {
	port, err := d.store.Get(ctx, id)
	if err != nil {
		return models.Port{}, fmt.Errorf("calling store get: %w", err)
	}

	return models.Port(port), nil
}

func (d Service) Delete(ctx context.Context, id string) error {
	err := d.store.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("calling store delete: %w", err)
	}

	return nil
}

func (d Service) List(ctx context.Context, limit int) ([]models.Port, error) {
	ports, err := d.store.List(ctx, limit)
	if err != nil {
		return nil, fmt.Errorf("calling store list: %w", err)
	}

	retPorts := make([]models.Port, 0, len(ports))
	for i := range ports {
		retPorts = append(retPorts, models.Port(ports[i]))
	}

	return retPorts, nil
}

package ports

import (
	"context"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("store resource not found")
var ErrAlreadyExists = errors.New("store resource already exists")
var ErrNoCoordinates = errors.New("both coordinates are required")
var ErrIncorrectLatitudeOrLongitudeValues = errors.New("incorrect latitude or longitude. latitude should range from -90 to 90 and longitude from -180 to 180")

type Storer interface {
	Create(context.Context, string, Port) error
	Update(context.Context, string, UpdatePort) error
	Get(context.Context, string) (Port, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit int) ([]Port, error)
}

type Domain struct {
	store Storer
}

func NewDomain(store Storer) *Domain {
	return &Domain{store: store}
}

func (d *Domain) Create(ctx context.Context, np NewPort) error {
	// Do some domain logic here like adding creation date or some other business logic checks
	if len(np.Port.Coordinates) != 2 {
		return ErrNoCoordinates
	}

	if np.Port.Coordinates[0] < -180 || np.Port.Coordinates[0] > 180 || np.Port.Coordinates[1] < -90 || np.Port.Coordinates[1] > 90 {
		return ErrIncorrectLatitudeOrLongitudeValues
	}

	err := d.store.Create(ctx, np.ID, np.Port)
	if err != nil {
		return errors.Wrap(err, "calling store create")
	}

	return nil
}

func (d *Domain) Update(ctx context.Context, id string, up UpdatePort) error {
	// Do some domain logic here like adding update date or some other business logic checks
	err := d.store.Update(ctx, id, up)
	if err != nil {
		return errors.Wrap(err, "calling store update")
	}

	return nil
}

func (d *Domain) Get(ctx context.Context, id string) (Port, error) {
	port, err := d.store.Get(ctx, id)
	if err != nil {
		return Port{}, errors.Wrap(err, "calling store get")
	}

	return port, nil
}

func (d *Domain) Delete(ctx context.Context, id string) error {
	err := d.store.Delete(ctx, id)
	if err != nil {
		return errors.Wrap(err, "calling store delete")
	}

	return nil
}

func (d *Domain) List(ctx context.Context, limit int) ([]Port, error) {
	ports, err := d.store.List(ctx, limit)
	if err != nil {
		return nil, errors.Wrap(err, "calling store list")
	}

	return ports, nil
}

package ports

import (
	"context"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("store resource not found")
var ErrAlreadyExists = errors.New("store resource already exists")

type Storer interface {
	Create(context.Context, string, Port) error
	Update(context.Context, string, UpdatePort) error
	Get(context.Context, string) (Port, error)
}

type Domain struct {
	store Storer
}

func NewDomain(store Storer) *Domain {
	return &Domain{store: store}
}

func (d *Domain) Create(ctx context.Context, np NewPort) error {
	// Do some domain logic here like adding creation date or some other business logic checks
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

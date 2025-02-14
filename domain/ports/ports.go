package ports

import (
	"context"
	"github.com/pkg/errors"
)

type Storer interface {
	Create(context.Context, string, Port) error
	Update(context.Context, string, Port) error
	Get(ctx context.Context, id string) (Port, error)
}

type Domain struct {
	store Storer
}

func NewDomain(store Storer) *Domain {
	return &Domain{store: store}
}

func (d *Domain) Create(ctx context.Context, id string, p Port) error {
	// Do some domain logic here like adding creation date or some other business logic checks
	err := d.store.Create(ctx, id, p)
	if err != nil {
		return errors.Wrap(err, "creating port in")
	}
	return d.store.Create(ctx, id, p)
}

func (d *Domain) Update(ctx context.Context, id string, p Port) error {
	// Do some domain logic here like adding update date or some other business logic checks
	err := d.store.Update(ctx, id, p)
	if err != nil {
		return errors.Wrap(err, "updating port in")
	}
	return d.store.Update(ctx, id, p)
}

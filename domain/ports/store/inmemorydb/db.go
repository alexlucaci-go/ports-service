package inmemorydb

import (
	"context"
	"github.com/alexlucaci-go/ports-service/domain/ports"
	"sync"
)

type InMemoryDB struct {
	data map[string]ports.Port
	mu   sync.RWMutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]ports.Port),
	}
}

func (db *InMemoryDB) Create(ctx context.Context, id string, p ports.Port) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.data[id]; ok {
		return ports.ErrAlreadyExists
	}

	db.data[id] = p
	return nil
}

func (db *InMemoryDB) Update(ctx context.Context, id string, up ports.UpdatePort) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	port, ok := db.data[id]
	if !ok {
		return ports.ErrNotFound
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

	db.data[id] = port
	return nil
}

func (db *InMemoryDB) Get(ctx context.Context, id string) (ports.Port, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	p, ok := db.data[id]
	if !ok {
		return ports.Port{}, ports.ErrNotFound
	}

	return p, nil
}

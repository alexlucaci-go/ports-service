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

func (db *InMemoryDB) Delete(ctx context.Context, id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.data[id]
	if !ok {
		return ports.ErrNotFound
	}

	delete(db.data, id)

	return nil
}

// List will list store ports; given the fact that the underlying implementation
// is using a map, subsequent calls to List using the same limit will not return the same data
// because iterating over map keys is not deterministic
func (db *InMemoryDB) List(ctx context.Context, limit int) ([]ports.Port, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if len(db.data) < limit {
		limit = len(db.data)
	}

	res := make([]ports.Port, 0, limit)
	count := 0
	for _, p := range db.data {
		if count == limit {
			break
		}
		res = append(res, p)
		count++
	}

	return res, nil
}

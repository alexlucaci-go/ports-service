package inmemorydb

import (
	"context"
	"sync"

	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/alexlucaci-go/ports-service/entities"
)

type InMemoryDB struct {
	data map[string]entities.Port
	mu   sync.RWMutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]entities.Port),
	}
}

//nolint:gocritic // it is intentionally passed as value as I don't want any unnecessary pointers to be passed
func (db *InMemoryDB) Create(_ context.Context, p entities.Port) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.data[p.ID]; ok {
		return ports.ErrAlreadyExists
	}

	db.data[p.ID] = p
	return nil
}

//nolint:gocritic // it is intentionally passed as value as I don't want any unnecessary pointers to be passed
func (db *InMemoryDB) Update(_ context.Context, p entities.Port) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.data[p.ID]
	if !ok {
		return ports.ErrNotFound
	}

	db.data[p.ID] = p
	return nil
}

func (db *InMemoryDB) Get(_ context.Context, id string) (entities.Port, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	p, ok := db.data[id]
	if !ok {
		return entities.Port{}, ports.ErrNotFound
	}

	return p, nil
}

func (db *InMemoryDB) Delete(_ context.Context, id string) error {
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
func (db *InMemoryDB) List(_ context.Context, limit int) ([]entities.Port, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if len(db.data) < limit {
		limit = len(db.data)
	}

	res := make([]entities.Port, 0, limit)
	count := 0
	for key := range db.data {
		if count == limit {
			break
		}
		res = append(res, db.data[key])
		count++
	}

	return res, nil
}

package loader

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/alexlucaci-go/ports-service/domain/ports"
)

type JSON struct {
	domain *ports.Domain
}

func NewJSON(domain *ports.Domain) *JSON {
	return &JSON{domain: domain}
}

func (l *JSON) LoadFromFile(ctx context.Context, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644) //nolint:mnd // 0644 is the default permission
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read the opening '{'
	token, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("reading opening token: %w", err)
	}

	if token != json.Delim('{') {
		return errors.New("expected opening '{'")
	}

	// the entire token decoding can probably be done more nicely, but I didn't invest
	// a lot of time in understanding how to do it properly
	for decoder.More() {
		err = l.decodeAndCreatePort(ctx, decoder)
		if err != nil {
			log.Printf("skipping port creation: %v\n", err)
			continue
		}
	}

	return nil
}

func (l *JSON) decodeAndCreatePort(ctx context.Context, decoder *json.Decoder) error {
	id, p, err := l.decodePort(decoder)
	if err != nil {
		return fmt.Errorf("decoding port: %w", err)
	}

	domainPort := ports.Port{
		ID:          id,
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}
	err = l.domain.Create(ctx, ports.NewPort{Port: domainPort})
	if err != nil {
		return fmt.Errorf("creating port with id %s: %w", id, err)
	}

	return nil
}

func (*JSON) decodePort(decoder *json.Decoder) (string, port, error) {
	var id string
	var p port
	idToken, err := decoder.Token()
	if err != nil {
		return "", port{}, fmt.Errorf("reading id token: %w", err)
	}

	id = idToken.(string)

	err = decoder.Decode(&p)
	if err != nil {
		return "", port{}, fmt.Errorf("decoding port: %w", err)
	}

	return id, p, nil
}

package loader

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexlucaci-go/ports-service/domain/ports"
	"os"
)

type Json struct {
	domain *ports.Domain
}

func NewJson(domain *ports.Domain) *Json {
	return &Json{domain: domain}
}

func (l *Json) LoadFromFile(ctx context.Context, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
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
		var id string
		var port ports.Port
		idToken, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("reading id token: %w", err)
		}

		id = idToken.(string)

		err = decoder.Decode(&port)
		if err != nil {
			return fmt.Errorf("decoding port: %w", err)
		}

		err = l.domain.Create(ctx, ports.NewPort{ID: id, Port: port})
		if err != nil {
			return fmt.Errorf("creating port: %w", err)
		}
	}

	return nil
}

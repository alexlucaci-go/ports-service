package loader

import (
	"context"
	"encoding/json"
	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/pkg/errors"
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
		return errors.Wrap(err, "opening file")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read the opening '{'
	token, err := decoder.Token()
	if err != nil {
		return errors.Wrap(err, "reading opening token")
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
			return errors.Wrap(err, "reading id token")
		}

		id = idToken.(string)

		err = decoder.Decode(&port)
		if err != nil {
			return errors.Wrap(err, "decoding port")
		}

		err = l.domain.Create(ctx, ports.NewPort{ID: id, Port: port})
		if err != nil {
			return errors.Wrap(err, "creating port")
		}
	}

	return nil
}

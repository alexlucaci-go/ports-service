package web

import (
	"encoding/json"
	"net/http"
)

func Decode(r *http.Request, val interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(val); err != nil {
		// can include here field validation logic to compose the FieldErrors that should be included in the response
		// otherwise defaults to NewRequestError with a bad request
		return NewRequestError(err, http.StatusBadRequest)
	}

	return nil
}

func Param(r *http.Request, name string) string {
	return r.PathValue(name)
}

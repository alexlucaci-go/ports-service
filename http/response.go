package web

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func Respond(ctx context.Context, w http.ResponseWriter, data interface{}, statusCode int) error {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

func RespondError(ctx context.Context, w http.ResponseWriter, err error) error {
	var res interface{}
	var code int

	switch v := errors.Cause(err).(type) {
	case *RequestError:
		res = ErrorResponse{Error: v.Err.Error()}
		code = v.Status
	case *FieldsValidationError:
		res = ErrorResponse{Error: v.Err.Error()} // should also include the fields error from the FieldErrors field
		code = http.StatusBadRequest
	default:
		code = http.StatusInternalServerError
	}

	if err := Respond(ctx, w, res, code); err != nil {
		return err
	}

	return nil
}

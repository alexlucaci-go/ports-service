package web

import (
	"context"
	"encoding/json"
	"errors"
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

	var reqError *RequestError
	var fieldsError *FieldsValidationError

	switch {
	case errors.As(err, &reqError):
		res = ErrorResponse{Error: reqError.Err.Error()}
		code = reqError.Status
	case errors.As(err, &fieldsError):
		res = ErrorResponse{Error: fieldsError.Err.Error()} // should also include the fields error from the FieldErrors field
		code = http.StatusBadRequest
	default:
		code = http.StatusInternalServerError
	}

	if err := Respond(ctx, w, res, code); err != nil {
		return err
	}

	return nil
}

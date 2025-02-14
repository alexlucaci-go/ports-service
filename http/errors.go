package web

type RequestError struct {
	Err    error
	Status int
}

func NewRequestError(err error, status int) error {
	return &RequestError{err, status}
}

func (err *RequestError) Error() string {
	return err.Err.Error()
}

type FieldsValidationError struct {
	Err         error
	FieldErrors interface{} // should include the descriptions of each field that failed validation
}

func (err *FieldsValidationError) Error() string {
	return err.Err.Error()
}

type ErrorResponse struct {
	Error string `json:"error"`
}

package exceptions

type InternalServerError struct {
	Err error
}

func NewInternalServerError(err error) *InternalServerError {

	return &InternalServerError{err}
}

func (e InternalServerError) Error() string {

	return e.Err.Error()
}

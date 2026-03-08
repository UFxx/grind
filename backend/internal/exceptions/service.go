package exceptions

type ServiceError struct {
	Err error
}

func NewServiceError(err error) *ServiceError {

	return &ServiceError{err}
}

func (e ServiceError) Error() string {

	return e.Err.Error()
}

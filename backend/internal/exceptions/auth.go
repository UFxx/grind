package exceptions

type AuthError struct {
	Err error
}

func NewAuthError(err error) *AuthError {

	return &AuthError{err}
}

func (e AuthError) Error() string {

	return e.Err.Error()
}

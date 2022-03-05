package internal

type Error struct {
	message string
	status  string
}

// Error :nodoc:
func (i *Error) Error() string {
	return i.message
}

// StatusCode :nodoc:
func (i *Error) StatusCode() string {
	return i.status
}

// NewError :nodoc:
func NewError(status, message string) *Error {
	return &Error{
		message: message,
		status:  status,
	}
}

var (
	ErrInvalidToken   = NewError("01", "invalid token")
	ErrTokenIsMissing = NewError("02", "token is missing")
	ErrUnauthorized   = NewError("03", "unauthorized")
	ErrCacheNotFound  = NewError("04", "cache not found")
)

package errors

import "net/http"

type CustomError interface {
    error
    StatusCode() int
}
type AuthenticationError struct {
    Msg string
}

func (e *AuthenticationError) Error() string {
    return e.Msg
}

func (e *AuthenticationError) StatusCode() int {
    return http.StatusUnauthorized
}

// Add other custom errors as needed

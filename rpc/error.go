package client

import (
	"fmt"
)

// Error ..
type Error struct {
	Status     string
	StatusCode int
	Body       string
	Detail     string
}

// Error error interface
func (e *Error) Error() string {
	return fmt.Sprintf("status_code: %d, status: %s, xlog: %s, body: %s", e.StatusCode, e.Status, e.Detail, e.Body)
}

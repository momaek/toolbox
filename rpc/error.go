package client

import (
	"fmt"
)

// Error ..
type Error struct {
	Status     string
	StatusCode int
	Body       string
}

func (e *Error) Error() string {
	return fmt.Sprintf("status_code: %d, status: %s, body: %s", e.StatusCode, e.Status, e.Body)
}

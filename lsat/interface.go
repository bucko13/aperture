package lsat

import (
	"time"
)

// TimeProvider is a way to get a Now() method while
// also being able to mock the same for tests
type TimeProvider interface {
	// In most cases Now() should just return the result
	// of Go's time.Now method, but sometimes, e.g. for tests
	// it can be useful to have a mocked version of this.
	Now() time.Time
}

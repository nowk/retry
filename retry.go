package retry

import (
	"fmt"
)

// RetryFunc is the function pattern to wrap your attempts. Returning an error
// will exit immediately from any further attemps.
// Successful attempts should return a (true, nil). Returning (false, nil) will
// continue to the next attempt till all have been exhausted.
type RetryFunc func(int) (bool, error)

// AttemptsExhausted implements error and is returned when all attemps have been
// exhausted and there was no successful operation
type AttemptsExhausted struct {
	N int
}

func (a AttemptsExhausted) Error() string {
	return fmt.Sprintf("error: %d attempts exhausted", a.N)
}

// Attempt attempts the RetryFunc n number times before returning an
// AttemptsExhausted error
func Attempt(fn RetryFunc, n int) error {
	for i := 0; i < n; i++ {
		ok, err := fn(i)
		if err != nil {
			return err
		}

		if ok {
			return nil
		}
	}

	return &AttemptsExhausted{
		N: n,
	}
}

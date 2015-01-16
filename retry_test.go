package retry

import (
	"errors"
	"gopkg.in/nowk/assert.v2"
	"testing"
)

func TestReturnsAttemptsExhaustedError(t *testing.T) {
	fn := func(_ int) (bool, error) {
		return false, nil
	}

	err := Attempt(fn, 3)
	assert.TypeOf(t, "*retry.AttemptsExhausted", err)
	assert.Equal(t, "error: 3 attempts exhausted", err.Error())
}

func TestExitsAttemptsOnError(t *testing.T) {
	var n int
	fn := func(i int) (bool, error) {
		if n = i; n == 1 {
			return false, errors.New("pow!")
		}

		return false, nil
	}

	err := Attempt(fn, 3)
	assert.Equal(t, 1, n)
	assert.Equal(t, "pow!", err.Error())
}

func TestErrorTrumpsOk(t *testing.T) {
	fn := func(_ int) (bool, error) {
		return true, errors.New("bam!")
	}

	err := Attempt(fn, 3)
	assert.Equal(t, "bam!", err.Error())
}

func TestExitsAttempsOnOk(t *testing.T) {
	var n int
	fn := func(i int) (bool, error) {
		if n = i; n == 3 {
			return true, nil
		}

		return false, nil
	}

	err := Attempt(fn, 9)
	assert.Nil(t, err)
	assert.Equal(t, 3, n)
}

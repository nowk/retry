# retry

[![Build Status](https://travis-ci.org/nowk/retry.svg?branch=master)](https://travis-ci.org/nowk/retry)
[![GoDoc](https://godoc.org/gopkg.in/nowk/retry.v1?status.svg)](http://godoc.org/gopkg.in/nowk/retry.v1)

Try, try and try again.

## Install

    go get gopkg.in/nowk/retry.v1

## Example

    fn := func(i int) (bool, error) {
      if true {
        return true, nil // operation success, end
      }

      return false, nil // operation failed, try again
    }

    err := retry.Attempt(fn, 3)
    if err != nil {
      // handle error or exhausted attemps
    }

## License

MIT


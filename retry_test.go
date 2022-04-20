package dutils

import (
	"errors"
	"testing"
)

func TestRetry(t *testing.T) {
	v := 0
	r := NewWithRetry[int](10, 100)
	r.Do(func() Result[int] {
		v += 1
		if v == 10 {
			return Result[int]{
				Value: v,
			}
		}
		return Result[int]{
			Error: errors.New("fake value"),
		}
	})
	if r.Result.Value != 10 {
		t.Errorf("invalid value, must be %d ", 10)
	}
}

func TestRetryWithError(t *testing.T) {
	r := NewWithRetry[int](10, 100)
	r.Do(func() Result[int] {
		return Result[int]{
			Error: errors.New("fake value"),
		}
	})
	if r.Result.Error.Error() != "fake value" {
		t.Errorf("result must has error")
	}
}

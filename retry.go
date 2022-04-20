package dutils

import "time"

type Result[T any] struct {
	Error error
	Value T
}

type WithRetry[T any] struct {
	delay  int64
	limit  uint
	count  uint
	Result Result[T]
}

func NewWithRetry[T any](limit uint, delayMs int64) *WithRetry[T] {
	return &WithRetry[T]{
		limit: limit,
		count: 0,
		delay: delayMs,
	}
}

func (r *WithRetry[T]) Do(fn func() Result[T]) *WithRetry[T] {
	for r.count < r.limit {
		r.count += 1
		r.Result = fn()
		if r.Result.Error == nil {
			return r
		}
		if r.delay != 0 {
			<-time.After(time.Duration(r.delay) * time.Millisecond)
		}
	}
	return r
}

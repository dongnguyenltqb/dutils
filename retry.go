package dutils

type Result[T any] struct {
	Error error
	Value T
}

type WithRetry[T any] struct {
	limit  uint
	count  uint
	Result Result[T]
}

func NewWithRetry[T any](limit uint) *WithRetry[T] {
	return &WithRetry[T]{
		limit: limit,
		count: 0,
	}
}

func (r *WithRetry[T]) Do(fn func() Result[T]) *WithRetry[T] {
	for r.count < r.limit {
		r.count += 1
		r.Result = fn()
		if r.Result.Error == nil {
			return r
		}
	}
	return r
}

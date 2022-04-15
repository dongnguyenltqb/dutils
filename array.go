package dutils

import (
	"math"
	"math/rand"
	"time"
)

type Heap[T any] struct {
	leng    int
	value   []T
	compare func(a, b T) int
}

func (h *Heap[T]) Insert(x T) *Heap[T] {
	h.leng++
	h.value[h.leng] = x
	h.UpHeap(h.leng)
	return h
}

func (h *Heap[T]) UpHeap(i int) *Heap[T] {
	// parent node
	k := int(math.Floor(float64(i) / 2))

	if i == 1 || h.compare(h.value[k], h.value[i]) <= 0 {
		return h
	}
	t := h.value[i]
	h.value[i] = h.value[k]
	h.value[k] = t
	h.UpHeap(k)
	return h
}

func (h *Heap[T]) DownHeap(i int) *Heap[T] {
	m := i * 2
	if m > h.leng {
		return h
	}
	if h.compare(h.value[m], h.value[m+1]) > 0 {
		m++
	}
	if h.compare(h.value[m], h.value[i]) < 0 {

		t := h.value[m]
		h.value[m] = h.value[i]
		h.value[i] = t
		h.DownHeap(m)
		return h

	}
	return h
}

func (h *Heap[T]) RemoveRoot() *Heap[T] {
	h.value[1] = h.value[h.leng]
	h.leng--
	if h.leng > 1 {
		h.DownHeap(1)
	}
	return h
}

func RandomInt(n int) []int {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = generator.Int()
	}
	return result
}

// Sort an array or slice
// compare function return -1 if a < b, 0 if a = b and 1 if a > b
func Sort[T comparable](input []T, compare func(a T, b T) int) []T {
	n := len(input)
	h := new(Heap[T])
	h.compare = compare
	h.value = make([]T, n+1)
	for i := 0; i < n; i++ {
		h.Insert(input[i])
	}
	result := make([]T, 0)
	for i := 0; i < n; i++ {
		result = append(result, h.value[1])
		h.RemoveRoot()
	}
	return result
}

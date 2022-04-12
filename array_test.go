package dutils

import (
	"testing"
)

func TestSortAsc(t *testing.T) {
	array := RandomInt(10000)
	array = Sort(array, func(a, b int) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}

		return -1

	})
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			t.Errorf("Sort result wrong : %v", array)
		}
	}
}

func TestSortDesc(t *testing.T) {
	array := RandomInt(10000)
	array = Sort(array, func(a, b int) int {
		if a == b {
			return 0
		}
		if a > b {
			return -1
		}

		return 1

	})
	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			t.Errorf("Sort result wrong : %v", array)
		}
	}
}

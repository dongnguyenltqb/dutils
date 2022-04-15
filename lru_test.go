package dutils

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	l := NewLRU(5)
	for i := 1; i <= 5; i++ {
		l.Inc(fmt.Sprintf("key%v", i))
	}
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 100-i; j++ {
			l.Inc(fmt.Sprintf("key%v", i))
		}
	}
	for i := 1; i <= 100; i++ {
		l.Inc("key6")
	}
	if l.Least() != "key4" {
		t.Error("Least recent use key must be key4")
	}
}

func BenchmarkLRU(b *testing.B) {
	nMax := uint(b.N * 2)
	l := NewLRU(nMax)

	for i := 1; i <= int(nMax); i++ {
		for j := 1; j <= 100; j++ {
			l.Inc(fmt.Sprintf("key%v", i))
		}
	}
}

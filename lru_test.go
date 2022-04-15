package dutils

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	l := NewLRU(15)
	for i := 1; i <= 10; i++ {
		l.Inc(fmt.Sprintf("key%v", i))
	}
	for i := 1; i <= 4; i++ {
		l.Inc(fmt.Sprintf("key%v", i))
		for j := 1; j <= i; j++ {
			l.Inc(fmt.Sprintf("key%v", i))
		}
	}
	for i := 6; i <= 10; i++ {
		l.Inc(fmt.Sprintf("key%v", i))
		for j := 1; j <= i; j++ {
			l.Inc(fmt.Sprintf("key%v", i))
		}
	}
	if l.Least() != "key5" {
		t.Errorf("Wrong result, must be key5")
	}
}

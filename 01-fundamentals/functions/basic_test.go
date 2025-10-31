package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(2, 3)
	if total != 5 {
		t.Errorf("Sum was incorrect got:%d,instead of:5", total)
	}
}

package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	attempt := fibonacci(5)
	if attempt != 5 {
		t.Errorf("expected %d got %d", 5, attempt)
	}

	attempt = fibonacci(8)
	if attempt != 21 {
		t.Errorf("expected %d got %d", 21, attempt)
	}

	attempt = fibonacci(15)
	if attempt != 610 {
		t.Errorf("expected %d got %d", 610, attempt)
	}
}
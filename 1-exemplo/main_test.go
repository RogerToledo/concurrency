package main

import (
	"sync"
	"testing"
)

func TestCountEven(t *testing.T) {
	var even int
	var wg sync.WaitGroup

	numbers := []int{1, 2, 3, 4}

	wg.Add(1)

	go func() {
		even = countEven(numbers, &wg)
	}()

	wg.Wait()

	if even != 2 {
		t.Errorf("Expected 2, got %d", even)
	}
}

func TestCountOdd(t *testing.T) {
	var odd int
	var wg sync.WaitGroup

	numbers := []int{1, 2, 3, 4, 5}

	wg.Add(1)

	go func() {
		odd = countOdd(numbers, &wg)
	}()

	wg.Wait()

	if odd != 3 {
		t.Errorf("Expected 2, got %d", odd)
	}
}

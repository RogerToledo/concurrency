package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var even int
	var odd int

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg.Add(2)

	go func(){
		even = countEven(numbers, &wg)
	}()

	go func() {
		odd = countOdd(numbers, &wg)
	}()

	wg.Wait()

	fmt.Printf("Even: %d - Odd: %d", even, odd)

}

func countOdd(numbers []int, wg *sync.WaitGroup) int {
	defer wg.Done()
	var odd int

	for _, number := range numbers {
		if number % 2 != 0 {
			odd++
		}
	}

	return odd
}

func countEven(numbers []int, wg *sync.WaitGroup) int {
	defer wg.Done()
	var even int

	for _, number := range numbers {
		if number % 2 == 0 {
			even++
		}
	}

	return even
}

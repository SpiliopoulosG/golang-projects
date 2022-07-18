package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting...")
}

func fibonacci(number int) int {
	if number == 0 {
		return 1
	}

	if number == 1 {
		return 1
	}
	var fibonacci_numbers = []int{1, 1}
	for num := 0; num < number-2; num++ {
		digiToAdd := fibonacci_numbers[len(fibonacci_numbers)-1] + fibonacci_numbers[len(fibonacci_numbers)-2]
		fibonacci_numbers = append(fibonacci_numbers, digiToAdd)
	}

	return fibonacci_numbers[len(fibonacci_numbers)-1]
}

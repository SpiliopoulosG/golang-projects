package main

import "fmt"

func main() {

	fmt.Println("Enter Your Number: ")

	// var then variable name then variable type
	var number int

	// Taking input from user
	fmt.Scanln(&number)

	for i := 1; i < number; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	fmt.Println("Program Finished")
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to the animal pub calculator")

	fmt.Println("What's the animal of your choice")
	var animal string
	fmt.Scanln(&animal)

	fmt.Println("How many do we get from a single birth")
	var babies_expected int
	fmt.Scanln(&babies_expected)

	fmt.Printf("We are expecting %v %s\n", babies_expected, animal)
}

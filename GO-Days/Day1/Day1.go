package main

import "fmt"

func main() {
	fmt.Println("Welcome to minecraft world name genarator!")

	fmt.Println("Enter your favorite fruit.")
	var fruit string
	fmt.Scanln(&fruit)
	// fmt.Printf("The city you were born is: %s\n", city)

	fmt.Println("What's your favorite color?")
	var color string
	fmt.Scanln(&color)

	final_name := fmt.Sprintf("%s%s", color, fruit)

	fmt.Printf("Your minecraft world name is: %v\n", final_name)
}

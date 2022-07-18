package main

import (
	"fmt"
	"strings"
)


func main() {

	// Printing
	fmt.Println(logo)
	fmt.Println("Welcome to Secret Auction")

	// Starting Constants
	var listOfPlayers []string
	var playerValues []int
	var endOfGame bool
	var currentWinner string
	var currentPrice int
	var maxPrice int

	// Starting Loop
	for endOfGame == false {
		var name string  
		fmt.Println("What's the name of the person?")
		fmt.Scanln(&name)
		listOfPlayers = append(listOfPlayers, name)
		
		var bid int
		fmt.Println("What is the bid amount?")
		fmt.Scanln(&bid)
		playerValues = append(playerValues, bid)
		
		var playerDecesion string
		fmt.Println("Is there another player? Y or N?")
		fmt.Scanln(&playerDecesion)

		if strings.ToLower(playerDecesion) == "y" {
			for index := 0; index < len(playerValues); index++ {
				currentPrice = playerValues[len(playerValues) - 1 ]
				if maxPrice < currentPrice {
					maxPrice = currentPrice
					currentWinner = listOfPlayers[len(playerValues) - 1]
				} else {
					continue
				}	
			}
		} else if strings.ToLower(playerDecesion) == "n" {
			endOfGame = true
			break
		} 
	}
	
	// Prints Winner
	fmt.Println("The winner is", currentWinner, "with the winning bid of", maxPrice)

}
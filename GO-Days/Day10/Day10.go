package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Welcome to Head or Tails Game!")

	for {
		var bet_amount int
		fmt.Println("What is your bet ammount?")
		fmt.Scanln(&bet_amount)

		if bet_amount < 0 {
			fmt.Println("You have to choose a value greater than 0!")
			continue
		}

	}

		var player_choice string
		fmt.Println("Head or Tails:")
		fmt.Scanln(&player_choice)

		if strings.ToLower(player_choice) != "head" && strings.ToLower(player_choice) != "tail"{
			fmt.Println("You must choose between Head or Tail!")
		}


}

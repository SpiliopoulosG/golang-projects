package main

import (
  "fmt"
  "math/rand"
  "time"
  "strings"
)

func main() {

	rand.Seed(time.Now().Unix())

	// ASCII Art

	var rock = `
		_______
	---'   ____)
		  (_____)
		  (_____)
		  (____)
	---.__(___)
	`
	var paper = `
		_______
	---'   ____)____
			  ______)
			  _______)
			 _______)
	---.__________)
	`

	var scissors = `
		_______
	---'   ____)____
			  ______)
		  __________)
		  (____)
	---.__(___)
	`

	pool := [3]string{"rock", "paper", "scissors"}
	fmt.Println("What do you choose? Type rock, paper or scissors.")
	computer_choice := fmt.Sprint(pool[rand.Intn(len(pool))])


	var player_choice string 
	fmt.Scanln(&player_choice)

	//Game Logic
	if strings.ToLower(player_choice) == "rock"  {
		fmt.Println(rock)
		if computer_choice == "rock" {
			fmt.Println("Computer chose:\n" + rock + "\nIt's a Tie\n")
		} else if computer_choice == "paper" {
			fmt.Println("Computer chose:\n" + paper  + "\nYou lose")
		} else if computer_choice == "scissors" {
			fmt.Println("Computer chose:\n" + scissors + "\nYou win")
		}


	} else if strings.ToLower(player_choice) == "paper"  {
		fmt.Println(paper)
		if computer_choice == "rock" {
			fmt.Println("Computer chose:\n" + rock + "\nYou win\n")
		} else if computer_choice == "paper" {
			fmt.Println("Computer chose:\n" + paper  + "\nIt's a tie")
		} else if computer_choice == "scissors" {
			fmt.Println("Computer chose:\n" + scissors + "\nYou lose")
		}


	} else if strings.ToLower(player_choice) == "scissors"  {
		fmt.Println(scissors)
		if computer_choice == "rock" {
			fmt.Println("Computer chose:\n" + rock + "\nYou lose\n")
		} else if computer_choice == "paper" {
			fmt.Println("Computer chose:\n" + paper  + "\nYou win")
		} else if computer_choice == "scissors" {
			fmt.Println("Computer chose:\n" + scissors + "\nIt's a tie")
		}


	}
}




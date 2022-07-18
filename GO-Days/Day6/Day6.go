package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func stringInSlice(character string, word []string) bool {
			for _, position := range word {
				if position == character {
					return true
				}
			}
			return false
		}

func stringInString(character string, word string) bool {
			for position := 0; position < len(word); position++ {
				if string(word[position]) == character {
					return true
				}
			}
			return false
		}


func main() {

	rand.Seed(time.Now().Unix())

	// HangMan

	fmt.Println("Starting... HangMan")

	// Word Selection
	var chosenWord string = wordGenerator()
	var wordLength int = len(chosenWord)

	// Starting Constants
	var lives int = 6

	fmt.Println("Pssst, the solution is:", chosenWord)

	// Create blanks
	var display []string
	for i := 0; i < wordLength; i++ {
		display = append(display, "_")
	}

	fmt.Println(display)

	// Start of Loop
	for lives > 0 {
		var guess string 
		fmt.Println("Guess a letter: ")
		fmt.Scanln(&guess)


		var guessList []string
	
		for position := 0; position < len(display); position++{
			if display[position] == guess {
				fmt.Println("You have already entered", guess ,", please enter another one")
			}
		}

		// Checks guessed letter
		for position := 0; position < wordLength; position++ {
			var letter string = string(chosenWord[position])
			// fmt.Println("Current position:", position, "Current letter:", letter, "Guessed letter:", guess)
			if letter == guess{
				display[position] = letter
			}
		}


		// Checks if user is wrong.
		if  stringInString(guess, chosenWord) == false && stringInSlice(guess, guessList) == false {
			if stringInSlice(guess, guessList) == false {
				guessList = append(guessList, guess)
				fmt.Println("The letter", guess, "is not in the word")
				lives -= 1
			}

			if lives == 0 {
				fmt.Println("You lose.\nThe word was.", chosenWord)
				break
			}

		}
			
		// Checks if user has got all letters.
		if stringInSlice("_", display) == false{
			print("You win.")
			break
		}

		fmt.Println(stages[lives])

		// Join all the elements in the list and turn it into a String.
		fmt.Println(strings.Join(display, " "))
		}


}

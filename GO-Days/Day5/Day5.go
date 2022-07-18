package main 

import (
	"fmt"
	"math/rand"
  	"time"
	"strings"
)

func main() {

	rand.Seed(time.Now().Unix())

	letters := [60]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r",
	  "s", "t", "u", "v", "w", "x", "y","z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
	  "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	numbers := [10]string{ "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	symbols := [10]string{"!", "#", "$", "%", "&", "(", ")", "*", "+" }


	fmt.Println("Welcome to the GoPassword Generator!")

	var lettersToAdd, symbolsToAdd, numbersToAdd int 
	var shufflePassword string

	fmt.Println("How many letters would you like in your password?")
	fmt.Scanln(&lettersToAdd)

	fmt.Println("How many symbols would you like?")
	fmt.Scanln(&symbolsToAdd)

	fmt.Println("How many numbers would you like?")
	fmt.Scanln(&numbersToAdd)

	fmt.Println("Would you like to shuffle your password. Y/N")
	fmt.Scanln(&shufflePassword)

	fmt.Println(shufflePassword)

	var password string 

	fmt.Println(password)

	for i := 0; i < lettersToAdd; i++ {
			randomLetter := fmt.Sprint(letters[rand.Intn(len(letters))]) 
			password += randomLetter
		}

	for i := 0; i < numbersToAdd; i++ {
			randomNumber := fmt.Sprint(numbers[rand.Intn(len(numbers))]) 
			password += randomNumber
		}

	for i := 0; i < symbolsToAdd; i++ {
			randomSymbol := fmt.Sprint(symbols[rand.Intn(len(symbols))]) 
			password += randomSymbol
		}	
	
	if strings.ToLower(shufflePassword) == "y" {
		passwordRune := []rune(password)
		rand.Shuffle(len(passwordRune), func(i, j int) { passwordRune[i], passwordRune[j] = passwordRune[j], passwordRune[i] })
		rand.Shuffle(len(passwordRune), func(i, j int) { passwordRune[i], passwordRune[j] = passwordRune[j], passwordRune[i] })
		fmt.Println("Your password is: ", string(passwordRune))
	}  else {
		fmt.Println("Your password is: ", password)
	}

}
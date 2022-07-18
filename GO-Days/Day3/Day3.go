package main

import "fmt"


func main() {
	
	fmt.Println("Welcome to Fairy Tale Game!")
	
	fmt.Println("The story of yours start now. You are in a post apocalyptic world travelling through countryside")

	fmt.Println("Suddenly in front of you, you see the path splitting one road leads to a dark cave and the other to a deep lake. Where do you go, cave or lake?")
	var answer string 
	fmt.Scanln(&answer)
	
	if answer == "cave" {
		fmt.Printf(`You proceed to the cave through a small hole in the front and you see a huge ravin ahead of you full of mushrooms and valueable jewels.
					You want to explore to continue finding the exit or gather jewels. Jewels or Exit`)
		fmt.Scanln(&answer)		
		if answer == "exit" {
			fmt.Println("You continue to the exit and you finally manage to leave the cave. Congrats you are out of the cave and happy to continue adventuring.")
		} else {
			fmt.Println("You start gathering jewels and before you understand a huge shadows hides the light from you and within seconds a humongous dragon burns you to crisp!")
		}
	} else if answer == "lake" {
		fmt.Printf(`You proceed to the lake  in the front and you see a huge open space in front of you. You hear songs from afar.
					 You want to explore to continue finding the exit or follow the songs. Songs or Exit`)
		fmt.Scanln(&answer)		
		if answer == "exit" {
			fmt.Println("You continue to the exit and you finally manage to leave the lake. Congrats you are out of the lake and happy to continue adventuring.")
		} else {
			fmt.Println("You start following the songs and you are drawn to an extremely beautiful woman. Within seconds she eats you alive!")
		}
	}

}


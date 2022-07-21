package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type joke struct {
	Error    bool   `json:"error"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Joke     string `json:"joke"`
	Flags    struct {
		Nsfw      bool `json:"nsfw"`
		Religious bool `json:"religious"`
		Political bool `json:"political"`
		Racist    bool `json:"racist"`
		Sexist    bool `json:"sexist"`
		Explicit  bool `json:"explicit"`
	} `json:"flags"`
	ID   int    `json:"id"`
	Safe bool   `json:"safe"`
	Lang string `json:"lang"`
}

func main() {

	fmt.Print("\n")
	fmt.Println("*********************\nWelcome to JokeWorld!\n*********************")
	fmt.Print("\n")

	options := []string{ "Programming", "Miscellaneous", "Dark", "Pun", "Spooky", "Christmas", "Any" }
	var optionsChoice int
	fmt.Println("Choose a joke category:\n-1 for Programming\n-2 for Misc\n-3 for Dark\n-4 for Pun\n-5 for Spooky\n-6 for Christmas\n-7 for Any")
	fmt.Print("\nInput: ")
	fmt.Scanln(&optionsChoice)
	
	

	url := "https://v2.jokeapi.dev/joke/" + options[optionsChoice - 1]

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	req.Header.Add("X-RapidAPI-Key", "6d9f285875mshfec9aedce5024a6p113401jsn61f3433754f0")
	req.Header.Add("X-RapidAPI-Host", "joke3.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error", err)
	}
	

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error", err)
	}

	var result joke
	if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Can not unmarshal JSON")
    }

	fmt.Println(result.Joke)

}
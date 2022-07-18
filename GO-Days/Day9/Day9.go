package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Joke struct {
	Categories []interface{} `json:"categories"`
	CreatedAt  string        `json:"created_at"`
	IconURL    string        `json:"icon_url"`
	ID         string        `json:"id"`
	UpdatedAt  string        `json:"updated_at"`
	URL        string        `json:"url"`
	Value      string        `json:"value"`
}

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	homePageVariables := getjoke()
	tpl.Execute(w, homePageVariables)
}

func main() {
	

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)

	fmt.Println("Server startedd")

}

func getjoke() Joke {

	url := "https://matchilling-chuck-norris-jokes-v1.p.rapidapi.com/jokes/random"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-RapidAPI-Key", "6d9f285875mshfec9aedce5024a6p113401jsn61f3433754f0")
	req.Header.Add("X-RapidAPI-Host", "matchilling-chuck-norris-jokes-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result Joke
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Can not unmarshal JSON")
    }

	return result

}
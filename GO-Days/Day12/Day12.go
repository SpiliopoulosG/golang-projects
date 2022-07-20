package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
)

type currencies []string

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	homePageVariables := getexchanges()
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

func getexchanges() currencies {

	url := "https://currency-exchange.p.rapidapi.com/listquotes"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "6d9f285875mshfec9aedce5024a6p113401jsn61f3433754f0")
	req.Header.Add("X-RapidAPI-Host", "currency-exchange.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	
	var result currencies
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Can not unmarshal JSON")
    }

	return result

}

func exchange() {

	var amount int
	var from_currency string
	var to_currency string

	url := fmt.Sprintf("https://currency-exchange.p.rapidapi.com/exchange?from=", from_currency, "&to=", to_currency, "&q=", amount)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "6d9f285875mshfec9aedce5024a6p113401jsn61f3433754f0")
	req.Header.Add("X-RapidAPI-Host", "currency-exchange.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}
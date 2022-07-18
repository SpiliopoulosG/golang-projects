package main

import (
   "io/ioutil"
   "net/http"
   "fmt"
   "encoding/json"
)

type Love struct {
	Fname      string `json:"fname"`
	Sname      string `json:"sname"`
	Percentage string `json:"percentage"`
	Result     string `json:"result"`
}

func main() {

	var fName string
	fmt.Println("What's the name of the first person?")
	fmt.Scanln(&fName)

	var sName string
	fmt.Println("What's the name of the second person?")
	fmt.Scanln(&sName)

	var url string = "https://love-calculator.p.rapidapi.com/"
	var query string = fmt.Sprintf("getPercentage?fname=%v&sname=%v", fName, sName)
	var finalUrl string = url + query

	req, _ := http.NewRequest("GET", finalUrl, nil)

	req.Header.Add("X-RapidAPI-Key", "6d9f285875mshfec9aedce5024a6p113401jsn61f3433754f0")
	req.Header.Add("X-RapidAPI-Host", "love-calculator.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(string(body))

	var result Love
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Can not unmarshal JSON")
    }

    fmt.Printf("%v and %v match with at a percentage of %v, %v\n", result.Fname, result.Sname, result.Percentage, result.Result)

}
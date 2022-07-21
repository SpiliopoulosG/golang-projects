package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func main() {

	fmt.Print("\n")
	fmt.Println("*********************\nCurrentWeather!\n*********************")
	fmt.Print("\n")

	var location string
	fmt.Println("Choose a location to look for weather")
	fmt.Print("\nInput: ")
	fmt.Scanln(&location)
	
	apiKey := "350503d35669a49c6bb3d94fdca9e777"

	url := "https://api.openweathermap.org/data/2.5/weather?id=524901&units=metric&appid=" + apiKey + "&q=" + location

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error", err)
	}
	

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error", err)
	}

	var result weather
	if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Can not unmarshal JSON")
    }
	fmt.Println(result.Weather[0].Main)
	fmt.Println(string(body))

}
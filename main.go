package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	loggly "github.com/jamespearly/loggly"
)

//Data structure
type Data struct {
	Cod     string  `json:"cod"`
	Message float32 `json:"message"`
	Cnt     int     `json:"cnt"`
	List    [2]List `json:"list"`
}

//LocationData structure
type LocationData struct {
	Dt          int       `json:"dt"`
	Name        string    `json:"name"`
	Coordinates Cord      `json:"coord"`
	WeatherData []Weather `json:"weather"`
	MainData    Main      `json:"main"`
	WindData    Wind      `json:"wind"`
}

//Cord structure
type Cord struct {
	Long float32 `json:"lon"`
	Lat  float32 `json:"lat"`
}

//Main struct
type Main struct {
	Temp     float32 `json:"temp"`
	Pressure float32 `json:"pressure"`
	Humidity int     `json:"humidity"`
}

//Weather structure
type Weather struct {
	Main string `json:"main"`
	Desc string `json:"description"`
}

//Wind structure
type Wind struct {
	Speed float32 `json:"speed"`
	Angle float32 `json:"deg"`
}

//List structure
type List struct {
	Dt          int       `json:"dt"`
	WeatherData []Weather `json:"weather"`
	WindData    Wind      `json:"wind"`
	TimeStamp   string    `json:"dt_txt"`
	MainData    Main      `json:"main"`
}

func main() {
	// os.Setenv("LOGGLY_TOKEN", "764e6b8e-5e82-41aa-abf6-0ce73b733b73")
	//for terminal : LOGGLY_TOKEN=764e6b8e-5e82-41aa-abf6-0ce73b733b73 go run main.go
	fmt.Println("Weather forecast by id")
	m := weatherForecast()
	fmt.Println("Weather by zipcode")
	n := weatherCurrent()
	var tag string
	// var message string
	// message = "morning"
	tag = "Weather"
	client := loggly.New(tag)
	client.EchoSend("info", m)
	client.EchoSend("info", n)
}

func weatherForecast() string {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/forecast?id=4070245&APPID=32d9771dc19c9e1910a8b88187d95573")
	if err != nil {
		return err.Error()
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		s := Data{}
		err := json.Unmarshal(data, &s)
		fmt.Println(s, err)
		return "Successful"
	}
}

func weatherCurrent() string {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=13126,US&APPID=32d9771dc19c9e1910a8b88187d95573")
	if err != nil {
		return err.Error()
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		s := LocationData{}
		err := json.Unmarshal(data, &s)
		fmt.Println(s, err)
		return "Successful"
	}
}

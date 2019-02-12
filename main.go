package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	loggly "github.com/jamespearly/loggly"
	"github.com/jasonlvhit/gocron"
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
	c := gocron.NewScheduler()
	c.Every(1).Hours().Do(weatherForecast)
	c.Every(1).Hours().Do(weatherCurrent)
	<-c.Start()
	// os.Setenv("LOGGLY_TOKEN", "764e6b8e-5e82-41aa-abf6-0ce73b733b73")
	//for terminal : LOGGLY_TOKEN=764e6b8e-5e82-41aa-abf6-0ce73b733b73 go run main.go
}

func weatherForecast() {
	var tag string
	tag = "weatherForecast"
	client := loggly.New(tag)
	fmt.Println("Weather forecast by id")
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/forecast?id=4070245&APPID=32d9771dc19c9e1910a8b88187d95573")
	if err != nil {
		client.EchoSend("info", err.Error())
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		s := Data{}
		err := json.Unmarshal(data, &s)
		fmt.Println(s, err)
		client.EchoSend("info", "success")
	}
}

func weatherCurrent() {
	var tag string
	tag = "weatherCurrent"
	client := loggly.New(tag)
	fmt.Println("Weather by zipcode")
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=13126,US&APPID=32d9771dc19c9e1910a8b88187d95573")
	if err != nil {
		client.EchoSend("info", err.Error())
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		s := LocationData{}
		err := json.Unmarshal(data, &s)
		fmt.Println(s, err)
		client.EchoSend("info", "success")
	}
}

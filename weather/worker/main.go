package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
//Item is the item data stored in DynamoDB a table in AWS
type Item struct {
	TimeInEpoch  int     `json:"TimeInEpoch"`
	Name         string  `json:"Name"`
	TempInKelvin float32 `json:"TempInKelvin"`
	Weather      string  `json:"Weather"`
	Humidity     int     `json:"Humidity"`
	WindSpeed    float32 `json:"WindSpeed"`
}

func main() {
	c := gocron.NewScheduler()
	c.Every(1).Hours().Do(weatherCharlotte)
	c.Every(1).Hours().Do(weatherOswego)
	<-c.Start()
}
//AddDBItem is a fuction for adding item to the DynamoDB table in AWS
func AddDBItem(item Item, svc *dynamodb.DynamoDB) {
	av, err := dynamodbattribute.MarshalMap(item)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Weather"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully added item to table")
}

func weatherCharlotte() {
	var tag string
	tag = "weatherCurrent"
	client := loggly.New(tag)
	fmt.Println("Weather by zipcode")
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=28025,US&APPID=32d9771dc19c9e1910a8b88187d95573")
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

func weatherOswego() {
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

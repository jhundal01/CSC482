package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Item is the item data stored in DynamoDB a table in AWS
type Item struct {
	TimeInEpoch  int     `json:"TimeInEpoch"`
	Name         string  `json:"Name"`
	TempInKelvin float32 `json:"TempInKelvin"`
	Weather      string  `json:"Weather"`
	Humidity     int     `json:"Humidity"`
	WindSpeed    float32 `json:"WindSpeed"`
}

//Status get table info
type Status struct {
	TableName   string `json:"table"`
	RecordCount int    `json:"recordCount"`
}

var status []Status

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/jhundal/all", GetDBItem).Methods("GET")
	router.HandleFunc("/jhundal/status", GetDBHC).Methods("GET")

	status = append(status, Status{TableName: "Weather", RecordCount: 200})

	logger := handlers.LoggingHandler(os.Stdout, router)

	srvr := http.Server{
		Addr:         ":8080",
		Handler:      logger,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	panic(srvr.ListenAndServe())
}

//GetDBItem prints out a message to a web server
func GetDBItem(w http.ResponseWriter, r *http.Request) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		fmt.Println("Error starting a new session")
	}
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	params := &dynamodb.ScanInput{
		TableName: aws.String("Weather"),
	}

	result, err := svc.Scan(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	item := []Item{}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Fprintf(w, "Weather Items:")
	fmt.Fprintf(w, "\n")
	fmt.Fprintln(w, item)
}

//GetDBHC prints out a message to a web server
func GetDBHC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

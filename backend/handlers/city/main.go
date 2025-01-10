package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
)

type City struct {
	Id   string  `json:"id"`
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func init() {
	_, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

}

func handleRequest(ctx context.Context, event json.RawMessage) (Response, error) {
	// Parse the input event
	var city City = City{
		Id:   "1",
		Name: "Amsterdam",
		Lat:  52.377956,
		Long: 4.897070,
	}

	content, _ := json.Marshal(city)
	return Response{
		StatusCode: 200,
		Body:       string(content),
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}

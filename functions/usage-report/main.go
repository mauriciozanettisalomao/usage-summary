package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// Endpoint TODO
type Endpoint struct {
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

// UsageSummary TODO
type UsageSummary struct {
	Date         string     `json:"date"`
	Platform     string     `json:"platform"`
	Requests     string        `json:"requests"`
	Availability string        `json:"availability"`
	Others       string        `json:"others"`
	Endpoints    []Endpoint `json:"endpoints"`
}


// Event TODO
type Event struct {
	UsageSummary UsageSummary `json:"UsageSummary"`
	DayOfWeek    string       `json:"dayOfWeek"`
	DayOfMonth   int          `json:"dayOfMonth"`
}

func handler(ctx context.Context, event Event) error {

	fmt.Println("event usage summary", event)

	return nil
}

func main() {

	lambda.Start(handler)
}

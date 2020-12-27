package main

import (
	"context"

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
	UsageSummary UsageSummary `json:"usageSummary"`
	DayOfWeek    string       `json:"dayOfWeek"`
	DayOfMonth   int          `json:"dayOfMonth"`
}

func handler(ctx context.Context) (Event, error) {

	return Event{
		DayOfWeek:  "Sunday",
		DayOfMonth: 1,
		UsageSummary: UsageSummary{
			Date:         "2020-12-27",
			Platform:     "test",
			Requests:     "10000",
			Availability: "6000",
			Others:       "4000",
			Endpoints: []Endpoint{
				{
					Name: "/test",
					Qty:  1000,
				},
				{
					Name: "/version",
					Qty:  2500,
				},
				{
					Name: "/",
					Qty:  500,
				},
			},
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}

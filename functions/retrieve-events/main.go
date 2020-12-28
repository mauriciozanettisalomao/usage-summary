package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

// UsageSummary TODO
type UsageSummary struct {
	Date         string            `json:"date"`
	Platform     string            `json:"platform"`
	Requests     string            `json:"requests"`
	Availability string            `json:"availability"`
	Others       string            `json:"others"`
	Endpoints    map[string]string `json:"endpoints"`
}

// Event TODO
type Event struct {
	UsageSummary UsageSummary `json:"usageSummary"`
	DayOfWeek    string       `json:"dayOfWeek"`
	DayOfMonth   int          `json:"dayOfMonth"`
}

// Input TODO
type Input struct {
	Platform string `json:"platform"`
	Source   string `json:"source"`
}

func handler(ctx context.Context, in Input) (Event, error) {

	return Event{
		DayOfWeek:  "Sunday",
		DayOfMonth: 1,
		UsageSummary: UsageSummary{
			Date:         "2020-12-27",
			Platform:     in.Platform,
			Requests:     "10000",
			Availability: "6000",
			Others:       "4000",
			Endpoints: map[string]string{
				"/test": "1000",
				"/":     "2500",
				"/v2":   "500",
			},
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}

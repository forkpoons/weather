package dto

import "time"

type ForecastsDb struct {
	RequestDate  string
	ForecastDate string
	Min, Max     float64
}

type Forecasts struct {
	date     time.Time
	min, max float64
}

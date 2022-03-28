package dto

import "time"

type ForecastsDb struct {
	Id           int     `db:"id"`
	RequestDate  string  `db:"requestDate"`
	ForecastDate string  `db:"forecastDate"`
	Min          float64 `db:"min"`
	Max          float64 `db:"max"`
}

type Forecasts struct {
	RequestDate  time.Time
	ForecastDate time.Time
	Min, Max     float64
}

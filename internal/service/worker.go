package service

import (
	"encoding/json"
	"github.com/forkpoons/weather/internal/dbf"
	"github.com/forkpoons/weather/internal/dto"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type data struct {
	DailyForecasts []struct {
		Date        time.Time `json:"Date"`
		Temperature struct {
			Minimum struct {
				Value float64 `json:"Value"`
			} `json:"Minimum"`
			Maximum struct {
				Value float64 `json:"Value"`
			} `json:"Maximum"`
		} `json:"Temperature"`
	} `json:"DailyForecasts"`
}

func GetFiveForecasts() {
	resp, err := http.Get("http://dataservice.accuweather.com/forecasts/v1/daily/5day/290868?apikey=3sB0cSMGeiqlMpJJpGZEZzOb7s3jax5K&language=en-us&details=false&metric=true")
	if err != nil {
		log.Fatalln(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}
	var dat data
	if err := json.Unmarshal(body, &dat); err != nil {
		log.Fatalln(err)
		return
	}
	var forecasts []dto.Forecasts
	for _, d := range dat.DailyForecasts {
		forecasts = append(forecasts, dto.Forecasts{Date: d.Date, Max: d.Temperature.Maximum.Value, Min: d.Temperature.Minimum.Value})
	}
	log.Println(forecasts)
	dbf.PostForecasts(forecasts)
}

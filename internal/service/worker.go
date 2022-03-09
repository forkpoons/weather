package service

import (
	"encoding/json"
	_ "github.com/forkpoons/weather/internal/dbf"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type DailyForecasts struct {
	date     string
	min, max float64
}

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
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//log.Println(string(body))
	var dat data
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
	var forecasts []DailyForecasts
	for _, d := range dat.DailyForecasts {
		forecasts = append(forecasts, DailyForecasts{date: d.Date.Format("2006-01-02"), max: d.Temperature.Maximum.Value, min: d.Temperature.Minimum.Value})
	}
	log.Println(forecasts)
	dbf.Qwe()
}

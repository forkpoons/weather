package service

import (
	"io/ioutil"
	"log"
	"net/http"
)

func MakeRequest() {
	resp, err := http.Get("http://dataservice.accuweather.com/forecasts/v1/daily/5day/290868?apikey=3sB0cSMGeiqlMpJJpGZEZzOb7s3jax5K&language=en-us&details=false&metric=true")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

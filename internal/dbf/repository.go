package dbf

import (
	"github.com/forkpoons/weather/internal/dto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func PostForecasts(Forecasts []dto.Forecasts) {
	db, err := sqlx.Connect("mysql", "admin:qwe123@tcp(localhost:3306)/test")
	if err != nil {
		log.Println(err)
		return
	}
	for _, f := range Forecasts {
		res, err := db.Exec(
			"INSERT INTO forecasts (requestDate, forecastDate, min, max) VALUES(?,?,?,?)",
			time.Now().Format("2006-01-02"),
			f.Date.Format("2006-01-02"),
			f.Min,
			f.Max,
		)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = res.LastInsertId()
		if err != nil {
			log.Printf("qwe")
			log.Println(err)
			return
		}
	}
	if err != nil {
		return
	}
	var forecast []dto.ForecastsDb
	err = db.Select(&forecast, "select * from forecasts")
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(forecast)
}

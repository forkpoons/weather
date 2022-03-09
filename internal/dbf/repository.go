package dbf

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Forecasts struct {
	Id       int
	Date     string
	Min, Max float64
}

func Qwe() {
	db, err := sqlx.Connect("mysql", "admin:qwe123@tcp(localhost:3306)/test")
	if err != nil {
		log.Println(err)
		return
	}
	res, err := db.Exec("INSERT INTO forecasts (date, min, max) VALUES('2022-03-09',4,5 )")
	if err != nil {
		panic(err)
	}
	_, err = res.LastInsertId()
	if err != nil {
		return
	}
	var forecast []Forecasts
	err = db.Select(&forecast, "select * from forecasts")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(forecast)
	if err != nil {
		log.Fatalln(err)
	}
}

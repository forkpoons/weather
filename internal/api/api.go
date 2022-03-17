package api

import (
	"log"
	"net/http"
)

func start() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		_, err := w.Write([]byte("hello"))
		if err != nil {
			log.Println(err)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

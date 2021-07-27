package main

import (
	"ch16/web21/app"
	"log"
	"net/http"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()
	
	log.Println("Started App")
	err := http.ListenAndServe(":4000", m)	
	if err != nil {
		panic(err)
	}
}
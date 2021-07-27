package main

import (
	"ch16/web15/app"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)
	http.ListenAndServe(":3000", n)	
}
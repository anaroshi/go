package main

import (
	"net/http"

	"restfuApi/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHandler())
}

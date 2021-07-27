package main

import (
	"net/http"

	"github.com/anaroshi/learngo/web/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}

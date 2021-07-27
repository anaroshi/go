package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, "Hello world")
		s := "Hello World"
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Hello</title>
			<script thpe="text/javascript" src="/src/web1/assets/hello.js"></script>
			<link rel="stylesheet" href="/src/web1/assets/hello.css">
		</head>
		<body>
			<stan class="hello">` + s + `</stan>
		</body>

		</html>		
		`
		res.Header().Set("Content-Type", "text/html")
		res.Write([]byte(html))
	})

	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
	http.ListenAndServe(":3000", nil)
}

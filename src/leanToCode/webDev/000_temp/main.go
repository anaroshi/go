package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	First string
	Last string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	p1 := Person{ First:"James", Last:"Bond"}
	tpl.ExecuteTemplate(w, "index.gohtml", p1)
}

func about(w http.ResponseWriter, r *http.Request) {
	xi := []int{3,5,7,9,17,749}
	tpl.ExecuteTemplate(w, "about.gohtml", xi)
}

func contact(w http.ResponseWriter, r *http.Request) {
	m := map[string]int{"james":32,"monpenny":24}
	tpl.ExecuteTemplate(w, "contact.gohtml", m)
}
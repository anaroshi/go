package main

import (
	"cnestForms/dbconn"
	"cnestForms/goq"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index(rw http.ResponseWriter, r *http.Request) {	
	tpl.ExecuteTemplate(rw, "index.gohtml", nil)
}

func goqHandler(rw http.ResponseWriter, r *http.Request) {
	goq.ExampleScrape(rw,r)
}

func getStntInfo(rw http.ResponseWriter, r *http.Request) {
	dbconn.GetInfoStudent(rw, r)
}

func processor(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(rw, r, "/", http.StatusSeeOther)
		return
	}
	fname := r.FormValue("firstName") 
	lname := r.FormValue("lastName")	 

	d := struct {
		FirstName string
		LastName string
	}{
		FirstName: fname,
		LastName: lname,
	}
	tpl.ExecuteTemplate(rw, "processor.gohtml", d)	
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/",index)
	mux.HandleFunc("/processor", processor).Methods("POST")
	mux.HandleFunc("/goq", goqHandler)
	mux.HandleFunc("/dbconn", getStntInfo)	
	http.ListenAndServe(":3000",mux)
}


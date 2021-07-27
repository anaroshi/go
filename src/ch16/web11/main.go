package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age int
	Email string
}

func (u User) IsOld() bool {
	return u.Age>30
}

func main() {
	user := User{
		Name:"Ann",
		Age:42,
		Email:"sundor@hanmail.net",
	}

	user2 := User{
		Name:"Sung",
		Age:23,
		Email:"sung@hanmail.net",
	}
	users := []User{user,user2}
	tmpl, err := template.New("Templ1").ParseFiles("templates/tmpl1.tmpl","templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(os.Stdout,"tmpl2.tmpl", users)
}
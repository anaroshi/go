package myapp

import (
	"cnestWeb/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World")
}

func barHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(rw, "Hello %s!", name)
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()

	data, err := json.Marshal(user)
	rw.Header().Add("content-type","application/json")
	utils.HandleErr(err)
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprint(rw, string(data))

}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", &fooHandler{})
	return mux	
}
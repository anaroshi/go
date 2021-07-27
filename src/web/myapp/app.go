package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hellow World")
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// json 주고 받기
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user) // json형태로 encoding함
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}

func BarHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // url에서 name 정보를 가져옴 http://localhost:3000/bar?name=sundor
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "Hellow %s!", name) // Hello sundor!
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux() // 라우터 인스턴스를 만들어 등록하는 방식

	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", BarHandler)

	mux.Handle("/foo", &fooHandler{})

	return mux
}

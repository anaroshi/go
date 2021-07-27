package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string 	`json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUser struct {
	ID int `json:"id"`
	UpdatedFirstName bool `json:"updated_first_name"`
	FirstName string `json:"first_name"`
	UpdatedLastName bool `json:"updated_last_name"`
	LastName string `json:"last_name"`
	UpdatedEmail bool `json:"updated_email"`
	Email string 	`json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var userMap map[int]*User
var lastID int

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello Ann")
}

func usersHandler(rw http.ResponseWriter, r *http.Request) {
	if len(userMap) == 0 {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprint(rw,"No Users")
		return
	}
	users := []*User{}
	for _, u := range userMap {
		users = append(users, u)
	}
	data, _ := json.Marshal(users)
	rw.Header().Add("Content-Type","applicaton/json")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, string(data))
}

func getUserInfoHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, err)
		return
	}

	user, ok := userMap[id]
	if !ok {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprint(rw, "No User Id:", id)
		return
	}
	// user := new(User)
	// user.ID = 2
	// user.FirstName="junghwa" 
	// user.LastName="sung"
	// user.Email="sundo@hanmail.net"

	rw.Header().Add("content-type","application/json")
	rw.WriteHeader(http.StatusOK)		
	data, _ := json.Marshal(user)
	fmt.Fprint(rw, string(data))
}

func createUserHandler(rw http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "Bad Request: ", err)
		return
	}

	// Created User
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	rw.Header().Add("content-type","application/json")
	rw.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(rw, string(data))
}

func deleteUserHandler (rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, err)
	}

	_, ok := userMap[id]
	if !ok {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprint(rw,"No User ID:", id)
		return
	}

	delete(userMap, id)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "Deleted User ID:", id)

}

func updateUserHandler(rw http.ResponseWriter, r *http.Request) {

	updateUser := new(User)
	err := json.NewDecoder(r.Body).Decode(updateUser)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "Bad Request: ", err)
		return
	}
	
	user, ok := userMap[updateUser.ID]

	if !ok {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprint(rw, "No User ID:", updateUser.ID)
		return
	}

	if updateUser.FirstName != "" {
		user.FirstName = updateUser.FirstName
	}

	if updateUser.LastName != "" {
		user.LastName = updateUser.LastName
	}

	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	rw.Header().Add("content-type","application/json")	
	rw.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(rw, string(data))

	//put(userMap, id)


}

func NewHandler() http.Handler {
	userMap = make(map[int]*User)
	lastID = 0
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler) 
	mux.HandleFunc("/users", usersHandler).Methods("GET")
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users", updateUserHandler).Methods("PUT")
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler)
	
	return mux
}
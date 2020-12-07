package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User struct
type User struct {
	Name string
	Age  int
}

// Users - global in-memory user database
var Users []User

func create(w http.ResponseWriter, r *http.Request) {
	log.Println("Create route hit:", r.UserAgent())

	var u User

	json.NewDecoder(r.Body).Decode(&u)

	Users = append(Users, u)
}

func get(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(Users)

	w.Write([]byte(data))
}

func update(w http.ResponseWriter, r *http.Request) {
}

func delete(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/user/create", create)
	http.HandleFunc("/user/get", get)
	http.HandleFunc("/user/update", update)
	http.HandleFunc("/user/delete", delete)

	http.ListenAndServe(":8080", nil)
}

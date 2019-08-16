package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User represents a user instance
type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	// Background *Background `json:"background,omitempty"`
}

// type Background struct {
// 	DateOfBirth          string `json:"dateofbirth,omitempty"`
// 	CityOfBirth          string `json:"cityofbirth,omitempty"`
// 	CountryOfCitizenship string `json:"countryofcitizenship,omitempty"`
// 	CertificationBody    string `json:"certificationbody,omitempty"`
// 	CertificationLevel   string `json:"certificationlevel,omitempty"`
// 	CertificationYear    string `json:"certificationyear,omitempty"`
// }

var users []User

// GetUsers returns a list of all User Instances
func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

// GetUser returns a specific instance of User by its ID number
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

// CreateUser receives data for persistance in the database of a User instance
func CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = params["id"]
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

// // UpdateUser receives data for updating a instance of a User
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode("User Updated")
// }

// DeleteUser deletes an instance of User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	router := mux.NewRouter()

	users = append(users, User{
		ID:        "1",
		FirstName: "hardcoded1",
		LastName:  "hardcoded1",
		Email:     "hardcoded@hardcoded1.hard",
	})

	users = append(users, User{
		ID:        "2",
		FirstName: "hardcoded2",
		LastName:  "hardcoded2",
		Email:     "hardcoded@hardcoded2.hard",
	})

	// endpoints
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
	// router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}

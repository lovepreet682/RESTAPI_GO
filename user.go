package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:rehan200@tcp(localhost:3306)/golang"

type GoData struct {
	//gorm.Model
	Id    int    `json:"Id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
	City  string `json:"City"`
}

func initialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Connot Connect to DB")
	}
	DB.AutoMigrate(&GoData{})
}

// Fetch the records
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var golang []GoData
	DB.Find(&golang)
	json.NewEncoder(w).Encode(golang)
}

// get user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var golang []GoData
	DB.First(&golang, params["id"])
	json.NewEncoder(w).Encode(golang)
}

// Creating the User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON data from the request body into a User struct
	var goData GoData
	err := json.NewDecoder(r.Body).Decode(&goData)
	if err != nil {
		// Handle parsing error
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Validate the user data
	if goData.Name == "" || goData.Email == "" || goData.City == "" {
		// Handle validation error
		http.Error(w, "Name and Email and City are required", http.StatusBadRequest)
		return
	}
	// Return a success response
	w.WriteHeader(http.StatusCreated)
	DB.Create(&goData)
	json.NewEncoder(w).Encode(goData)
	fmt.Println("User created successfully")
}

// delete the data by using ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var golang []GoData
	DB.Delete(&golang, params["id"])
	json.NewEncoder(w).Encode("The Data is delete successfully")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var golang GoData
	DB.First(&golang, params["id"])
	json.NewDecoder(r.Body).Decode(&golang)
	DB.Save(&golang)
	json.NewEncoder(w).Encode(golang)
}

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	id    int
	Make  string
	Model string
	Price int
}

var vehicles = []Vehicle{
	{1, "Toyota", "Corolla", 10000},
	{2, "Toyota", "Camry", 20000},
	{1, "HONDA", "Civic", 10000},
}

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}
func returnCarsByBrand(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	carM := vars["make"]
	cars := []Vehicle{}
	for _, car := range vehicles {
		if car.Make == carM {
			cars = append(cars, car)
		}
	}
	json.NewEncoder(w).Encode(cars)
}
func returnCarsbyID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func updateCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func createCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func removeCarByIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/cars", returnAllCars).Methods("GET")
	router.HandleFunc("/cars/make/{make}", returnCarsByBrand).Methods("GET")
	router.HandleFunc("/cars/{id}", returnCarsbyID).Methods("GET")
	router.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	router.HandleFunc("/cars", createCar).Methods("POST")
	router.HandleFunc("/cars/{id}", removeCarByIndex).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}

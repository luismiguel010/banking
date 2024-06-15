package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {

	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomer)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World!")
}

func getAllCustomer(writer http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{Name: "Luis", City: "Medellín", Zipcode: "0000"},
		{Name: "Daniel", City: "Medellín", Zipcode: "0000"},
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(customers)
}

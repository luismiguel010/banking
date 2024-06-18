package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city"      xml:"city"`
	Zipcode string `json:"zip_code"  xml:"zipcode"`
}

func greet(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World!")
}

func getAllCustomer(writer http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{Name: "Luis", City: "Medellín", Zipcode: "0000"},
		{Name: "Daniel", City: "Medellín", Zipcode: "0000"},
	}

	if request.Header.Get("Content-Type") == "application/xml" {
		writer.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(writer).Encode(customers)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(customers)
	}
}

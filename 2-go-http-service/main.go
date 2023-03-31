package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	http.HandleFunc("/person", personHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	// Örnek bir kişi verisi
	person := Person{Name: "Ahmet", Age: 27}

	// Veriyi JSON formatına dönüştürme
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// HTTP yanıtı olarak JSON verisini yazdırma
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

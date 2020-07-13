package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User : people info
type User struct {
	Name string
	Habits []string
}

func write(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-custom-Header", "custom")
	w.WriteHeader(201)

	user := &User{
		Name: "yeshan",
		Habits: []string{"breaking", "coding"},
	}

	json, _ := json.Marshal(user)

	w.Write(json)
}

func main() {
	http.HandleFunc("/write", write)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
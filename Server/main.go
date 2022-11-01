package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var user Person

func main() {
	r := mux.NewRouter()
	r.Handlefunc("/person/{id}", GetUser)
	log.Fatalf(http.ListenAndServe("localhost:8080", r))

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatalf("Could not process your request")
		return
	}
	var user []Person
	err = json.Unmarshal(data, &user)

	for _, person := range user {
		if person.ID == id {
			response, err := json.Marshal(user)
			if err != nil {
				log.Fatalf("Could not process your request")
				return
			}
		}
	}
}

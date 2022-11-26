package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	filePtr := flag.String("file", "items.json", "location of product JSON file")
	flag.Parse()
	initProducts(*filePtr)

	r := mux.NewRouter()

	r.HandleFunc("/items", getItemsHandler)
	r.HandleFunc("/items", createItemHandler).Methods(http.MethodPost)
	r.HandleFunc("/items/{id}", updateItemHandler).Methods(http.MethodPut)
	r.HandleFunc("/items/{id}", deleteItemHandler).Methods(http.MethodDelete)
	r.HandleFunc("/items/{id}", getItemHandler)

}

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func (m *Item) validate() error {
	if m.Name == "" {
		return fmt.Errorf("missing item name")
	}
	if m.Price == 0 {
		return fmt.Errorf("missing item price")
	}
	return nil
}

var items []Item

func getItemsHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(items); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Item
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Printf("problem decoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := m.validate(); err != nil {
		log.Printf("product validation error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m.ID = len(items) + 1

	if err := m.validate(); err != nil {
		log.Printf("product validation error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	items = append(items, m)
	w.WriteHeader(http.StatusCreated)
}

func getItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("error converting id to int: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := lookupItem(id)
	if m == nil {
		log.Printf("product with id %d not found", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("error encoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func updateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("error converting id to int: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := lookupItem(id)
	if m == nil {
		log.Printf("product with id %d not found", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Printf("error decoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	m.ID = id

	if err := m.validate(); err != nil {
		log.Printf("cannot validate item: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for p := range items {
		if id == items[p].ID {
			items[p] = *m
			log.Printf("item updated with id %d", id)
			break
		}
	}
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("failed to convert id to int: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if i := lookupItem(id); i == nil {
		log.Printf("product with id %d not found", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for p, prod := range items {
		if prod.ID == items[p].ID {
			items = append(items[:p], items[p+1:]...)
			log.Printf("item with id %d deleted", id)
			break
		}
	}
}

func lookupItem(id int) *Item {
	for _, m := range items {
		if m.ID == id {
			return &m
		}
	}
	return nil
}

func initProducts(filepath string) {
	bs, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bs, &items); err != nil {
		log.Fatal(err)
	}
}

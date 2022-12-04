package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "items.db")
	if err != nil {
		fmt.Println("could not reach database")
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("no response from database")
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/items", GetItems).Methods("GET")
	r.HandleFunc("/items/{id}", GetItem).Methods("GET")
	r.HandleFunc("/items", AddItem).Methods("POST")
	r.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	var items []Item
	limit := r.FormValue("limit")
	stmt, err := db.Prepare("SELECT * FROM items LIMIT ?;")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := stmt.Query(limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = stmt.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for result.Next() {
		var item Item
		err := result.Scan(&item.ID, &item.Name, &item.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringid := vars["id"]
	intid, err := strconv.Atoi(stringid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if intid > 100 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	stmt, err := db.Prepare("SELECT * FROM item WHERE ID =?;")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := stmt.QueryRow(intid)
	err = stmt.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var item Item
	err = result.Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func AddItem(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newitem := Item{
		ID:    0,
		Name:  "",
		Price: 0,
	}

	err = json.Unmarshal(request, &newitem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	stmt, err := db.Prepare("INSERT INTO item (ID, Name, Price) VALUES (?,?,?);")
	_, err = stmt.Exec(newitem.ID, newitem.Name, newitem.Price)
	w.WriteHeader(http.StatusCreated)

}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringid := vars["id"]
	id, err := strconv.Atoi(stringid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if id > 100 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	stmt, err := db.Prepare("DELETE FROM items WHERE ID =?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringid := vars["id"]
	id, err := strconv.Atoi(stringid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if id > 100 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newitem := Item{
		ID:    0,
		Name:  "",
		Price: 0,
	}

	err = json.Unmarshal(request, &newitem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("UPDATE items SET Name =?, Price =?, WHERE ID =?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(newitem.Name, newitem.Price, newitem.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}

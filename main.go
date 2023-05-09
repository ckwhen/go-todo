package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "wenchihkai"
	DB_PASSWORD = "password"
	DB_NAME     = "go-todo"
)

type Todo struct {
	ID          string `json:id`
	Task        string `json:task`
	IsCompleted bool   `json:isCompleted`
}
type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Todo `json:"data"`
	Message string `json:"message"`
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	rows, err := db.Query("SELECT * FROM todos")
	checkErr(err)

	todos := make([]Todo, 0)

	for rows.Next() {
		todo := new(Todo)

		err = rows.Scan(&todo.ID, &todo.Task, &todo.IsCompleted)

		checkErr(err)

		todos = append(todos, *todo)
	}

	var response = JsonResponse{Type: "success", Data: todos}

	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	api := r.PathPrefix("/v1").Subrouter()

	todos := api.PathPrefix("/todos").Subrouter()
	todos.HandleFunc("/", getTodos).Methods("GET")

	http.Handle("/", r)

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

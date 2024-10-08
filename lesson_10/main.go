package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ToDo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type ToDoStore struct {
	tasks []ToDo
}

func (store *ToDoStore) getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.tasks)
}

func (store *ToDoStore) getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range store.tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&ToDo{})
}

func (store *ToDoStore) createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task ToDo
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(100000000))
	store.tasks = append(store.tasks, task)
	json.NewEncoder(w).Encode(task)
}

func (store *ToDoStore) updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range store.tasks {
		if item.ID == params["id"] {
			store.tasks = append(store.tasks[:index], store.tasks[index+1:]...)
			var task ToDo
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			store.tasks = append(store.tasks, task)
			json.NewEncoder(w).Encode(task)
			return
		}
	}
}

func (store *ToDoStore) deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range store.tasks {
		if item.ID == params["id"] {
			store.tasks = append(store.tasks[:index], store.tasks[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(store.tasks)
}

func main() {
	store := &ToDoStore{
		tasks: []ToDo{
			{ID: "1", Title: "Task One"},
			{ID: "2", Title: "Task Two"},
		},
	}

	r := mux.NewRouter()
	r.HandleFunc("/tasks", store.getTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", store.getTask).Methods("GET")
	r.HandleFunc("/tasks", store.createTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", store.updateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", store.deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	ID     int                `json:"id"`
	Name   string             `json:"name"`
	Grades map[string]float64 `json:"grades"`
}

type Class struct {
	ClassName string          `json:"class_name"`
	Students  map[int]Student `json:"students"`
}

var classData = Class{
	ClassName: "10-A",
	Students: map[int]Student{
		1: {ID: 1, Name: "John Doe", Grades: map[string]float64{"math": 90, "science": 85}},
		2: {ID: 2, Name: "Jane Smith", Grades: map[string]float64{"math": 95, "science": 80}},
	},
}

func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !validateUser(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func validateUser(username, password string) bool {
	return username == "teacher" && password == "password123"
}

func getClassInfo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(classData)
}

func getStudentInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	student, exists := classData.Students[studentID]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/class", basicAuth(getClassInfo)).Methods("GET")
	r.HandleFunc("/student/{id}", basicAuth(getStudentInfo)).Methods("GET")

	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

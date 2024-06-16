package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Student struct {
	ID     int
	Name   string
	Grades map[string]float64
}

type Class struct {
	Name     string
	Teacher  string
	Students []Student
}

var classData Class

func init() {
	classData = Class{
		Name:    "5-B",
		Teacher: "Mr. Smith",
		Students: []Student{
			{ID: 1, Name: "John Doe", Grades: map[string]float64{"Math": 85, "English": 90}},
			{ID: 2, Name: "Jane Doe", Grades: map[string]float64{"Math": 92, "English": 88}},
		},
	}
}

func main() {
	http.HandleFunc("/class", classHandler)
	http.HandleFunc("/student/", studentHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func classHandler(w http.ResponseWriter, r *http.Request) {
	if !isTeacher(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(classData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	if !isTeacher(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/student/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	student, found := getStudentByID(id)
	if !found {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func isTeacher(r *http.Request) bool {
	return r.Header.Get("X-Teacher") == classData.Teacher
}

func getStudentByID(id int) (Student, bool) {
	for _, student := range classData.Students {
		if student.ID == id {
			return student, true
		}
	}
	return Student{}, false
}

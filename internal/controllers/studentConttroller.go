package controllers

import (
	"net/http"
	"rest-project/internal/services"
)

func New() {
	http.HandleFunc("/students", services.GetAllStudents)
	http.HandleFunc("/students/one", services.GetStudentById)
	http.HandleFunc("/students/add", services.CreateStudent)
	http.HandleFunc("/students/update", services.UpdateStudent)
	http.HandleFunc("/students/delete", services.DeleteStudent)
}

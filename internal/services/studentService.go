package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"rest-project/internal/models"
	"strconv"
)

var (
	id       = 4
	students = []models.Student{
		{Id: 1, FullName: "Муратов Алихан Сейдахметович", Birthdate: "2005-05-13", Age: 18},
		{Id: 2, FullName: "Болатов Акбар Нуркенович", Birthdate: "2000-07-22", Age: 23},
		{Id: 3, FullName: "Даулетханова Айнур Муратовна", Birthdate: "2009-02-18", Age: 14},
	}
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(students)
	if err != nil {
		log.Fatal(err)
	}
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	for _, student := range students {
		if student.Id == id {
			err := json.NewEncoder(w).Encode(student)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var studentCreate models.StudentEdit

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &studentCreate)
	if err != nil {
		log.Fatal(err)
	}

	newStudent := models.Student{
		Id:        id,
		FullName:  studentCreate.FullName,
		Birthdate: studentCreate.Birthdate,
		Age:       studentCreate.Age,
	}

	id += 1

	students = append(students, newStudent)

	err = json.NewEncoder(w).Encode(newStudent)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var studentEdit models.StudentEdit
	studentId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &studentEdit)
	if err != nil {
		log.Fatal(err)
	}

	updatedStudent := models.Student{
		FullName:  studentEdit.FullName,
		Birthdate: studentEdit.Birthdate,
		Age:       studentEdit.Age,
	}

	for i, student := range students {
		if student.Id == studentId {

			updatedStudent.Id = student.Id
			students = append(students[:i], students[i+1:]...)
			students = append(students, updatedStudent)
			break
		}
	}

	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteStudent(writer http.ResponseWriter, request *http.Request) {
	studentId, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	for i, student := range students {
		if student.Id == studentId {
			students = append(students[:i], students[i+1:]...)
			break
		}
	}
	
	err = json.NewEncoder(writer).Encode(students)
	if err != nil {
		log.Fatal(err)
	}
}

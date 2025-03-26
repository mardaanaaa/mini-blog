package services

import (
	"errors"
	"rest-project/internal/models"
)

type StudentService struct{}

var (
	id       = 4
	students = []models.Student{
		{Id: 1, FullName: "Муратов Алихан Сейдахметович", Birthdate: "2005-05-13", Age: 18},
		{Id: 2, FullName: "Болатов Акбар Нуркенович", Birthdate: "2000-07-22", Age: 23},
		{Id: 3, FullName: "Даулетханова Айнур Муратовна", Birthdate: "2009-02-18", Age: 14},
	}
)

// 📌 Получить список всех студентов
func (s *StudentService) GetAllStudents() []models.Student {
	return students
}

// 📌 Получить студента по ID
func (s *StudentService) GetStudentById(id int) (models.Student, error) {
	for _, student := range students {
		if student.Id == id {
			return student, nil
		}
	}
	return models.Student{}, errors.New("student not found")
}

// 📌 Создать нового студента
func (s *StudentService) CreateStudent(studentEdit models.StudentEdit) models.Student {
	newStudent := models.Student{
		Id:        id,
		FullName:  studentEdit.FullName,
		Birthdate: studentEdit.Birthdate,
		Age:       studentEdit.Age,
	}

	id++
	students = append(students, newStudent)

	return newStudent
}

// 📌 Обновить данные студента
func (s *StudentService) UpdateStudent(studentId int, studentEdit models.StudentEdit) (models.Student, error) {
	for i, student := range students {
		if student.Id == studentId {
			updatedStudent := models.Student{
				Id:        student.Id,
				FullName:  studentEdit.FullName,
				Birthdate: studentEdit.Birthdate,
				Age:       studentEdit.Age,
			}

			students[i] = updatedStudent
			return updatedStudent, nil
		}
	}
	return models.Student{}, errors.New("student not found")
}

// 📌 Удалить студента
func (s *StudentService) DeleteStudent(studentId int) error {
	for i, student := range students {
		if student.Id == studentId {
			students = append(students[:i], students[i+1:]...)
			return nil
		}
	}
	return errors.New("student not found")
}

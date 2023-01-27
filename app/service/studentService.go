package service

import (
	"api-sekolah/app/repository"
	"api-sekolah/app/request"
	"api-sekolah/models"
	"errors"
	"reflect"
)

type StudentService interface {
	List() ([]models.Student, error)
	FindByID(ID int) (models.Student, error)
	FindByNIM(NIM string) (models.Student, error)
	Store(studentRequest request.StudentRequest) (models.Student, error)
	Update(ID int, studentRequest request.StudentUpdateRequest) (models.Student, error)
	Destroy(ID int) (models.Student, error)
}

type studentService struct {
	studentRepository repository.StudentRepository
}

func NewStudentService(studentRepository repository.StudentRepository) *studentService {
	return &studentService{studentRepository}
}

func (s *studentService) List() ([]models.Student, error) {
	students, err := s.studentRepository.List()

	return students, err
}

func (s *studentService) FindByID(ID int) (models.Student, error) {
	student, err := s.studentRepository.FindByID(ID)

	return student, err
}

func (s *studentService) FindByNIM(NIM string) (models.Student, error) {
	student, err := s.studentRepository.FindByNIM(NIM)

	return student, err
}

func (s *studentService) Store(studentRequest request.StudentRequest) (models.Student, error) {
	ageReq, _ := studentRequest.Age.Int64()
	age := int(ageReq)

	student := models.Student{
		Name:  studentRequest.Name,
		NIM:   studentRequest.NIM,
		Age:   age,
		Grade: studentRequest.Grade,
	}

	student, err := s.studentRepository.Store(student)

	return student, err
}

func (s *studentService) Update(ID int, studentRequest request.StudentUpdateRequest) (models.Student, error) {
	student, err := s.studentRepository.FindByID(ID)

	if err != nil {
		return student, errors.New("ada error nih")
	}

	ageReq, _ := studentRequest.Age.Int64()
	age := int(ageReq)

	student.Name = studentRequest.Name
	student.Age = age
	student.Grade = studentRequest.Grade

	updateStudent, err := s.studentRepository.Update(student)

	return updateStudent, err
}

func (s *studentService) Destroy(ID int) (models.Student, error) {
	student, err := s.studentRepository.FindByID(ID)

	if err != nil {
		return student, errors.New("ada error nih")
	}

	if !reflect.ValueOf(student).IsZero() {
		return student, err
	}

	deleteStudent, err := s.studentRepository.Destroy(student)

	return deleteStudent, err
}

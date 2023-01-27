package repository

import (
	"api-sekolah/models"

	"gorm.io/gorm"
)

type StudentRepository interface {
	List() ([]models.Student, error)
	FindByID(ID int) (models.Student, error)
	FindByNIM(NIM string) (models.Student, error)
	Store(student models.Student) (models.Student, error)
	Update(student models.Student) (models.Student, error)
	Destroy(student models.Student) (models.Student, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *studentRepository {
	return &studentRepository{db}
}

func (r *studentRepository) List() ([]models.Student, error) {
	var students []models.Student
	err := r.db.Find(&students).Error

	return students, err
}

func (r *studentRepository) FindByID(ID int) (models.Student, error) {
	var student models.Student
	err := r.db.Find(&student, ID).Error

	return student, err
}

func (r *studentRepository) FindByNIM(NIM string) (models.Student, error) {
	var student models.Student
	err := r.db.Where("nim = ?", NIM).Find(&student).Error

	return student, err
}

func (r *studentRepository) Store(student models.Student) (models.Student, error) {
	err := r.db.Create(&student).Error

	return student, err
}

func (r *studentRepository) Update(student models.Student) (models.Student, error) {
	err := r.db.Save(&student).Error

	return student, err
}

func (r *studentRepository) Destroy(student models.Student) (models.Student, error) {
	err := r.db.Delete(&student).Error

	return student, err
}

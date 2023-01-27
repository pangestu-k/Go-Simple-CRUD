package controller

import (
	"api-sekolah/app/request"
	"api-sekolah/app/service"
	"api-sekolah/helper"
	"api-sekolah/initialize"
	"api-sekolah/models"
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type studentController struct {
	studentService service.StudentService
}

func init() {
	initialize.LoadDatabase()
}

func NewStudentController(studentService service.StudentService) *studentController {
	return &studentController{studentService}
}

func (ctr *studentController) GetList(c *gin.Context) {
	students, err := ctr.studentService.List()

	if err != nil {
		c.JSON(400, gin.H{
			"message": "ups Something Wrong",
		})
	}

	var studentsResponse []helper.StudentResponse
	for _, s := range students {
		studentResponse := convertToResponse(s)

		studentsResponse = append(studentsResponse, studentResponse)
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   studentsResponse,
	})
}

func (ctr *studentController) GetByID(c *gin.Context) {
	idParams := c.Param("id")
	id, _ := strconv.Atoi(idParams)

	student, err := ctr.studentService.FindByID(id)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "ups Something Wrong",
		})
	}

	studentResonse := convertToResponse(student)

	if reflect.ValueOf(studentResonse).IsZero() {
		c.JSON(200, gin.H{
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": studentResonse,
	})
}

func (ctr *studentController) StoreData(c *gin.Context) {
	var studentRequest request.StudentRequest

	c.ShouldBind(&studentRequest)

	v := validator.New()
	err := v.Struct(studentRequest)

	if err != nil {
		var requestValidate []string

		for _, e := range err.(validator.ValidationErrors) {
			requestValidate = append(requestValidate, e.Error())
			fmt.Println(e)
		}

		c.JSON(422, gin.H{
			"validation": requestValidate,
		})
		return
	}

	student, _ := ctr.studentService.FindByNIM(studentRequest.NIM)

	if !reflect.ValueOf(student).IsZero() {
		c.JSON(422, gin.H{
			"validation": "NIM already been taken.",
		})
		return
	}

	post, err := ctr.studentService.Store(studentRequest)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Ups SOmething Wrong with Your Query",
		})
		return
	}

	studentResponse := convertToResponse(post)

	c.JSON(200, gin.H{
		"data": studentResponse,
	})
}

func (ctr *studentController) UpdateData(c *gin.Context) {
	idParams := c.Param("id")
	id, _ := strconv.Atoi(idParams)

	var studentRequest request.StudentUpdateRequest

	c.ShouldBind(&studentRequest)

	v := validator.New()
	err := v.Struct(studentRequest)

	if err != nil {
		var requestValidate []string

		for _, e := range err.(validator.ValidationErrors) {
			requestValidate = append(requestValidate, e.Error())
			fmt.Println(e)
		}

		c.JSON(422, gin.H{
			"validation": requestValidate,
		})
		return
	}

	updateStudent, err := ctr.studentService.Update(id, studentRequest)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Ups SOmething Wrong with Your Query",
		})
		return
	}

	updateResponse := convertToResponse(updateStudent)

	c.JSON(200, gin.H{
		"data": updateResponse,
	})
}

func (ctr *studentController) DeleteData(c *gin.Context) {
	idParams := c.Param("id")
	id, _ := strconv.Atoi(idParams)

	deleteStudent, err := ctr.studentService.Destroy(id)

	if !reflect.ValueOf(deleteStudent).IsZero() {
		println(err)
		c.JSON(404, gin.H{
			"message": "Data not found",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Something wrong with it",
		})
		return
	}

	deleteResponse := convertToResponse(deleteStudent)

	c.JSON(200, gin.H{
		"data":    deleteResponse,
		"message": "Data berhasi di Hapus",
	})
}

func convertToResponse(s models.Student) helper.StudentResponse {
	studentResponse := helper.StudentResponse{
		ID:    s.ID,
		Name:  s.Name,
		NIM:   s.NIM,
		Age:   s.Age,
		Grade: s.Grade,
	}

	return studentResponse
}

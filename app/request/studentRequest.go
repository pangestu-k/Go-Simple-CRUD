package request

import "encoding/json"

type StudentRequest struct {
	Name  string      `form:"name" json:"name" bson:"name" validate:"required"`
	NIM   string      `form:"nim" json:"nim" bson:"nim" validate:"required"`
	Age   json.Number `form:"age" json:"age" bson:"age" validate:"required,number"`
	Grade string      `form:"grade" json:"grade" bson:"grade" validate:"required"`
}

type StudentUpdateRequest struct {
	Name  string      `form:"name" json:"name" bson:"name" validate:"required"`
	Age   json.Number `form:"age" json:"age" bson:"age" validate:"required,number"`
	Grade string      `form:"grade" json:"grade" bson:"grade" validate:"required"`
}

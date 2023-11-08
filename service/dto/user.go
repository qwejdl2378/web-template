package dto

import "github.com/qwejdl2378/web-template/model"

type UserDTO struct {
	UserType int8     `json:"UserType"`
	UserName string   `json:"UserName"`
	Age      int16    `json:"Age"`
	Habits   []string `json:"Habits"`
}

func (u UserDTO) ToUserModel() *model.User {
	return &model.User{
		Name: u.UserName,
		Age:  u.Age,
	}
}

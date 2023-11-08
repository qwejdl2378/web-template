package service

import (
	"github.com/qwejdl2378/web-template/model"
	"github.com/qwejdl2378/web-template/service/dto"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) CreateUser(dto *dto.UserDTO) error {
	newUser := dto.ToUserModel()
	_, err := model.CreateUser(u.db, newUser)
	return err
}

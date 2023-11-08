package model

import "gorm.io/gorm"

type User struct {
	ID   int32
	Name string
	Age  int16
}

func CreateUser(db *gorm.DB, user *User) (*User, error) {
	err := db.Create(user).Error
	return user, err
}

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string
	Email      string
	ProfileUrl string
	Password   string
}

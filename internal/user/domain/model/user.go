package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Name     string
	Password string `json:"-"`
}

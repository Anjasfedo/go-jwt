package models

import ()

type User struct {
	Id       uint
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

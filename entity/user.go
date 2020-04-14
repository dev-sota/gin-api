package entity

import "github.com/jinzhu/gorm"

// User is users tabel property
type User struct {
	gorm.Model
	Name  string
	Email string
	Role  string
}

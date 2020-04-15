package users

import (
	"github.com/dev-sota/gin-api/entities"
	"github.com/dev-sota/gin-api/infrastructure/database"
	"github.com/gin-gonic/gin"
)

// Repository procides user'r behavior
type Repository struct{}

// User is alias of entities.User struct
type User entities.User

// GetAll is get all User
func (r Repository) GetAll() ([]User, error) {
	db := database.Get()
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// CreateModel is create User model
func (r Repository) CreateModel(c *gin.Context) (User, error) {
	db := database.Get()
	var u User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// GetByID is get a User
func (r Repository) GetByID(id string) (User, error) {
	db := database.Get()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// UpdateByID is update a User
func (r Repository) UpdateByID(id string, c *gin.Context) (User, error) {
	db := database.Get()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

// DeleteByID is delete a User
func (r Repository) DeleteByID(id string) error {
	db := database.Get()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}

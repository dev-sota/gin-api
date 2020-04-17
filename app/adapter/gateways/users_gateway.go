package users

import (
	"github.com/dev-sota/gin-api/app/domain"
	"github.com/dev-sota/gin-api/app/infrastructure/database"
	"github.com/gin-gonic/gin"
)

// Gateway procides user's behavior
type Gateway struct{}

// User is alias of domain.User struct
type User domain.User

// GetAll is get all User
func (g Gateway) GetAll() ([]User, error) {
	db := database.Get()
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// CreateModel is create User model
func (g Gateway) CreateModel(c *gin.Context) (User, error) {
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
func (g Gateway) GetByID(id string) (User, error) {
	db := database.Get()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// UpdateByID is update a User
func (g Gateway) UpdateByID(id string, c *gin.Context) (User, error) {
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
func (g Gateway) DeleteByID(id string) error {
	db := database.Get()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}

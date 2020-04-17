package users

import "github.com/dev-sota/gin-api/app/domain"

// Repository interface
type Repository interface {
	Store(domain.User) (int, error)
	FindByID(id int) ([]domain.User, error)
	FindAll() ([]domain.User, error)
}

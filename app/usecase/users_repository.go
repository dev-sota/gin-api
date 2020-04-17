package users

import "github.com/dev-sota/gin-api/app/domain"

// UserRepository interface
type UserRepository interface {
	Store(domain.User) (int, error)
	FindByID(id int) ([]domain.User, error)
	FindAll() ([]domain.User, error)
}

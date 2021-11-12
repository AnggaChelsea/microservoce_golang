package repository

import (
	"simplelogic/domain/entity"
)

type UserRepository interface {
	GetUser(id int) (entity.User, error)
	GetUsers() ([]entity.User, error)
	CreateUser(user entity.User) error
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
}

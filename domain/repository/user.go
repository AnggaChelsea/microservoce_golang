package repository

import (
	"context"
	"simplelogic/domain/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(ctx context.Context, ID int  ) (entity.User, error)
	GetUsers() ([]entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) GetUsers(ctx context.Context) ([]entity.User, error) {
// 	var user []entity.User
// 	err := r.db.Find(ctx, &user)
// 	return user
// }
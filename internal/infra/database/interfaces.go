package database

import "github.com/jadson-medeiros/goapi/internal/entity"

type Userinterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
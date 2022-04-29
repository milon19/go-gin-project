package repository

import "healthousedemo/internal/app/domain"

type IUser interface {
	GetUserByEmail(email string) (*domain.User, error)
	CreateUser(user domain.User) error
}

package factory

import "healthousedemo/internal/app/domain"

type User struct{}

func (u User) Generate(id string, name string, email string, userType string) *domain.User {
	return &domain.User{ID: id, Name: name, Email: email, Type: userType}
}

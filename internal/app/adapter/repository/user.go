package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	db "healthousedemo/internal/app/adapter/db/connections"
	"healthousedemo/internal/app/adapter/db/models"
	"healthousedemo/internal/app/domain"
	"healthousedemo/internal/app/domain/factory"
)

const (
	CollectionUser = "user"
)

type User struct{}

func (u User) GetUserByEmail(email string) (*domain.User, error) {
	_db := db.GetDB(CollectionUser)
	ctx, cancel := db.GetCtx()
	defer cancel()
	filter := bson.M{"email": bson.M{"$eq": email}}
	var user models.User
	if err := _db.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	userFactory := factory.User{}
	return userFactory.Generate(user.UserId, user.Name, user.Email, user.Type), nil
}

func (u User) CreateUser(user domain.User) error {
	_db := db.GetDB(CollectionUser)
	ctx, cancel := db.GetCtx()
	defer cancel()
	_, err := _db.InsertOne(ctx, models.User{UserId: user.ID, Email: user.Email, Name: user.Name, Type: "supplier"})
	if err != nil {
		return err
	}
	return nil
}

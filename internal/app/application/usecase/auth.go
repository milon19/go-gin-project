package usecase

import (
	"encoding/json"
	"fmt"
	"healthousedemo/internal/app/adapter/utils"
	domain "healthousedemo/internal/app/domain"
	"healthousedemo/internal/app/domain/repository"
)

type Auth struct{}

func (au Auth) GoogleAuthenticationCallBack(url string, userRepository repository.IUser) error {
	api := utils.Api{Url: url, Method: "GET", Body: nil}
	response, err := api.MakeRequest()
	if err != nil {
		fmt.Println(err)
		return err
	}
	var authUser domain.OAuthResponse
	if err := json.Unmarshal(response, &authUser); err != nil {
		fmt.Println(err)
		return err
	}
	_, err = userRepository.GetUserByEmail(authUser.Email)
	if err != nil {
		err := userRepository.CreateUser(domain.User{ID: authUser.ID, Email: authUser.Email, Name: authUser.Name})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

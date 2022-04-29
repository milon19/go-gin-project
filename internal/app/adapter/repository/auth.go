package repository

import "healthousedemo/internal/app/adapter/config"

type Auth struct{}

func (a Auth) GoogleLogin() string {
	u := config.AppConfig.GoogleLoginConfig
	return u.AuthCodeURL("auth0-demo")
}

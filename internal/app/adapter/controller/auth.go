package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"healthousedemo/internal/app/adapter/config"
	"healthousedemo/internal/app/adapter/repository"
	"healthousedemo/internal/app/application/usecase"
	"net/http"
)

type AuthController struct{}

type authUseCase = usecase.Auth

func AuthRoutes(r *gin.Engine) *gin.Engine {
	authController := AuthController{}

	r.GET("/login", func(c *gin.Context) {
		authController.googleLogin(c)
	})
	r.GET("/callback", func(c *gin.Context) {
		authController.callBack(c)
	})
	return r
}

func (au AuthController) googleLogin(c *gin.Context) {
	url := repository.Auth{}.GoogleLogin()
	c.Redirect(302, url)
}

func (au AuthController) callBack(c *gin.Context) {
	code := c.Request.URL.Query().Get("code")
	c.Request.Header.Add("content-type", "application/json")
	token, err := config.AppConfig.GoogleLoginConfig.Exchange(context.Background(), code)
	var userRepository repository.User
	err = authUseCase{}.GoogleAuthenticationCallBack(config.OauthGoogleUrlAPI+token.AccessToken, userRepository)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadGateway, gin.H{"err": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": token.AccessToken, "err": nil})
}

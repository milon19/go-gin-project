package repository

type IAuthentication interface {
	GoogleLogin() string
	GoogleCallBack() string
}

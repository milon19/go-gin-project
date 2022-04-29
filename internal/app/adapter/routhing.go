package adapter

import (
	"github.com/gin-gonic/gin"
	"healthousedemo/internal/app/adapter/controller"
)

func Routes(r *gin.Engine) *gin.Engine {
	r = controller.AuthRoutes(r)
	return r
}

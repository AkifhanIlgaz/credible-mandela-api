package routers

import (
	"github.com/AkifhanIlgaz/credible-mandela-api/controllers"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/constants"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	authController controllers.AuthController
}

func NewAuthRouter(authController controllers.AuthController) AuthRouter {
	return AuthRouter{
		authController: authController,
	}
}

func (r AuthRouter) Setup(rg *gin.RouterGroup) {
	router := rg.Group(constants.AuthPath)

	router.POST("/login", r.authController.Login)
	router.POST("/register", r.authController.Register)
	router.POST("/logout", r.authController.Logout)
}

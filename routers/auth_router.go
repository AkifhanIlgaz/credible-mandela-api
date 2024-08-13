package routers

import (
	"github.com/AkifhanIlgaz/credible-mandela-api/controllers"
	"github.com/gin-gonic/gin"
)

const (
	AuthPath     string = "/auth"
	LoginPath    string = "/login"
	RegisterPath string = "/register"
	LogoutPath   string = "/logout"
	RefreshPath  string = "/refresh"
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
	router := rg.Group(AuthPath)

	router.POST(LoginPath, r.authController.Login)
	router.POST(RegisterPath, r.authController.Register)
	router.POST(LogoutPath, r.authController.Logout)
	router.POST(RefreshPath, r.authController.Refresh)
}

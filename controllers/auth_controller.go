package controllers

import (
	"github.com/AkifhanIlgaz/credible-mandela-api/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

func (controller AuthController) Login(ctx *gin.Context) {

}

func (controller AuthController) Register(ctx *gin.Context) {

}

func (controller AuthController) Logout(ctx *gin.Context) {

}

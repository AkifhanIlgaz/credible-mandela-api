package controllers

import (
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/credible-mandela-api/models"
	"github.com/AkifhanIlgaz/credible-mandela-api/services"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/response"
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
	var form models.RegisterFormWithSignature

	if err := ctx.ShouldBindJSON(&form); err != nil {
		log.Println(err.Error())
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := form.Validate(); err != nil {
		log.Println(err.Error())
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := form.ToUser()
	if err != nil {
		log.Println(err.Error())
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: Create user on DB
}

func (controller AuthController) Logout(ctx *gin.Context) {

}

func (controller AuthController) Refresh(ctx *gin.Context) {

}

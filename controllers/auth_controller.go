package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/credible-mandela-api/models"
	"github.com/AkifhanIlgaz/credible-mandela-api/services"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/mande"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/response"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService  services.AuthService
	tokenService services.TokenService
	mandeClient  mande.Client
}

func NewAuthController(authService services.AuthService, tokenService services.TokenService, mandeClient mande.Client) AuthController {
	return AuthController{
		authService:  authService,
		mandeClient:  mandeClient,
		tokenService: tokenService,
	}
}

func (controller AuthController) Login(ctx *gin.Context) {

}

func (controller AuthController) Register(ctx *gin.Context) {
	var form models.RegisterForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := form.Validate(); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := form.ToUser()
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	cred, err := controller.mandeClient.GetCredOfUser(user.Address)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if hasEnoughCred := mande.IsEnoughCredToRegister(cred); !hasEnoughCred {
		log.Printf("%v does not have enough cred to register", user.Address)
		response.WithError(ctx, http.StatusInternalServerError, fmt.Sprintf("%v does not have enough cred to register", user.Address))
		return
	}

	userId, err := controller.authService.CreateUser(user)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusConflict, err.Error())
		return
	}

	accessToken, err := controller.tokenService.GenerateAccessToken(userId.Hex())
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusConflict, err.Error())
		return
	}

	refreshToken, err := controller.tokenService.GenerateRefreshToken(userId.Hex())
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusConflict, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, models.RegisterResponse{
		Uid:          userId.Hex(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (controller AuthController) Logout(ctx *gin.Context) {

}

func (controller AuthController) Refresh(ctx *gin.Context) {

}

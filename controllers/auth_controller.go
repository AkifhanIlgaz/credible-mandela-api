package controllers

import (
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/credible-mandela-api/models"
	"github.com/AkifhanIlgaz/credible-mandela-api/services"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/constants"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/crypto"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/mande"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/message"
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
	var form models.LoginForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := controller.authService.GetUserByUsername(form.Username)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	if err := crypto.VerifyPassword(user.PasswordHash, form.Password); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusUnauthorized, message.WrongPassword)
		return
	}

	accessToken, err := controller.tokenService.GenerateToken(constants.AccessTokenType, user.Id.Hex(), user.Address, user.Username)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := controller.tokenService.GenerateToken(constants.RefreshTokenType, user.Id.Hex(), user.Address, user.Username)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.LoginSuccess, models.LoginResponse{
		Uid:          user.Id.Hex(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
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
		response.WithError(ctx, http.StatusInternalServerError, message.InsufficientCred)
		return
	}

	userId, err := controller.authService.CreateUser(user)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusConflict, err.Error())
		return
	}

	accessToken, err := controller.tokenService.GenerateToken(constants.AccessTokenType, userId.Hex(), user.Address, user.Username)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := controller.tokenService.GenerateToken(constants.RefreshTokenType, userId.Hex(), user.Address, user.Username)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.RegisterSuccess, models.RegisterResponse{
		Uid:          userId.Hex(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (controller AuthController) Refresh(ctx *gin.Context) {
	var form models.RefreshTokenForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := controller.tokenService.ExtractUserFromRefreshToken(form.RefreshToken)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	accessToken, err := controller.tokenService.GenerateToken(constants.AccessTokenType, user.Id.Hex(), user.Address, user.Username)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := controller.tokenService.GenerateToken(constants.RefreshTokenType, user.Id.Hex(), user.Address, user.Username)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.TokensRefreshed, models.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

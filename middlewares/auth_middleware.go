package middlewares

import (
	"net/http"
	"strings"

	"github.com/AkifhanIlgaz/credible-mandela-api/services"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/constants"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/response"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authService  services.AuthService
	tokenService services.TokenService
}

func NewAuthMiddleware(authService services.AuthService, tokenService services.TokenService) AuthMiddleware {
	return AuthMiddleware{
		authService:  authService,
		tokenService: tokenService,
	}
}

func (middleware AuthMiddleware) ExtractUidFromAuthHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		t := strings.Fields(authHeader)
		if len(t) == 2 {
			accessToken := t[1]
			user, err := middleware.tokenService.ExtractUserFromAccessToken(accessToken)
			if err != nil {
				response.WithError(ctx, http.StatusUnauthorized, err.Error())
				return
			}

			ctx.Set(constants.CtxUid, user.Id)
			ctx.Set(constants.CtxUsername, user.Username)
			ctx.Set(constants.CtxAddress, user.Address)
			ctx.Next()
			return
		}
		response.WithError(ctx, http.StatusUnauthorized, "Not authorized")
	}
}

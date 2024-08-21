package routers

import (
	"github.com/AkifhanIlgaz/credible-mandela-api/controllers"
	"github.com/AkifhanIlgaz/credible-mandela-api/middlewares"
	"github.com/gin-gonic/gin"
)

const AdsPath = "/ads"

type AdRouter struct {
	adController   controllers.AdController
	authMiddleware middlewares.AuthMiddleware
}

func NewAdRouter(adController controllers.AdController, authMiddleware middlewares.AuthMiddleware) AdRouter {
	return AdRouter{
		adController:   adController,
		authMiddleware: authMiddleware,
	}
}

func (r AdRouter) Setup(rg *gin.RouterGroup) {
	router := rg.Group(AdsPath)

	router.Use(r.authMiddleware.ExtractUidFromAuthHeader())

	router.POST("/", r.adController.PublishAd)     // Publish an ad
	router.DELETE("/:id", r.adController.DeleteAd) // Delete ad by ID
	router.PUT("/:id", r.adController.UpdateAd)    // Edit ad by ID

	router.GET("/:id", r.adController.GetAdById)              // Get ad by ID
	router.GET("/", r.adController.GetAllAds)                 // Get all ads
	router.GET("/user/:address", r.adController.GetAdsOfUser) // Get all ads of user
	router.GET("/page/:page", r.adController.GetAdsByPage)    // Get ads by page
}

package routers

import (
	"github.com/AkifhanIlgaz/credible-mandela-api/controllers"
	"github.com/gin-gonic/gin"
)

const AdsPath = "/ads"

type AdRouter struct {
	adController controllers.AdController
}

func NewAdRouter(adController controllers.AdController) AdRouter {
	return AdRouter{
		adController: adController,
	}
}

func (r AdRouter) Setup(rg *gin.RouterGroup) {
	router := rg.Group(AdsPath)

	router.GET("/", r.adController.GetAllAds)                 // Get all ads
	router.POST("/", r.adController.PublishAd)                // Publish an ad
	router.DELETE("/:id", r.adController.DeleteAd)            // Delete ad by ID
	router.PUT("/:id", r.adController.EditAd)                 // Edit ad by ID
	router.GET("/user/:address", r.adController.GetAdsOfUser) // Get all ads of user
}

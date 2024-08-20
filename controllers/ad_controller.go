package controllers

import (
	"github.com/AkifhanIlgaz/credible-mandela-api/services"
	"github.com/gin-gonic/gin"
)

type AdController struct {
	adService services.AdService
}

func NewAdController(adService services.AdService) AdController {
	return AdController{
		adService: adService,
	}
}

func (controller AdController) GetAllAds(ctx *gin.Context) {

}

func (controller AdController) GetAdsByPage(ctx *gin.Context) {
	
}

func (controller AdController) GetAdById(ctx *gin.Context) {

}

func (controller AdController) GetAdsOfUser(ctx *gin.Context) {

}

func (controller AdController) PublishAd(ctx *gin.Context) {

}

func (controller AdController) DeleteAd(ctx *gin.Context) {

}

func (controller AdController) EditAd(ctx *gin.Context) {

}

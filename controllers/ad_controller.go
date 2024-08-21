package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/AkifhanIlgaz/credible-mandela-api/models"
	"github.com/AkifhanIlgaz/credible-mandela-api/services"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/constants"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/message"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/response"
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
	var form models.PublishAdForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if form.Amount <= 0 {
		response.WithError(ctx, http.StatusBadRequest, message.StakeAmountInvalid)
		return
	}

	ad := models.Ad{
		Advertiser: ctx.GetString(constants.CtxAddress),
		Amount:     form.Amount,
		CreatedAt:  time.Now(),
	}

	adId, err := controller.adService.CreateAd(ad)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.AdPublished, models.PublishAdResponse{
		AdId: adId,
	})
}

func (controller AdController) DeleteAd(ctx *gin.Context) {
	adId := ctx.Param("id")

	if len(adId) == 0 {
		response.WithError(ctx, http.StatusBadRequest, message.MissingId)
		return
	}

	err := controller.adService.DeleteAdById(adId)
	if err != nil {
		// TODO: Update error handling
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.AdDeleted, nil)
}

func (controller AdController) EditAd(ctx *gin.Context) {

}

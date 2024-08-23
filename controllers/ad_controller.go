package controllers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/AkifhanIlgaz/credible-mandela-api/models"
	"github.com/AkifhanIlgaz/credible-mandela-api/services"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/constants"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/message"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
	ads, err := controller.adService.GetAds()
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.SomethingWentWrong)
		return
	}

	// ? Ad sayısı 0 ise farklı mesaj göster
	response.WithSuccess(ctx, http.StatusOK, message.AdFound, ads)
}

func (controller AdController) GetAdById(ctx *gin.Context) {
	adId := ctx.Param(constants.ParamId)

	if len(adId) == 0 {
		response.WithError(ctx, http.StatusBadRequest, message.MissingId)
		return
	}

	ad, err := controller.adService.GetById(adId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			response.WithError(ctx, http.StatusNotFound, message.AdNotFound)
			return
		}
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.SomethingWentWrong)
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.AdFound, ad)
}

func (controller AdController) GetAdsByAddress(ctx *gin.Context) {
	address := ctx.Param(constants.ParamAddress)

	ads, err := controller.adService.GetAdsByAddress(address)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.SomethingWentWrong)
		return
	}

	// ? Ad sayısı 0 ise farklı mesaj göster
	response.WithSuccess(ctx, http.StatusOK, message.AdFound, ads)
}

func (controller AdController) GetAdsOfCurrentUser(ctx *gin.Context) {
	address := ctx.GetString(constants.CtxAddress)

	ads, err := controller.adService.GetAdsByAddress(address)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.SomethingWentWrong)
		return
	}

	// ? Ad sayısı 0 ise farklı mesaj göster
	response.WithSuccess(ctx, http.StatusOK, message.AdFound, ads)
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

	adId, err := controller.adService.Create(ad)
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
	adId := ctx.Param(constants.ParamId)

	if len(adId) == 0 {
		response.WithError(ctx, http.StatusBadRequest, message.MissingId)
		return
	}

	ad, err := controller.adService.GetById(adId)
	if err != nil {
		// TODO: Update error handling
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if ad.Advertiser != ctx.GetString(constants.CtxAddress) {
		response.WithError(ctx, http.StatusUnauthorized, message.NotAuthorizedToDelete)
		return
	}

	err = controller.adService.DeleteById(adId)
	if err != nil {
		// TODO: Update error handling
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.AdDeleted, nil)
}

func (controller AdController) UpdateAd(ctx *gin.Context) {
	adId := ctx.Param(constants.ParamId)

	if len(adId) == 0 {
		response.WithError(ctx, http.StatusBadRequest, message.MissingId)
		return
	}

	var form models.UpdateAdForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if form.NewAmount <= 0 {
		response.WithError(ctx, http.StatusBadRequest, message.StakeAmountInvalid)
		return
	}

	err := controller.adService.UpdateById(adId, form.NewAmount)
	if err != nil {
		// TODO: Update error handling
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.AdEdited, nil)
}

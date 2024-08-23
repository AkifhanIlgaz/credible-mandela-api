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

type CommunityNoteController struct {
	communityNoteService services.CommunityNoteService
}

func NewCommunityNoteController(communityNoteService services.CommunityNoteService) CommunityNoteController {
	return CommunityNoteController{
		communityNoteService: communityNoteService,
	}
}

func (controller CommunityNoteController) Publish(ctx *gin.Context) {
	var form models.PublishCommunityNoteForm

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

	communityNote := models.CommunityNote{
		Publisher:          ctx.GetString(constants.CtxUsername),
		CreatedAt:          time.Now(),
		Title:              form.Title,
		Content:            form.Content,
		CoverImageIPFSHash: form.CoverImageIPFSHash,
	}

	communityNoteId, err := controller.communityNoteService.Create(communityNote)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.CommunityNotePublished, models.PublishCommunityNoteResponse{
		CommunityNoteId: communityNoteId,
	})
}

func (controller CommunityNoteController) Delete(ctx *gin.Context) {
	communityNoteId := ctx.Param(constants.ParamId)

	if len(communityNoteId) == 0 {
		response.WithError(ctx, http.StatusBadRequest, message.MissingId)
		return
	}

	communityNote, err := controller.communityNoteService.GetById(communityNoteId)
	if err != nil {
		// TODO: Update error handling
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if communityNote.Publisher != ctx.GetString(constants.CtxUsername) {
		response.WithError(ctx, http.StatusUnauthorized, message.NotAuthorizedToDelete)
		return
	}

	err = controller.communityNoteService.DeleteById(communityNoteId)
	if err != nil {
		// TODO: Update error handling
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.CommunityNoteDeleted, nil)
}

func (controller CommunityNoteController) Like(ctx *gin.Context) {

}

func (controller CommunityNoteController) Unlike(ctx *gin.Context) {

}

func (controller CommunityNoteController) GetById(ctx *gin.Context) {
	communityNoteId := ctx.Param(constants.ParamId)

	if len(communityNoteId) == 0 {
		response.WithError(ctx, http.StatusBadRequest, message.MissingId)
		return
	}

	communityNote, err := controller.communityNoteService.GetById(communityNoteId)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.CommunityNoteFound, communityNote)
}

func (controller CommunityNoteController) GetAll(ctx *gin.Context) {
	communityNotes, err := controller.communityNoteService.GetAll()
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.SomethingWentWrong)
		return
	}

	// ? Ad sayısı 0 ise farklı mesaj göster
	response.WithSuccess(ctx, http.StatusOK, message.CommunityNoteFound, communityNotes)
}

func (controller CommunityNoteController) GetNotesOfUser(ctx *gin.Context) {
	username := ctx.Param(constants.CtxUsername)

	communityNotes, err := controller.communityNoteService.GetByUsername(username)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.SomethingWentWrong)
		return
	}

	// ? Ad sayısı 0 ise farklı mesaj göster
	response.WithSuccess(ctx, http.StatusOK, message.CommunityNoteFound, communityNotes)
}

func (controller CommunityNoteController) GetNotesOfCurrentUser(ctx *gin.Context) {
	username := ctx.GetString(constants.CtxUsername)

	communityNotes, err := controller.communityNoteService.GetByUsername(username)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.SomethingWentWrong)
		return
	}

	// ? Ad sayısı 0 ise farklı mesaj göster
	response.WithSuccess(ctx, http.StatusOK, message.CommunityNoteFound, communityNotes)
}

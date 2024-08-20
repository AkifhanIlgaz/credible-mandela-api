package controllers

import "github.com/AkifhanIlgaz/credible-mandela-api/services"

type CommunityNoteController struct {
	communityNoteService services.CommunityNoteService
}

func NewCommunityNoteController(communityNoteService services.CommunityNoteService) CommunityNoteController {
	return CommunityNoteController{
		communityNoteService: communityNoteService,
	}
}

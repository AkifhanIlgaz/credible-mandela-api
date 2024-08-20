package routers

import (
	"github.com/AkifhanIlgaz/credible-mandela-api/controllers"
	"github.com/gin-gonic/gin"
)

const CommunityNotesPath = "/community-notes"

type CommunityNoteRouter struct {
	CommunityNoteController controllers.CommunityNoteController
}

func NewCommunityNoteRouter(communityNoteController controllers.CommunityNoteController) CommunityNoteRouter {
	return CommunityNoteRouter{
		CommunityNoteController: communityNoteController,
	}
}

func (r CommunityNoteRouter) Setup(rg *gin.RouterGroup) {
	router := rg.Group(CommunityNotesPath)

	_ = router
}

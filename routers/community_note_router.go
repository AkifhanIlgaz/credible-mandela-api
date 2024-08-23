package routers

import (
	"github.com/AkifhanIlgaz/credible-mandela-api/controllers"
	"github.com/AkifhanIlgaz/credible-mandela-api/middlewares"
	"github.com/gin-gonic/gin"
)

const CommunityNotesPath = "/community-notes"

type CommunityNoteRouter struct {
	communityNoteController controllers.CommunityNoteController
	authMiddleware          middlewares.AuthMiddleware
}

func NewCommunityNoteRouter(communityNoteController controllers.CommunityNoteController, authMiddleware middlewares.AuthMiddleware) CommunityNoteRouter {
	return CommunityNoteRouter{
		communityNoteController: communityNoteController,
		authMiddleware:          authMiddleware,
	}
}

func (r CommunityNoteRouter) Setup(rg *gin.RouterGroup) {
	router := rg.Group(CommunityNotesPath)

	router.Use(r.authMiddleware.ExtractUidFromAuthHeader())

	router.POST("/", r.communityNoteController.Publish)          // Publish a community note
	router.POST("/like/:id", r.communityNoteController.Like)     // Like a community note
	router.POST("/unlike/:id", r.communityNoteController.Unlike) // Undo like a community note
	router.DELETE("/:id", r.communityNoteController.Delete)      // Delete community note by id

	router.GET("/:id", r.communityNoteController.GetById)                   // Get community note by ID
	router.GET("/", r.communityNoteController.GetById)                      // Get all community notes + Sort
	router.GET("/user/:address", r.communityNoteController.GetNotesOfUser)  // Get all community notes of user + Sort
	router.GET("/user/me", r.communityNoteController.GetNotesOfCurrentUser) // Get community notes of current user + Sort

}

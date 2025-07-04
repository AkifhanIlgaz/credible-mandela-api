package main

import (
	"context"
	"fmt"
	"log"

	cfg "github.com/AkifhanIlgaz/credible-mandela-api/config"
	"github.com/AkifhanIlgaz/credible-mandela-api/controllers"
	"github.com/AkifhanIlgaz/credible-mandela-api/middlewares"
	"github.com/AkifhanIlgaz/credible-mandela-api/routers"
	"github.com/AkifhanIlgaz/credible-mandela-api/services"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/constants"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/db"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/mande"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := cfg.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.TODO()

	client, err := db.Connect(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	db := client.Database(db.DatabaseName)

	mandeClient := mande.NewClient()

	adService := services.NewAdService(ctx, db)
	authService := services.NewAuthService(ctx, db)
	tokenService := services.NewTokenService(config)
	communityNoteService := services.NewCommunityNoteService(ctx, db)

	// TODO: optimize with separate function, interface
	err = authService.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	adController := controllers.NewAdController(adService)
	authController := controllers.NewAuthController(authService, tokenService, mandeClient)
	communityNoteController := controllers.NewCommunityNoteController(communityNoteService)

	authMiddleware := middlewares.NewAuthMiddleware(authService, tokenService)

	adRouter := routers.NewAdRouter(adController, authMiddleware)
	authRouter := routers.NewAuthRouter(authController, authMiddleware)
	communityNoteRouter := routers.NewCommunityNoteRouter(communityNoteController, authMiddleware)

	server := gin.Default()
	setCors(server)

	router := server.Group("/api")

	adRouter.Setup(router)
	authRouter.Setup(router)
	communityNoteRouter.Setup(router)

	err = server.Run(fmt.Sprintf(":%v", constants.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func setCors(server *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
}

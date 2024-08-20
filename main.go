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
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
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
	tokenService := services.NewTokenService(ctx, db, config)

	// TODO: optimize with separate function, interface
	err = authService.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	err = tokenService.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	adController := controllers.NewAdController(adService)
	authController := controllers.NewAuthController(authService, tokenService, mandeClient)

	authMiddleware := middlewares.NewAuthMiddleware(authService, tokenService)

	adRouter := routers.NewAdRouter(adController)
	authRouter := routers.NewAuthRouter(authController, authMiddleware)

	server := gin.Default()
	setCors(server)

	router := server.Group("/api")

	adRouter.Setup(router)
	authRouter.Setup(router)

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

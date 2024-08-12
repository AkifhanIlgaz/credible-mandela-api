package main

import (
	"context"
	"fmt"
	"log"

	cfg "github.com/AkifhanIlgaz/credible-mandela-api/config"
	"github.com/AkifhanIlgaz/credible-mandela-api/controllers"
	"github.com/AkifhanIlgaz/credible-mandela-api/routers"
	"github.com/AkifhanIlgaz/credible-mandela-api/services"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/constants"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/database"
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

	client, err := database.Connect(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	db := client.Database(constants.DatabaseName)

	adService := services.NewAdService(ctx, db)
	authService := services.NewAuthService(ctx, db)

	adController := controllers.NewAdController(adService)
	authController := controllers.NewAuthController(authService)

	adRouter := routers.NewAdRouter(adController)
	authRouter := routers.NewAuthRouter(authController)

	server := gin.Default()
	setCors(server)

	router := server.Group("/api")

	adRouter.Setup(router)
	authRouter.Setup(router)

	err = server.Run(fmt.Sprintf(":%v", config.Port))
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

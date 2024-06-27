package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AkifhanIlgaz/credible-mandela-api/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	godotenv.Load()

	ctx := context.TODO()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	adController, err := controllers.NewAdController(ctx, client.Database("credible-mandela"))
	if err != nil {
		panic(err)
	}

	server := gin.Default()
	setCors(server)

	router := server.Group("/api")

	router.POST("/ad", adController.CreateAd)
	router.GET("/ad", adController.GetAd)
	router.GET("/ads", adController.GetAllAds)

	err = server.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}

func setCors(server *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"https://www.crediblemandela.xyz"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
}

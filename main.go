// VERSÃO BETA 1.3
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Jp-Roberto/challengerklever/controllers"
	"github.com/Jp-Roberto/challengerklever/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	ctx            context.Context
	coinservice    services.CoinService
	coincontroller controllers.CoinController
	coincollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	godotenv.Load(".env")
	ctx = context.TODO()

	mongodbURI := os.Getenv("MONGODB_URL")
	fmt.Println("mongodbURI", mongodbURI)
	if mongodbURI == "" {
		mongodbURI = "mongodb://localhost:27017"
	}

	mongoconn := options.Client().ApplyURI(mongodbURI)
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo connection established")

	coincollection = (*mongo.Collection)(mongoclient.Database("voteSystem").Collection("coins"))
	coinservice = services.NewCoinService(coincollection, ctx)
	coincontroller = controllers.NewCoinController(coinservice)

	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	port := os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}

	basepath := server.Group("/beta")
	coincontroller.RegisterCoinRoutes(basepath)
	log.Fatal(server.Run(":" + port))
}

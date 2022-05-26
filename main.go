package main

import (
	"context"
	"log"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"example/bookmark-api/controllers"
	"example/bookmark-api/services"

	"github.com/gin-gonic/gin"
)

var (
	ctx 	context.Context
	err 	error

	bms 	services.BookmarkService
	bmc 	controllers.BookmarkController

	client *mongo.Client
	bookmarkcoll *mongo.Collection
	server *gin.Engine
)

func init() {
	ctx = context.TODO()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected successfully")

	bookmarkcoll = client.Database("bookmarks").Collection("entries")
	bms = services.NewBookmarkService(bookmarkcoll, ctx)
	bmc = controllers.New(bms)

	server = gin.Default()
}

func main() {
	defer client.Disconnect(ctx)

	defaultpath := server.Group("/v1")
	bmc.RegisterRoutes(defaultpath)

	log.Fatal(server.Run(":8080"))
}
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"godowns/controllers"
	"godowns/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Mongo error : ", err)
	}

	db := client.Database("krishi")
	fmt.Println("🚀 connected to mongodb : krishi")

	controllers.InitGodownController(db)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5174"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	routes.GodownRoutes(r)

	fmt.Println("🚀 Server running on port 8010")
	r.Run(":8010")
}

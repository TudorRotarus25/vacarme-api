package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/TudorRotarus25/vacarme-api/controllers"
)

func getMongoConnectionString() string {
	return os.Getenv("MONGO_URL")
}

func getMongoClient() *mongo.Client {
	connectionString := getMongoConnectionString()
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

func main() {
	godotenv.Load()

	r := gin.Default()
	mongoClient := getMongoClient()
	db := mongoClient.Database("vacarme")

	pc := controllers.NewProjectController(db)

	r.GET("/projects", pc.GetAllProjects)
	r.GET("/projects/:slug", pc.GetProject)
	r.GET("/categories", pc.GetAllCategories)

	fmt.Println("Listening on port 8080")

	err := r.Run()

	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
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
	r := httprouter.New()
	mongoClient := getMongoClient()
	db := mongoClient.Database("vacarme")

	pc := controllers.NewProjectController(db)

	r.GET("/projects", pc.GetAllProjects)
	r.GET("/projects/:slug", pc.GetProject)
	r.GET("/categories", pc.GetAllCategories)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}

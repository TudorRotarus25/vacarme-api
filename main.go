package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TudorRotarus25/vacarme-api/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := httprouter.New()
	mongoClient := getMongoClient()
	db := mongoClient.Database("vacarme")

	pc := controllers.NewProjectController(db)

	r.GET("/projects", pc.GetAllProjects)
	r.GET("/projects/:slug", pc.GetProject)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}

func getMongoClient() *mongo.Client {
	mongoUrl := os.Getenv("MONGO_URL")
	mongoUsername := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASS")

	clientOptions := options.Client().ApplyURI("mongodb+srv://" + mongoUsername + ":" + mongoPassword + "@" + mongoUrl + "?retryWrites=true&w=majority")

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

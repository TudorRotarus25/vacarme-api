package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/TudorRotarus25/vacarme-api/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := httprouter.New()
	mongoClient := getMongoClient()
	db := mongoClient.Database("vacarme")
	// defer mongoClient.Disconnect(context.TODO())

	pc := controllers.NewProjectController(db)

	r.GET("/projects", pc.GetAllProjects)
	r.GET("/projects/insert", pc.AddProject)

	http.ListenAndServe(":8080", r)
}

func getMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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

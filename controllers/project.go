package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/TudorRotarus25/vacarme-api/models"
	"github.com/julienschmidt/httprouter"
)

// ProjectController struct
type ProjectController struct {
	db *mongo.Database
}

// NewProjectController constructor
func NewProjectController(db *mongo.Database) *ProjectController {
	return &ProjectController{db: db}
}

// GetAllProjects list all the projects
func (pc ProjectController) GetAllProjects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	collection := pc.db.Collection("projects")
	filter := bson.D{{}}
	findOptions := options.Find()
	findOptions.SetLimit(500)

	var results []*models.ProjectModel

	cur, err := collection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.ProjectModel
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	if len(results) == 0 {
		w.Write([]byte("[]"))
		return
	}

	cur.Close(context.TODO())

	jsonResults, err := json.Marshal(results)

	if err != nil {
		http.Error(w, "{\"message\":\"Error parsing result\"}", 500)
		return
	}

	w.Write(jsonResults)
}

// GetProject get project details
func (pc ProjectController) GetProject(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	slug := params.ByName("slug")

	collection := pc.db.Collection("projects")
	filter := bson.D{{"slug", slug}}
	findOptions := options.Find()
	findOptions.SetLimit(500)

	var result models.ProjectModel

	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		http.Error(w, "{\"message\":\"Project not found\"}", 404)
		return
	}

	jsonResult, err := json.Marshal(result)

	if err != nil {
		http.Error(w, "{\"message\":\"Error parsing result\"}", 500)
		return
	}

	w.Write(jsonResult)
}

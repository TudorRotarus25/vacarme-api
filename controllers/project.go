package controllers

import (
	"context"
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

func sendResponse(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err := w.Write(response)

	if err != nil {
		log.Fatal(err)
	}
}

// GetAllCategories return all the categories sorted by `order`
func (pc ProjectController) GetAllCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	collection := pc.db.Collection("categories")
	filter := bson.D{{}}
	findOptions := options.Find()
	findOptions.SetLimit(500)
	findOptions.SetSort(bson.D{{"order", 1}})

	var results []*models.CategoryModel

	cur, err := collection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.CategoryModel
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
		sendResponse(w, []byte("[]"))
		return
	}

	cur.Close(context.TODO())

	categories, err := models.ParseCategories(results)

	sendResponse(w, categories)
}

// GetAllProjects list all the projects
func (pc ProjectController) GetAllProjects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	collection := pc.db.Collection("projects")
	filter := bson.D{{}}
	findOptions := options.Find()
	findOptions.SetLimit(500)
	findOptions.SetSort(bson.D{{"order", 1}})

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
		sendResponse(w, []byte("[]"))
		return
	}

	cur.Close(context.TODO())

	projects, err := models.ParseProjectsListBasicInfo(results)

	sendResponse(w, projects)
}

// GetProject get project details
func (pc ProjectController) GetProject(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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

	projectData, err := result.ParseProjectDetails()

	if err != nil {
		http.Error(w, "{\"message\":\"Error parsing result\"}", 500)
		return
	}

	sendResponse(w, projectData)
}

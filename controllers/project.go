package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/TudorRotarus25/vacarme-api/models"
	"github.com/julienschmidt/httprouter"
)

// ProjectController struct
type ProjectController struct {
	db *mongo.Database
}

// NewProjectController constructor
func NewProjectController(db *mongo.Database) *ProjectController {
	return &ProjectController{}
}

// GetAllProjects list all the projects
func (pc ProjectController) GetAllProjects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "{\"ok\": true}")
}

// AddProject adds a project to the DB
func (pc ProjectController) AddProject(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := models.ProjectModel{
		Slug:      "test",
		ImagePath: "other test",
	}
	insertResult, err := pc.db.Collection("projects").InsertOne(context.TODO(), p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, insertResult.InsertedID)
}

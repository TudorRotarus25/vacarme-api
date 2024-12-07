package controllers

import (
	"context"
	"github.com/TudorRotarus25/vacarme-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ProjectController struct
type ProjectController struct {
	db *mongo.Database
}

// NewProjectController constructor
func NewProjectController(db *mongo.Database) *ProjectController {
	return &ProjectController{db: db}
}

// GetAllCategories return all the categories sorted by `order`
func (pc ProjectController) GetAllCategories(c *gin.Context) {
	collection := pc.db.Collection("categories")
	filter := bson.D{{}}
	findOptions := options.Find()
	findOptions.SetLimit(500)
	findOptions.SetSort(bson.D{{"order", 1}})

	var results []*models.CategoryModel

	cur, err := collection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		print("Failed to load the categories collection")
		c.JSON(500, gin.H{"error": "Failed to load the categories collection"})
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var elem models.CategoryModel

		err := cur.Decode(&elem)
		if err != nil {
			print("Failed to parse collection item")
			continue
		}

		results = append(results, &elem)
	}

	categories := models.ParseCategories(results)

	c.JSON(200, gin.H{"data": categories})
}

// GetAllProjects list all the projects
func (pc ProjectController) GetAllProjects(c *gin.Context) {
	collection := pc.db.Collection("projects")
	filter := bson.D{{}}
	findOptions := options.Find()
	findOptions.SetLimit(500)
	findOptions.SetSort(bson.D{{"order", 1}})

	var results []*models.ProjectModel

	cur, err := collection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		print("Failed to load the projects collection")
		c.JSON(500, gin.H{"error": "Failed to load the projects collection"})
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var elem models.ProjectModel
		err := cur.Decode(&elem)
		if err != nil {
			print("Failed to parse collection item")
			continue
		}

		results = append(results, &elem)
	}

	projects := models.ParseProjectsListBasicInfo(results)

	c.JSON(200, gin.H{"data": projects})
}

// GetProject get project details
func (pc ProjectController) GetProject(c *gin.Context) {
	slug := c.Param("slug")

	collection := pc.db.Collection("projects")
	filter := bson.D{{"slug", slug}}
	findOptions := options.Find()
	findOptions.SetLimit(500)

	var result models.ProjectModel

	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		c.JSON(404, gin.H{"error": "Project not found"})
		return
	}

	projectData := result.ParseProjectDetails()

	c.JSON(200, gin.H{"data": projectData})
}

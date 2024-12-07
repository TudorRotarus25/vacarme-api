package controllers

import (
	"context"
	"github.com/TudorRotarus25/vacarme-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProjectController struct {
	db *mongo.Database
}

func NewProjectController(db *mongo.Database) *ProjectController {
	return &ProjectController{db: db}
}

func (pc ProjectController) GetAllCategories(c *gin.Context) {
	collection := pc.db.Collection("categories")
	filter := bson.D{{}}
	findOptions := options.Find()
	findOptions.SetLimit(500)
	findOptions.SetSort(bson.D{{"order", 1}})

	c.Header("Access-Control-Allow-Origin", "*")

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

	c.JSON(200, categories)
}

func (pc ProjectController) GetAllProjects(c *gin.Context) {
	collection := pc.db.Collection("projects")
	filter := bson.D{{}}
	findOptions := options.Find()
	findOptions.SetLimit(500)
	findOptions.SetSort(bson.D{{"order", 1}})

	c.Header("Access-Control-Allow-Origin", "*")

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

	c.JSON(200, projects)
}

func (pc ProjectController) GetProject(c *gin.Context) {
	slug := c.Param("slug")

	collection := pc.db.Collection("projects")
	filter := bson.D{{"slug", slug}}
	findOptions := options.Find()
	findOptions.SetLimit(500)

	c.Header("Access-Control-Allow-Origin", "*")

	var result models.ProjectModel

	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		c.JSON(404, gin.H{"error": "Project not found"})
		return
	}

	projectData := result.ParseProjectDetails()

	c.JSON(200, projectData)
}

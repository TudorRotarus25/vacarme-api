package models

import "strings"

type CategoryModel struct {
	Name  string `json:"name" bson:"name"`
	Order int    `json:"order" bson:"order"`
}

type CategoryResponse struct {
	Name string `json:"name"`
}

func (c *CategoryModel) ParseCategory() *CategoryResponse {
	response := CategoryResponse{
		Name: strings.ToLower(c.Name),
	}

	return &response
}

func ParseCategories(c []*CategoryModel) []*CategoryResponse {
	var categories []*CategoryResponse

	for _, cur := range c {
		pp := cur.ParseCategory()

		categories = append(categories, pp)
	}

	return categories
}

package models

import "strings"

// CategoryModel mongo model for categories
type CategoryModel struct {
	Name  string `json:"name" bson:"name"`
	Order int    `json:"order" bson:"order"`
}

type CategoryResponse struct {
	Name string `json:"name"`
}

// ParseCategory parse category
func (c *CategoryModel) ParseCategory() *CategoryResponse {
	response := CategoryResponse{
		Name: strings.ToLower(c.Name),
	}

	return &response
}

// ParseCategories parse a list of categories
func ParseCategories(c []*CategoryModel) []*CategoryResponse {
	var categories []*CategoryResponse

	for _, cur := range c {
		pp := cur.ParseCategory()

		categories = append(categories, pp)
	}

	return categories
}

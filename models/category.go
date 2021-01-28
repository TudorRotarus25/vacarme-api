package models

import (
	"encoding/json"
	"strings"
)

// CategoryModel mongo model for categories
type CategoryModel struct {
	Name  string `json:"name" bson:"name"`
	Order int    `json:"order" bson:"order"`
}

// ParseCategory parse category
func (c *CategoryModel) ParseCategory() ([]byte, error) {
	type responseType struct {
		Name string `json:"name"`
	}

	response := responseType{
		Name: strings.ToLower(c.Name),
	}

	return json.Marshal(response)
}

// ParseCategories parse a list of categories
func ParseCategories(c []*CategoryModel) ([]byte, error) {
	var categories []string

	for _, cur := range c {
		pp, err := cur.ParseCategory()

		if err != nil {
			return nil, err
		}

		categories = append(categories, string(pp))
	}

	projects := "[" + strings.Join(categories, ",") + "]"

	return []byte(projects), nil
}

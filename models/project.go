package models

// ProjectModel mongo model for projects
type ProjectModel struct {
	Slug          string  `json:"slug"`
	Name          string  `json:"name"`
	Order         int16   `json:"order"`
	Layout        string  `json:"layout"`
	Width         float32 `json:"width"`
	ShapeURL      string  `json:"shapeUrl"`
	HoverImageURL string  `json:"hoverImageUrl"`
	Cta           string  `json:"cta"`
	HomeText1     string  `json:"homeText1"`
	HomeText2     string  `json:"homeText2"`
}

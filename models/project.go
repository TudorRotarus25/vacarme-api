package models

// ProjectModel mongo model for projects
type ProjectModel struct {
	Slug      string `json:"slug"`
	ImagePath string `json:"imagePath"`
}

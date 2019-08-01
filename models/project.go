package models

import (
	"encoding/json"
)

// ProjectModel mongo model for projects
type ProjectModel struct {
	Order          int16    `json:"order" bson:"order"`
	Name           string   `json:"name" bson:"name"`
	Slug           string   `json:"slug" bson:"slug"`
	Project        string   `json:"project" bson:"project"`
	Domain         string   `json:"domain" bson:"domain"`
	Year           int16    `json:"year" bson:"year"`
	Client         string   `json:"client" bson:"client"`
	Layout         string   `json:"layout" bson:"layout"`
	TaglineAngle   int16    `json:"taglineAngle" bson:"taglineAngle"`
	TaglineTop     string   `json:"taglineTop" bson:"taglineTop"`
	TaglineBottom  string   `json:"taglineBottom" bson:"taglineBottom"`
	Width          float32  `json:"width" bson:"width"`
	ShapeURL       string   `json:"shapeUrl" bson:"shapeUrl"`
	HoverImageURL  string   `json:"hoverImageUrl" bson:"hoverImageUrl"`
	MobileImageURL string   `json:"mobileImageUrl" bson:"mobileImageUrl"`
	Cta            string   `json:"cta" bson:"cta"`
	Paragraphs     []string `json:"paragraphs" bson:"paragraphs"`
}

// ParseProjectDetails parse project leaving only the details
func (project *ProjectModel) ParseProjectDetails() ([]byte, error) {
	type responseType struct {
		Name       string   `json:"name"`
		Project    string   `json:"project"`
		Domain     string   `json:"domain"`
		Year       int16    `json:"year"`
		Client     string   `json:"client"`
		Paragraphs []string `json:"paragraphs"`
	}

	response := responseType{
		Name:       project.Name,
		Project:    project.Project,
		Domain:     project.Domain,
		Year:       project.Year,
		Client:     project.Client,
		Paragraphs: project.Paragraphs,
	}

	return json.Marshal(response)
}

// ParseProjectBasicInfo parse project leaving only basic info
func (project *ProjectModel) ParseProjectBasicInfo() ([]byte, error) {
	type responseType struct {
		Slug           string  `json:"slug"`
		Layout         string  `json:"layout"`
		TaglineAngle   int16   `json:"taglineAngle"`
		TaglineTop     string  `json:"taglineTop"`
		TaglineBottom  string  `json:"taglineBottom"`
		Width          float32 `json:"width"`
		ShapeURL       string  `json:"shapeUrl"`
		HoverImageURL  string  `json:"hoverImageUrl"`
		MobileImageURL string  `json:"mobileImageUrl"`
		Cta            string  `json:"cta"`
	}

	response := responseType{
		Slug:           project.Slug,
		Layout:         project.Layout,
		TaglineAngle:   project.TaglineAngle,
		TaglineTop:     project.TaglineTop,
		TaglineBottom:  project.TaglineBottom,
		Width:          project.Width,
		ShapeURL:       project.ShapeURL,
		HoverImageURL:  project.HoverImageURL,
		MobileImageURL: project.MobileImageURL,
		Cta:            project.Cta,
	}

	return json.Marshal(response)
}

// ParseProjectsListBasicInfo parse a list of projects to only contain basic info
func ParseProjectsListBasicInfo(p []*ProjectModel) ([]byte, error) {
	p0, err := p[0].ParseProjectBasicInfo()

	if err != nil {
		return nil, err
	}

	projects := "[" + string(p0)

	for i := 1; i < len(p); i++ {
		np, err := p[i].ParseProjectBasicInfo()

		if err != nil {
			return nil, err
		}

		projects += "," + string(np)
	}

	projects += "]"

	return []byte(projects), nil
}

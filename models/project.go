package models

import (
	"encoding/json"
	"strings"
)

// ProjectModel mongo model for projects
type ProjectModel struct {
	Order                int      `json:"order" bson:"order"`
	Name                 string   `json:"name" bson:"name"`
	Slug                 string   `json:"slug" bson:"slug"`
	Project              string   `json:"project" bson:"project"`
	Domain               string   `json:"domain" bson:"domain"`
	Year                 int      `json:"year" bson:"year"`
	Client               string   `json:"client" bson:"client"`
	Layout               string   `json:"layout" bson:"layout"`
	TaglineAngle         int      `json:"taglineAngle" bson:"taglineAngle"`
	TaglineTop           string   `json:"taglineTop" bson:"taglineTop"`
	TaglineBottom        string   `json:"taglineBottom" bson:"taglineBottom"`
	Width                float32  `json:"width" bson:"width"`
	ShapeURL             string   `json:"shapeUrl" bson:"shapeUrl"`
	HoverImageURL        string   `json:"hoverImageUrl" bson:"hoverImageUrl"`
	MobileImageURL       string   `json:"mobileImageUrl" bson:"mobileImageUrl"`
	Color                string   `json:"color" bson:"color"`
	Cta                  string   `json:"cta" bson:"cta"`
	Paragraphs           []string `json:"paragraphs" bson:"paragraphs"`
	ProjectImagesPreset  int      `json:"projectImagesPreset" bson:"projectImagesPreset"`
	ProjectDesktopImages []string `json:"projectDesktopImages" bson:"projectDesktopImages"`
	ProjectMobileImages  []string `json:"projectMobileImages" bson:"projectMobileImages"`
}

// ParseProjectDetails parse project leaving only the details
func (project *ProjectModel) ParseProjectDetails() ([]byte, error) {
	type responseType struct {
		Name                 string   `json:"name"`
		Project              string   `json:"project"`
		Domain               string   `json:"domain"`
		Year                 int      `json:"year"`
		Client               string   `json:"client"`
		Paragraphs           []string `json:"paragraphs"`
		ProjectImagesPreset  int      `json:"projectImagesPreset"`
		ProjectDesktopImages []string `json:"projectDesktopImages"`
		ProjectMobileImages  []string `json:"projectMobileImages"`
	}

	response := responseType{
		Name:                 project.Name,
		Project:              project.Project,
		Domain:               project.Domain,
		Year:                 project.Year,
		Client:               project.Client,
		Paragraphs:           project.Paragraphs,
		ProjectImagesPreset:  project.ProjectImagesPreset,
		ProjectDesktopImages: project.ProjectDesktopImages,
		ProjectMobileImages:  project.ProjectMobileImages,
	}

	return json.Marshal(response)
}

// ParseProjectBasicInfo parse project leaving only basic info
func (project *ProjectModel) ParseProjectBasicInfo() ([]byte, error) {
	type responseType struct {
		Slug           string  `json:"slug"`
		Layout         string  `json:"layout"`
		TaglineAngle   int     `json:"taglineAngle"`
		TaglineTop     string  `json:"taglineTop"`
		TaglineBottom  string  `json:"taglineBottom"`
		Width          float32 `json:"width"`
		ShapeURL       string  `json:"shapeUrl"`
		HoverImageURL  string  `json:"hoverImageUrl"`
		MobileImageURL string  `json:"mobileImageUrl"`
		Color          string  `json:"color"`
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
		Color:          project.Color,
		Cta:            project.Cta,
	}

	return json.Marshal(response)
}

// ParseProjectsListBasicInfo parse a list of projects to only contain basic info
func ParseProjectsListBasicInfo(p []*ProjectModel) ([]byte, error) {
	var parsedProjects []string

	for _, cur := range p {
		pp, err := cur.ParseProjectBasicInfo()

		if err != nil {
			return nil, err
		}

		parsedProjects = append(parsedProjects, string(pp))
	}

	projects := "[" + strings.Join(parsedProjects, ",") + "]"

	return []byte(projects), nil
}

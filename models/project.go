package models

import (
	"encoding/json"
	"sort"
	"strings"
)

// Video type
type Video struct {
	Src         string `json:"src" bson:"src"`
	Orientation string `json:"orientation" bson:"orientation"`
}

// Media type
type Media struct {
	Type        string `json:"type" bson:"type"`
	Src         string `json:"src" bson:"src"`
	Orientation string `json:"orientation" bson:"orientation"`
	Order       int    `json:"order" bson:"order"`
}

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
	ThirdOverlayImageURL string   `json:"thirdOverlayImageUrl" bson:"thirdOverlayImageUrl"`
	MobileImageURL       string   `json:"mobileImageUrl" bson:"mobileImageUrl"`
	Color                string   `json:"color" bson:"color"`
	Cta                  string   `json:"cta" bson:"cta"`
	Paragraphs           []string `json:"paragraphs" bson:"paragraphs"`
	URL                  string   `json:"url" bson:"url"`
	ProjectImagesPreset  int      `json:"projectImagesPreset" bson:"projectImagesPreset"`
	ProjectDesktopImages []string `json:"projectDesktopImages" bson:"projectDesktopImages"`
	ProjectMobileImages  []string `json:"projectMobileImages" bson:"projectMobileImages"`
	ProjectVideos        []Video  `json:"projectVideos" bson:"projectVideos"`
	ProjectDesktopMedia  []Media  `json:"projectDesktopMedia" bson:"projectDesktopMedia"`
	ProjectMobileMedia   []Media  `json:"projectMobileMedia" bson:"projectMobileMedia"`
	DarkMode             bool     `json:"darkMode" bson:"darkMode"`
	TaglineDarkMode      bool     `json:"taglineDarkMode" bson:"taglineDarkMode"`
	MetaTitle            string   `json:"metaTitle" bson:"metaTitle"`
	MetaDescription      string   `json:"metaDescription" bson:"metaDescription"`
}

func sortMediaItems(items []Media) []Media {
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Order < items[j].Order
	})

	return items
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
		URL                  string   `json:"url"`
		ProjectImagesPreset  int      `json:"projectImagesPreset"`
		ProjectDesktopImages []string `json:"projectDesktopImages"`
		ProjectMobileImages  []string `json:"projectMobileImages"`
		ProjectVideos        []Video  `json:"projectVideos"`
		ProjectDesktopMedia  []Media  `json:"projectDesktopMedia"`
		ProjectMobileMedia   []Media  `json:"projectMobileMedia"`
		DarkMode             bool     `json:"darkMode"`
		MetaTitle            string   `json:"metaTitle"`
		MetaDescription      string   `json:"metaDescription"`
	}

	response := responseType{
		Name:                 project.Name,
		Project:              project.Project,
		Domain:               project.Domain,
		Year:                 project.Year,
		Client:               project.Client,
		Paragraphs:           project.Paragraphs,
		URL:                  project.URL,
		ProjectImagesPreset:  project.ProjectImagesPreset,
		ProjectDesktopImages: project.ProjectDesktopImages,
		ProjectMobileImages:  project.ProjectMobileImages,
		ProjectVideos:        project.ProjectVideos,
		ProjectDesktopMedia:  sortMediaItems(project.ProjectDesktopMedia),
		ProjectMobileMedia:   sortMediaItems(project.ProjectMobileMedia),
		DarkMode:             project.DarkMode,
		MetaTitle:            project.MetaTitle,
		MetaDescription:      project.MetaDescription,
	}

	return json.Marshal(response)
}

// ParseProjectBasicInfo parse project leaving only basic info
func (project *ProjectModel) ParseProjectBasicInfo() ([]byte, error) {
	type responseType struct {
		Name                 string  `json:"name"`
		Slug                 string  `json:"slug"`
		Layout               string  `json:"layout"`
		TaglineAngle         int     `json:"taglineAngle"`
		TaglineTop           string  `json:"taglineTop"`
		TaglineBottom        string  `json:"taglineBottom"`
		Width                float32 `json:"width"`
		ShapeURL             string  `json:"shapeUrl"`
		HoverImageURL        string  `json:"hoverImageUrl"`
		ThirdOverlayImageURL string  `json:"thirdOverlayImageUrl"`
		MobileImageURL       string  `json:"mobileImageUrl"`
		Color                string  `json:"color"`
		Cta                  string  `json:"cta"`
		TaglineDarkMode      bool    `json:"taglineDarkMode"`
	}

	response := responseType{
		Name:                 project.Name,
		Slug:                 project.Slug,
		Layout:               project.Layout,
		TaglineAngle:         project.TaglineAngle,
		TaglineTop:           project.TaglineTop,
		TaglineBottom:        project.TaglineBottom,
		Width:                project.Width,
		ShapeURL:             project.ShapeURL,
		HoverImageURL:        project.HoverImageURL,
		ThirdOverlayImageURL: project.ThirdOverlayImageURL,
		MobileImageURL:       project.MobileImageURL,
		Color:                project.Color,
		Cta:                  project.Cta,
		TaglineDarkMode:      project.TaglineDarkMode,
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

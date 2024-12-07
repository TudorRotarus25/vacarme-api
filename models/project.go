package models

import (
	"sort"
)

type Media struct {
	Type        string `json:"type" bson:"type"`
	Src         string `json:"src" bson:"src"`
	Orientation string `json:"orientation" bson:"orientation"`
	Order       int    `json:"order" bson:"order"`
}

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
	ProjectDesktopMedia  []Media  `json:"projectDesktopMedia" bson:"projectDesktopMedia"`
	ProjectMobileMedia   []Media  `json:"projectMobileMedia" bson:"projectMobileMedia"`
	DarkMode             bool     `json:"darkMode" bson:"darkMode"`
	TaglineDarkMode      bool     `json:"taglineDarkMode" bson:"taglineDarkMode"`
	MetaTitle            string   `json:"metaTitle" bson:"metaTitle"`
	MetaDescription      string   `json:"metaDescription" bson:"metaDescription"`
}

type ProjectBasicInfoResponse struct {
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

type ProjectDetailsResponse struct {
	Name                string   `json:"name"`
	Project             string   `json:"project"`
	Domain              string   `json:"domain"`
	Year                int      `json:"year"`
	Client              string   `json:"client"`
	Paragraphs          []string `json:"paragraphs"`
	URL                 string   `json:"url"`
	ProjectDesktopMedia []Media  `json:"projectDesktopMedia"`
	ProjectMobileMedia  []Media  `json:"projectMobileMedia"`
	DarkMode            bool     `json:"darkMode"`
	MetaTitle           string   `json:"metaTitle"`
	MetaDescription     string   `json:"metaDescription"`
}

func sortMediaItems(items []Media) []Media {
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Order < items[j].Order
	})

	return items
}

func (project *ProjectModel) ParseProjectDetails() *ProjectDetailsResponse {
	response := ProjectDetailsResponse{
		Name:                project.Name,
		Project:             project.Project,
		Domain:              project.Domain,
		Year:                project.Year,
		Client:              project.Client,
		Paragraphs:          project.Paragraphs,
		URL:                 project.URL,
		ProjectDesktopMedia: sortMediaItems(project.ProjectDesktopMedia),
		ProjectMobileMedia:  sortMediaItems(project.ProjectMobileMedia),
		DarkMode:            project.DarkMode,
		MetaTitle:           project.MetaTitle,
		MetaDescription:     project.MetaDescription,
	}

	return &response
}

func (project *ProjectModel) ParseProjectBasicInfo() *ProjectBasicInfoResponse {
	response := ProjectBasicInfoResponse{
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

	return &response
}

func ParseProjectsListBasicInfo(p []*ProjectModel) []*ProjectBasicInfoResponse {
	var parsedProjects []*ProjectBasicInfoResponse

	for _, cur := range p {
		pp := cur.ParseProjectBasicInfo()

		parsedProjects = append(parsedProjects, pp)
	}

	return parsedProjects
}

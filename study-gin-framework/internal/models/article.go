package models

type Article struct {
	Title       string `json: "title"`
	Description string `json: "description"`
	URL         string `json: "URL"`
}

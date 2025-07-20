package models

type Collection struct {
	ID      string  `json:"id"`
	SiteID  string  `json:"siteId"`
	Name    string  `json:"name"`
	Slug    string  `json:"slug"`
	Fields  []Field `json:"fields"`
	Entries []Entry `json:"entries"`
}

type CreateCollectionRequest struct {
	SiteID string  `json:"siteId"`
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

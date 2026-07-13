package models

import "time"

type MediaFile struct {
	ID           string    `json:"id"`
	SiteID       string    `json:"siteId"`
	OriginalName string    `json:"originalName"`
	MIMEType     string    `json:"mimeType"`
	Size         int64     `json:"size"`
	URL          string    `json:"url"`
	CreatedAt    time.Time `json:"createdAt"`
}

type MediaFilePage struct {
	Files      []MediaFile `json:"files"`
	Pagination Pagination  `json:"pagination"`
}

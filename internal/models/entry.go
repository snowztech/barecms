package models

import "encoding/json"

type Entry struct {
	ID           string          `json:"id"`
	CollectionID string          `json:"collection_id"`
	Data         json.RawMessage `json:"data"`
}

type CreateEntryRequest struct {
	CollectionID string          `json:"collectionId"`
	Data         json.RawMessage `json:"data"`
}

type UpdateEntryRequest struct {
	Data json.RawMessage `json:"data"`
}

type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"totalPages"`
}

type EntryPage struct {
	Entries    []Entry    `json:"entries"`
	Pagination Pagination `json:"pagination"`
}

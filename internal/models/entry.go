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

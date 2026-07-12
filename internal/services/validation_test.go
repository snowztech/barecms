package services

import (
	"barecms/internal/models"
	"encoding/json"
	"errors"
	"testing"
)

func TestValidateEntryData(t *testing.T) {
	fields := []models.Field{
		{Name: "title", Type: models.FieldTypeString},
		{Name: "rating", Type: models.FieldTypeNumber},
		{Name: "published", Type: models.FieldTypeBoolean},
		{Name: "date", Type: models.FieldTypeDate},
		{Name: "image", Type: models.FieldTypeImage, Optional: true},
	}
	valid := json.RawMessage(`{"title":{"value":"Post","type":"string"},"rating":{"value":"4.5","type":"number"},"published":{"value":"true","type":"boolean"},"date":{"value":"2026-07-12","type":"date"},"image":{"value":null,"type":"image"}}`)
	if err := validateEntryData(valid, fields); err != nil {
		t.Fatalf("valid entry rejected: %v", err)
	}

	invalid := json.RawMessage(`{"rating":{"value":"many","type":"number"},"published":{"value":"yes","type":"boolean"},"date":{"value":"tomorrow","type":"date"},"extra":{"value":"x","type":"string"}}`)
	err := validateEntryData(invalid, fields)
	var validationError *ValidationError
	if !errors.As(err, &validationError) {
		t.Fatalf("expected validation error, got %v", err)
	}
	for _, field := range []string{"title", "rating", "published", "date", "extra"} {
		if validationError.Fields[field] == "" {
			t.Errorf("expected error for %s", field)
		}
	}
}

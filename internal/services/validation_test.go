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

func TestValidateCollectionSchemaRejectsInvalidDefinitions(t *testing.T) {
	err := validateCollectionSchema("", []models.Field{
		{Name: "title", Type: models.FieldTypeString},
		{Name: "title", Type: "unknown"},
	})
	var validationError *ValidationError
	if !errors.As(err, &validationError) {
		t.Fatalf("expected validation error, got %v", err)
	}
	for _, field := range []string{"name", "fields.1.name", "fields.1.type"} {
		if validationError.Fields[field] == "" {
			t.Errorf("expected error for %s", field)
		}
	}
}

func TestValidateEntryDataEnforcesLengthAndRangeConstraints(t *testing.T) {
	minLength, maxLength := 3, 5
	min, max := 1.0, 10.0
	fields := []models.Field{
		{Name: "title", Type: models.FieldTypeString, MinLength: &minLength, MaxLength: &maxLength},
		{Name: "rating", Type: models.FieldTypeNumber, Min: &min, Max: &max},
	}
	invalid := json.RawMessage(`{"title":{"value":"ab","type":"string"},"rating":{"value":"11","type":"number"}}`)
	var validationError *ValidationError
	if !errors.As(validateEntryData(invalid, fields), &validationError) {
		t.Fatal("expected constraint validation error")
	}
	if validationError.Fields["title"] == "" || validationError.Fields["rating"] == "" {
		t.Fatalf("missing constraint errors: %+v", validationError.Fields)
	}
	valid := json.RawMessage(`{"title":{"value":"hello","type":"string"},"rating":{"value":"10","type":"number"}}`)
	if err := validateEntryData(valid, fields); err != nil {
		t.Fatalf("valid constrained entry rejected: %v", err)
	}
}

func TestValidateCollectionSchemaRejectsInvalidConstraints(t *testing.T) {
	minLength, maxLength := 5, 2
	min, max := 10.0, 1.0
	err := validateCollectionSchema("Posts", []models.Field{
		{Name: "title", Type: models.FieldTypeString, MinLength: &minLength, MaxLength: &maxLength},
		{Name: "rating", Type: models.FieldTypeNumber, Min: &min, Max: &max},
	})
	var validationError *ValidationError
	if !errors.As(err, &validationError) {
		t.Fatal("expected invalid constraints to be rejected")
	}
	if validationError.Fields["fields.0.maxLength"] == "" || validationError.Fields["fields.1.max"] == "" {
		t.Fatalf("missing schema errors: %+v", validationError.Fields)
	}
}

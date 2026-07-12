package services

import (
	"barecms/internal/models"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type ValidationError struct {
	Fields map[string]string
}

func (e *ValidationError) Error() string { return "entry validation failed" }

type entryFieldValue struct {
	Value any              `json:"value"`
	Type  models.FieldType `json:"type"`
}

func validateEntryData(data json.RawMessage, fields []models.Field) error {
	var values map[string]entryFieldValue
	if err := json.Unmarshal(data, &values); err != nil {
		return &ValidationError{Fields: map[string]string{"data": "must be a JSON object"}}
	}
	errors := make(map[string]string)
	schema := make(map[string]models.Field, len(fields))
	for _, field := range fields {
		schema[field.Name] = field
		value, exists := values[field.Name]
		if !exists || value.Value == nil || fmt.Sprint(value.Value) == "" {
			if !field.Optional {
				errors[field.Name] = "is required"
			}
			continue
		}
		if value.Type != field.Type {
			errors[field.Name] = "type does not match collection schema"
			continue
		}
		if message := validateFieldValue(field.Type, value.Value); message != "" {
			errors[field.Name] = message
		}
	}
	for name := range values {
		if _, exists := schema[name]; !exists {
			errors[name] = "is not defined in the collection schema"
		}
	}
	if len(errors) > 0 {
		return &ValidationError{Fields: errors}
	}
	return nil
}

func validateFieldValue(fieldType models.FieldType, value any) string {
	text, ok := value.(string)
	if !ok {
		return "must be a string value"
	}
	switch fieldType {
	case models.FieldTypeNumber:
		if _, err := strconv.ParseFloat(text, 64); err != nil {
			return "must be a number"
		}
	case models.FieldTypeBoolean:
		if _, err := strconv.ParseBool(text); err != nil {
			return "must be true or false"
		}
	case models.FieldTypeDate:
		if _, err := time.Parse("2006-01-02", text); err != nil {
			return "must use YYYY-MM-DD format"
		}
	case models.FieldTypeURL, models.FieldTypeImage:
		parsed, err := url.ParseRequestURI(text)
		if err != nil || (parsed.Scheme != "" && parsed.Scheme != "http" && parsed.Scheme != "https") {
			return "must be a valid URL"
		}
	}
	return ""
}

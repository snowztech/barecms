package services

import (
	"barecms/internal/models"
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
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
		if message := validateFieldValue(field, value.Value); message != "" {
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

func validateFieldValue(field models.Field, value any) string {
	text, ok := value.(string)
	if !ok {
		return "must be a string value"
	}
	switch field.Type {
	case models.FieldTypeNumber:
		number, err := strconv.ParseFloat(text, 64)
		if err != nil || math.IsNaN(number) || math.IsInf(number, 0) {
			return "must be a number"
		}
		if field.Min != nil && number < *field.Min {
			return fmt.Sprintf("must be at least %g", *field.Min)
		}
		if field.Max != nil && number > *field.Max {
			return fmt.Sprintf("must be at most %g", *field.Max)
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
	if field.MinLength != nil && len([]rune(text)) < *field.MinLength {
		return fmt.Sprintf("must contain at least %d characters", *field.MinLength)
	}
	if field.MaxLength != nil && len([]rune(text)) > *field.MaxLength {
		return fmt.Sprintf("must contain at most %d characters", *field.MaxLength)
	}
	return ""
}

func validateCollectionSchema(name string, fields []models.Field) error {
	problems := make(map[string]string)
	if strings.TrimSpace(name) == "" {
		problems["name"] = "is required"
	}
	if len(fields) == 0 {
		problems["fields"] = "must contain at least one field"
	}
	seen := make(map[string]bool, len(fields))
	for index, field := range fields {
		key := fmt.Sprintf("fields.%d", index)
		fieldName := strings.TrimSpace(field.Name)
		if fieldName == "" {
			problems[key+".name"] = "is required"
		} else if seen[fieldName] {
			problems[key+".name"] = "must be unique"
		}
		seen[fieldName] = true
		if !models.IsValidFieldType(string(field.Type)) {
			problems[key+".type"] = "is invalid"
		}
		stringLike := field.Type == models.FieldTypeString || field.Type == models.FieldTypeText || field.Type == models.FieldTypeURL
		if (field.MinLength != nil || field.MaxLength != nil) && !stringLike {
			problems[key+".length"] = "is only supported for string, text, and URL fields"
		}
		if field.MinLength != nil && *field.MinLength < 0 {
			problems[key+".minLength"] = "must be zero or greater"
		}
		if field.MaxLength != nil && *field.MaxLength < 0 {
			problems[key+".maxLength"] = "must be zero or greater"
		}
		if field.MinLength != nil && field.MaxLength != nil && *field.MinLength > *field.MaxLength {
			problems[key+".maxLength"] = "must be greater than or equal to minLength"
		}
		if (field.Min != nil || field.Max != nil) && field.Type != models.FieldTypeNumber {
			problems[key+".range"] = "is only supported for number fields"
		}
		if field.Min != nil && field.Max != nil && *field.Min > *field.Max {
			problems[key+".max"] = "must be greater than or equal to min"
		}
	}
	if len(problems) > 0 {
		return &ValidationError{Fields: problems}
	}
	return nil
}

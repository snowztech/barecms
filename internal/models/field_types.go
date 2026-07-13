package models

type FieldType string

const (
	FieldTypeString  FieldType = "string"
	FieldTypeText    FieldType = "text"
	FieldTypeNumber  FieldType = "number"
	FieldTypeBoolean FieldType = "boolean"
	FieldTypeDate    FieldType = "date"
	FieldTypeImage   FieldType = "image"
	FieldTypeURL     FieldType = "url"
)

func ValidFieldTypes() []FieldType {
	return []FieldType{
		FieldTypeString,
		FieldTypeText,
		FieldTypeNumber,
		FieldTypeBoolean,
		FieldTypeDate,
		FieldTypeImage,
		FieldTypeURL,
	}
}

func IsValidFieldType(fieldType string) bool {
	for _, validType := range ValidFieldTypes() {
		if string(validType) == fieldType {
			return true
		}
	}
	return false
}

type Field struct {
	Name      string    `json:"name"`
	Type      FieldType `json:"type"`
	Optional  bool      `json:"optional,omitempty"`
	MinLength *int      `json:"minLength,omitempty"`
	MaxLength *int      `json:"maxLength,omitempty"`
	Min       *float64  `json:"min,omitempty"`
	Max       *float64  `json:"max,omitempty"`
}

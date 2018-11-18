package gqlutils

import (
	"testing"

	"github.com/graphql-go/graphql"
)

func TestTypeToGQLTypeInt(t *testing.T) {
	intList := []string{"byte", "uintptr", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64"}

	for _, name := range intList {
		result := TypeToGQLType(name)
		if result != graphql.Int {
			t.Errorf("Expected 'graphql.Int', got '%T'", *result)
		}
	}
}

func TestTypeToGQLTypeFloat(t *testing.T) {
	floatList := []string{"float32", "float64", "complex64", "complex128"}
	for _, name := range floatList {
		result := TypeToGQLType(name)
		if result != graphql.Float {
			t.Errorf("Expected 'graphql.Float', got '%T'", result)
		}

	}
}

func TestTypeToGQLTypeBool(t *testing.T) {
	result := TypeToGQLType("bool")
	if result != graphql.Boolean {
		t.Errorf("Expected 'graphql.Boolean', got '%T'", result)
	}
}

func TestTypeToGQLTypeDateTime(t *testing.T) {
	dateTimeList := []string{"date", "time"}
	for _, name := range dateTimeList {
		result := TypeToGQLType(name)
		if result != graphql.DateTime {
			t.Errorf("Expected 'graphql.DateTime', got '%T'", result)
		}
	}
}

func TestTypeToGQLTypeString(t *testing.T) {
	strList := []string{"string", "rune"}
	for _, name := range strList {
		result := TypeToGQLType(name)
		if result != graphql.String {
			t.Errorf("Expected 'graphql.String', got '%T'", result)
		}
	}
}

func TestTypeToGQLTypeUnknown(t *testing.T) {
	result := TypeToGQLType("foo")
	if result != graphql.String {
		t.Errorf("Expected 'graphql.String', got '%T'", result)
	}
}

func TestStructToFields(t *testing.T) {

}

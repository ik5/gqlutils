package gqlutils

import (
	"testing"
	"time"

	"github.com/graphql-go/graphql"
)

type test struct {
	ID       int       `json:"id,omitempty" desc:"foo bar" type:"int64"`
	Name     string    `json:"-"`
	DateTime time.Time `json:"date_time" type:"time" desc:"DateTime"`
	KeyName  int64     `json:"-," deprecation:"Will not be used next version"`
	Nothing  string
}

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

func TestStructToFieldsNames(t *testing.T) {
	fields := StructToFields(test{})

	keyNames := []string{"id", "date_time", "key_name"}
	if len(keyNames) != len(fields) {
		t.Errorf("Expected '#%d' fields, found '#%d'", len(keyNames), len(fields))
	}
}

func TestStructToFieldsFieldsValidTypes(t *testing.T) {
	fields := StructToFields(test{})

	if fields["id"].Type != graphql.Int {
		t.Errorf("Expected 'id.Type == graphql.Int' found: '%T'", fields["id"].Type)
	}

	if fields["date_time"].Type != graphql.DateTime {
		t.Errorf("Expected 'date_time.Type == graphql.DateTime', found '%T'", fields["date_time"].Type)
	}

	if fields["key_name"].Type != graphql.Int {
		t.Errorf("Expected 'key_name.Type == graphql.Int', found '%T'", fields["date_time"].Type)
	}
}

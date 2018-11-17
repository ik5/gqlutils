package gqlutils

import (
	"reflect"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/ik5/gostrutils"
)

// TypeToGQLType takes a string with a type and converts it into graphql Scalar.
// If unknown type is found, it returns graphql.String.
//
// Supported data types:
//   - integer (byte, uintptr, int, int8, int16, int32, int64, uint..uint64)
//   - floating point (float32, float64, complex64, complex128)
//   - string (string, rune)
//   - date/time
//   - boolean
func TypeToGQLType(typ string) *graphql.Scalar {
	switch typ {
	case "string", "rune":
		return graphql.String
	case
		"byte", "uintptr",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64":
		return graphql.Int
	case "float32", "float64", "complex64", "complex128":
		return graphql.Float
	case "bool":
		return graphql.Boolean
	case "date", "time":
		return graphql.DateTime
	default:
		return graphql.String
	}
}

// StructToFields takes a json tag and generate fields.
//
// The following function supports the json tag rules, and adding additional tags:
//
//  - type - The datatype to use (Golang's basic data types). If non provided, the
//    default data type of the field will be used.
//  - desc - description for the field
//  - deprecation - A deprecation message for the field
//
// Example for struct:
//
// type test struct {
//   ID       int       `json:"id,omitempty" desc:"foo bar" type:"int64"`
//   Name     string    `json:"-"`
//   DateTime time.Time `json:"date_time" type:"time" desc:"DateTime"`
//   KeyName  int64     `json:"-," deprecation:""`
//   Nothing  string
// }
func StructToFields(strct interface{}) graphql.Fields {
	fields := make(graphql.Fields)

	types := reflect.TypeOf(strct)

	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		json := field.Tag.Get("json")
		if json == "" || json == "-" {
			continue
		}
		jsonFields := strings.Split(json, ",")
		name := jsonFields[0]
		if name == "-" {
			name = gostrutils.CamelCaseToUnderscore(field.Name)
		}

		fields[name] = &graphql.Field{}
		typ := field.Tag.Get("type")
		if typ == "" {
			typ = field.Type.String()
		}
		fields[name].Type = TypeToGQLType(typ)
		fields[name].Description = field.Tag.Get("desc")
		fields[name].DeprecationReason = field.Tag.Get("deprecation")
	}

	return fields
}

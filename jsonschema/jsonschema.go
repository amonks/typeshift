// Package jsonschema contains go struct types representing json-schema
package jsonschema

type Type string

const (
	TypeObject  Type = "object"
	TypeInteger Type = "integer"
	TypeNumber  Type = "number"
	TypeBoolean Type = "boolean"
	TypeString  Type = "string"
	TypeArray   Type = "array"
	TypeNull    Type = "null"
)

type Schema interface{}

type SchemaObjectProperties = map[string]Schema

type StringFormat = string

const (
	StringFormatDateTime StringFormat = "date-time"
	StringFormatEmail    StringFormat = "email"
	StringFormatHostname StringFormat = "hostname"
	StringFormatIPv4     StringFormat = "ipv4"
	StringFormatIPv6     StringFormat = "ipv6"
	StringFormatURI      StringFormat = "uri"
)

type Any struct {
	Title       string         `json:"title"`
	Description string         `json:"description,omitempty"`
	Default     *interface{}   `json:"default,omitempty"`
	Enum        *[]interface{} `json:"enum,omitempty"`
}

type String struct {
	Title       string         `json:"title"`
	Description string         `json:"description,omitempty"`
	Type        Type           `json:"type"`
	Default     *interface{}   `json:"default,omitempty"`
	Enum        *[]interface{} `json:"enum,omitempty"`
	MinLength   *int           `json:"minLength,omitempty"`
	MaxLength   *int           `json:"maxLength,omitempty"`
	Pattern     *string        `json:"pattern,omitempty"`
	Format      *StringFormat  `json:"format,omitempty"`
}

type Number struct {
	Title            string         `json:"title"`
	Description      string         `json:"description,omitempty"`
	Type             Type           `json:"type"`
	Default          *interface{}   `json:"default,omitempty"`
	Enum             *[]interface{} `json:"enum,omitempty"`
	Minimum          *float32       `json:"minimum,omitempty"`
	Maximum          *float32       `json:"maximum,omitempty"`
	ExclusiveMaximum *bool          `json:"exclusiveMaximum,omitempty"`
}

// An Integer in json-schema.
type Integer struct {
	Title            string         `json:"title"`
	Description      string         `json:"description,omitempty"`
	Type             Type           `json:"type"`
	Default          *interface{}   `json:"default,omitempty"`
	Enum             *[]interface{} `json:"enum,omitempty"`
	Minimum          *int           `json:"minimum,omitempty"`
	Maximum          *int           `json:"maximum,omitempty"`
	ExclusiveMaximum *bool          `json:"exclusiveMaximum,omitempty"`
}

// An Object in json-schema.
type Object struct {
	Title         string                 `json:"title"`
	Description   string                 `json:"description,omitempty"`
	Type          Type                   `json:"type"`
	Default       *interface{}           `json:"default,omitempty"`
	Enum          *[]interface{}         `json:"enum,omitempty"`
	Properties    SchemaObjectProperties `json:"properties"`
	Required      *[]string              `json:"required,omitempty"`
	MinProperties *int                   `json:"minProperties,omitempty"`
	MaxProperties *int                   `json:"maxProperties,omitempty"`
}

// A Map in json-schema.
type Map struct {
	Title                string         `json:"title"`
	Description          string         `json:"description,omitempty"`
	Type                 Type           `json:"type"`
	Default              *interface{}   `json:"default,omitempty"`
	Enum                 *[]interface{} `json:"enum,omitempty"`
	AdditionalProperties Schema         `json:"additionalProperties"`
	Required             *[]string      `json:"required,omitempty"`
	MinProperties        *int           `json:"minProperties,omitempty"`
	MaxProperties        *int           `json:"maxProperties,omitempty"`
}

// An Array in json-schema.
type Array struct {
	Title       string         `json:"title"`
	Description string         `json:"description,omitempty"`
	Type        Type           `json:"type"`
	Default     *interface{}   `json:"default,omitempty"`
	Enum        *[]interface{} `json:"enum,omitempty"`
	Items       Schema         `json:"items"`
	MinItems    *int           `json:"minItems,omitempty"`
	MaxItems    *int           `json:"maxItems,omitempty"`
	UniqueItems bool           `json:"uniqueItems"`
}

// A Tuple in json-schema must have the "Array" type.
type Tuple struct {
	Title       string       `json:"title"`
	Description string       `json:"description,omitempty"`
	Default     *interface{} `json:"default,omitempty"`
	Items       []Schema     `json:"items"`
	Type        Type         `json:"type"`
	UniqueItems bool         `json:"uniqueItems"`
}

// Boolean is the boolean type.
type Boolean struct {
	Title       string       `json:"title"`
	Description string       `json:"description,omitempty"`
	Type        Type         `json:"type"`
	Default     *interface{} `json:"default,omitempty"`
}

// Null is the null type.
type Null struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Type        Type   `json:"type"`
}

// AnyOf represents a typescript union.
type AnyOf struct {
	AnyOf []Schema `json:"anyOf"`
}

// Nullable makes a schema nullable.
func Nullable(s Schema) Schema {
	return AnyOf{
		AnyOf: []Schema{
			s,
			Null{
				Type: TypeNull,
			},
		},
	}
}

// SchemaPtr gives you a pointer to a schema.
func SchemaPtr(s Schema) *Schema {
	return &s
}

package objects

// PublicObject is an object
type PublicObject struct {
	// DocStringField is a cool field
	DocStringField            string
	privateStringField        string
	NullableStringField       *string
	RenamedStringField        string `json:"new_string_field_name"`
	OmittedStringField        string `json:"-"`
	WeirdlyOmittedStringField string `json:"-,omitempty"`
	StructField               struct {
		Int    int
		String string
	}
}

// privateObject's docstring
type privateObject struct {
	// DocStringField is a cool field
	DocStringField            string
	privateStringField        string
	NullableStringField       *string
	RenamedStringField        string `json:"new_string_field_name"`
	OmittedStringField        string `json:"-"`
	WeirdlyOmittedStringField string `json:"-,omitempty"`
	StructField               struct {
		Int    int
		String string
	}
}

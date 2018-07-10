package readtypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func strPtr(s string) *string {
	return &s
}

func TestJsontag(t *testing.T) {
	testcases := []struct {
		input       string
		expectation jsonTag
	}{
		{``, jsonTag{}},
		{`json:""`, jsonTag{}},
		{`json:"new_string_field_name"`, jsonTag{
			tsName: strPtr("new_string_field_name"),
		}},
		{"`json:\"new_string_field_name\"`", jsonTag{
			tsName: strPtr("new_string_field_name"),
		}},
		{`json:"new_string_field_name,omitempty"`, jsonTag{
			tsName:    strPtr("new_string_field_name"),
			omitempty: true,
		}},
		{`json:"This is a very fancy fieldname!,omitempty"`, jsonTag{
			tsName:    strPtr("This is a very fancy fieldname!"),
			omitempty: true,
		}},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.expectation, ReadJsonTag(testcase.input))
	}
}

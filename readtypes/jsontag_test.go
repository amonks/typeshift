package readtypes_test

import (
	"testing"

	"github.com/amonks/typeshift/readtypes"
	"github.com/stretchr/testify/assert"
)

func strPtr(s string) *string {
	return &s
}

func TestJsontag(t *testing.T) {
	testcases := []struct {
		input       string
		expectation readtypes.JsonTag
	}{
		{``, readtypes.JsonTag{}},
		{`json:""`, readtypes.JsonTag{}},
		{`json:"new_string_field_name"`, readtypes.JsonTag{
			TSName: strPtr("new_string_field_name"),
		}},
		{"`json:\"new_string_field_name\"`", readtypes.JsonTag{
			TSName: strPtr("new_string_field_name"),
		}},
		{`json:"new_string_field_name,omitempty"`, readtypes.JsonTag{
			TSName:    strPtr("new_string_field_name"),
			Omitempty: true,
		}},
		{`json:"This is a very fancy fieldname!,omitempty"`, readtypes.JsonTag{
			TSName:    strPtr("This is a very fancy fieldname!"),
			Omitempty: true,
		}},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.expectation, readtypes.ReadJsonTag(testcase.input))
	}
}

package util_test

import (
	"testing"

	"github.com/amonks/typeshift/util"
	"github.com/stretchr/testify/assert"
)

func TestJsontag(t *testing.T) {
	testcases := []struct {
		input       string
		expectation util.JsonTag
	}{
		{``, util.JsonTag{}},
		{`json:""`, util.JsonTag{}},
		{`json:"new_string_field_name"`, util.JsonTag{
			TSName: util.StrPtr("new_string_field_name"),
		}},
		{"`json:\"new_string_field_name\"`", util.JsonTag{
			TSName: util.StrPtr("new_string_field_name"),
		}},
		{`json:"new_string_field_name,omitempty"`, util.JsonTag{
			TSName:    util.StrPtr("new_string_field_name"),
			Omitempty: true,
		}},
		{`json:"This is a very fancy fieldname!,omitempty"`, util.JsonTag{
			TSName:    util.StrPtr("This is a very fancy fieldname!"),
			Omitempty: true,
		}},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.expectation, util.ReadJsonTag(testcase.input))
	}
}

package readtypes

import (
	"reflect"
	"strings"
)

type jsonTag struct {
	skip      bool
	omitempty bool
	tsName    *string
}

func readJsonTag(t string) jsonTag {
	t = strings.Replace(t, "`", "", 2)
	var tag = reflect.StructTag(t)
	jsontag := tag.Get("json")
	if jsontag == "" {
		return jsonTag{}
	}

	omitempty := strings.Contains(jsontag, ",omitempty")
	jsontag = strings.Replace(jsontag, ",omitempty", "", 1)

	if jsontag == "-" {
		return jsonTag{skip: true}
	}

	var tsname *string
	if jsontag != "" {
		tsname = &jsontag
	}

	return jsonTag{
		omitempty: omitempty,
		tsName:    tsname,
	}
}

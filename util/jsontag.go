package util

import (
	"reflect"
	"strings"
)

type JsonTag struct {
	Skip      bool
	Omitempty bool
	TSName    *string
}

func ReadJsonTag(t string) JsonTag {
	var tag = reflect.StructTag(t)
	jsontag := tag.Get("json")
	if jsontag == "" {
		return JsonTag{}
	}

	omitempty := strings.Contains(jsontag, ",omitempty")
	jsontag = strings.Replace(jsontag, ",omitempty", "", 1)

	if jsontag == "-" {
		return JsonTag{Skip: true}
	}

	var tsname *string
	if jsontag != "" {
		tsname = &jsontag
	}

	return JsonTag{
		Omitempty: omitempty,
		TSName:    tsname,
	}
}

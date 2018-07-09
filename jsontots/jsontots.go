// Package jsontots contains a wrapper around json-schema-to-typescript
package jsontots

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/amonks/typeshift/jsonschema"
)

// JSONToTs converts jsonschema into typescript declarations, by feeding it into
// json-schema-to-typescript. If you pass it input that is _not_ json schema,
// it'll still marshal it and feed it to json-schema-to-typescript.
func JSONToTs(s jsonschema.Schema) string {
	command := exec.Command("json2ts")
	stdin, err := command.StdinPipe()
	if err != nil {
		panic(err)
	}

	json, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(stdin, string(json))
	stdin.Close()

	output, err := command.CombinedOutput()
	if err != nil {
		log.Println(output)
		panic(err)
	}

	return string(output)
}

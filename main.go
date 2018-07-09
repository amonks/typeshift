package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/amonks/typeshift/jsontots"
	"github.com/amonks/typeshift/readtypes"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
}

func usage() {
	fmt.Println("typeshift")
	fmt.Println("convert go types to json-schema")
	fmt.Println()
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	pkgpath := flag.String("path", ".", "path to the go package directory")
	format := flag.String("format", "json", "output format, 'ts' or 'json'")
	outputDir := flag.String("output", "stdout", "output dir or 'stdout'")
	flag.Parse()

	schemas, err := readtypes.ReadPackageDirectoryTypes(*pkgpath)
	if err != nil {
		panic(err)
	}

	output := func(name, s string) {
		fmt.Println(s)
	}
	if *outputDir != "stdout" {
		err := os.MkdirAll(*outputDir, 0755)
		if err != nil {
			panic(err)
		}
		output = func(name, s string) {
			filename := path.Join(*outputDir, name+"."+*format)
			err := ioutil.WriteFile(filename, []byte(s), 0644)
			if err != nil {
				panic(err)
			}
		}
	}

	for name, schema := range *schemas {
		switch *format {
		case "ts":
			output(name, jsontots.JSONToTs(schema))
		case "json":
			json, err := json.MarshalIndent(schema, "", "  ")
			if err != nil {
				panic(err)
			}
			output(name, string(json)+"\n\n\n\n")
		default:
			panic("unsupported format!")
		}
	}

}

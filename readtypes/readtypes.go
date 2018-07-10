// Package readtypes uses go's ast api to convert go type declarations into
// json-schema
package readtypes

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/amonks/typeshift/jsonschema"
)

func nonEmptyStringPtrOr(s *string, alt string) string {
	if s == nil || *s == "" {
		return alt
	}
	return *s
}

func docToDescription(docs ...*ast.CommentGroup) string {
	out := ""
	for _, doc := range docs {
		if doc == nil {
			continue
		}
		out += doc.Text()
	}
	return strings.Trim(out, "  \n")
}

func structExprToSchema(name, description string, structType ast.StructType) (*jsonschema.Schema, error) {
	fields := map[string]jsonschema.Schema{}
	requiredFields := []string{}
	for _, f := range structType.Fields.List {
		var name string
		for _, id := range f.Names {
			name = id.Name
		}

		var jsontag = jsonTag{}
		if f.Tag != nil {
			jsontag = readJsonTag(f.Tag.Value)
			name = nonEmptyStringPtrOr(jsontag.tsName, name)
			if jsontag.skip {
				continue
			}
		}
		description := docToDescription(f.Doc, f.Comment)
		t, err := ExprToSchema(name, description, f.Type)
		if err != nil {
			return nil, err
		}

		isOptional := false
		if !isOptional {
			requiredFields = append(requiredFields, name)
		}

		fields[name] = t
	}
	return jsonschema.SchemaPtr(jsonschema.Object{
		Title:       name,
		Description: description,
		Properties:  fields,
		Required:    &requiredFields,
	}), nil
}

func ifaceExprToSchema(name, description string, iface ast.InterfaceType) (*jsonschema.Schema, error) {
	return jsonschema.SchemaPtr(jsonschema.Any{Title: name, Description: description}), nil
}

func mapExprToSchema(name, description string, record ast.MapType) (*jsonschema.Schema, error) {
	log.Println("map", name, record)
	recordType, err := ExprToSchema(name, description, record.Value)
	if err != nil {
		return nil, err
	}
	return jsonschema.SchemaPtr(jsonschema.Map{
		Title:                name,
		Description:          description,
		Type:                 jsonschema.TypeObject,
		AdditionalProperties: recordType,
	}), nil
}

func arrayExprToSchema(name, description string, array ast.ArrayType) (*jsonschema.Schema, error) {
	arrayType, err := ExprToSchema("", "", array.Elt)
	if err != nil {
		return nil, err
	}
	return jsonschema.SchemaPtr(jsonschema.Array{
		Title:       name,
		Description: description,
		Type:        jsonschema.TypeArray,
		Items:       arrayType,
	}), nil
}

func identToSchema(name, description string, ident ast.Ident) (*jsonschema.Schema, error) {
	// handle scalars

	switch ident.Name {
	case "string":
		return jsonschema.SchemaPtr(jsonschema.String{
			Title:       name,
			Description: description,
			Type:        jsonschema.TypeString,
		}), nil
	case "float32", "float64":
		return jsonschema.SchemaPtr(jsonschema.Number{
			Title:       name,
			Description: description,
			Type:        jsonschema.TypeNumber,
		}), nil
	case "int64", "int":
		return jsonschema.SchemaPtr(jsonschema.Integer{
			Title:       name,
			Description: description,
			Type:        jsonschema.TypeInteger,
		}), nil
	case "bool":
		return jsonschema.SchemaPtr(jsonschema.Boolean{
			Title:       name,
			Description: description,
			Type:        jsonschema.TypeBoolean,
		}), nil

	default:
		// handle references

		decl := ident.Obj
		log.Println("kind", decl.Kind)
		log.Println(reflect.ValueOf(decl.Decl))
		log.Println(reflect.TypeOf(decl.Decl))
		switch spec := decl.Decl.(type) {
		case *ast.TypeSpec:
			log.Println("reference type", spec.Name)
			refType, err := ExprToSchema(name, description, spec.Type)
			if err != nil {
				return nil, err
			}
			return refType, nil
		}

		panic(fmt.Sprintf("unsupported scalar or bad reference %s (%v)", name, ident))
	}
}

func starExprToSchema(name, description string, star ast.StarExpr) (*jsonschema.Schema, error) {
	nullableType, err := ExprToSchema(name, description, star.X)
	if err != nil {
		return nil, err
	}
	if nullableType == nil {
		panic("nill nullable")
	}
	return jsonschema.SchemaPtr(jsonschema.Nullable(*nullableType)), nil
}

// ExprToSchema converts a named ast.Expr into a jsonschema.Schema declaration.
func ExprToSchema(name, description string, expr ast.Expr) (*jsonschema.Schema, error) {
	// we'll do a big type switch on expr

	if structType, isStruct := expr.(*ast.StructType); isStruct {
		return structExprToSchema(name, description, *structType)
	} else if iface, isIface := expr.(*ast.InterfaceType); isIface {
		return ifaceExprToSchema(name, description, *iface)
	} else if record, isMap := expr.(*ast.MapType); isMap {
		return mapExprToSchema(name, description, *record)
	} else if array, isArray := expr.(*ast.ArrayType); isArray {
		return arrayExprToSchema(name, description, *array)
	} else if ident, isIdent := expr.(*ast.Ident); isIdent {
		return identToSchema(name, description, *ident)
	} else if nullable, isNullable := expr.(*ast.StarExpr); isNullable {
		return starExprToSchema(name, description, *nullable)
	}

	log.Println("not found", name, expr)
	panic(fmt.Sprintf("unsupported type declaration %s", name))
}

type namedExpr struct {
	name        string
	description string
	expr        ast.Expr
}

type schemas = map[string]jsonschema.Schema

func (t namedExpr) schema() (jsonschema.Schema, error) {
	return ExprToSchema(t.name, t.description, t.expr)
}

type namedSchema struct {
	schema jsonschema.Schema
	name   string
}

func readPackageTypes(p *ast.Package) (*[]namedSchema, error) {
	types := []namedExpr{}
	for _, f := range p.Files {
		for _, d := range f.Decls {
			if gen, isGen := d.(*ast.GenDecl); isGen && gen.Tok == token.TYPE {
				description := docToDescription(gen.Doc)
				specs := gen.Specs
				var expr ast.Expr
				var name string
				for _, spec := range specs {
					if typespec, isTypespec := spec.(*ast.TypeSpec); isTypespec {
						name = typespec.Name.Name
						description += docToDescription(typespec.Doc, typespec.Comment)
						expr = typespec.Type
					}
				}
				types = append(types, namedExpr{name: name, description: description, expr: expr})
			}
		}
	}

	typedescriptions := []namedSchema{}
	for _, t := range types {
		schema, err := t.schema()
		if err != nil {
			return nil, err
		}
		typedescriptions = append(typedescriptions, namedSchema{schema: schema, name: t.name})
	}
	return &typedescriptions, nil
}

// ReadPackageDirectoryTypes produces a map from names to json-schema
// declarations for each type declaration in a go package directory.
func ReadPackageDirectoryTypes(path string) (*map[string]jsonschema.Schema, error) {
	set := token.NewFileSet()
	packs, err := parser.ParseDir(set, path, nil, parser.ParseComments)
	if err != nil {
		log.Println("Failed to parse package:", err)
		os.Exit(1)
	}

	types := map[string]jsonschema.Schema{}
	for _, pack := range packs {
		ts, err := readPackageTypes(pack)
		if err != nil {
			return nil, err
		}
		for _, t := range *ts {
			types[t.name] = t.schema
		}
	}
	return &types, err
}

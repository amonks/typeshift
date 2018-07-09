Typeshift is some code for converting go types to json-schema or to
typescript type declarations.

Converting to typescript type declarations depends on
json-schema-to-typescript.

## Install it

    npm install -g json-schema-to-typescript
    go get github.com/amonks/typeshift
    go install github.com/amonks/typeshift

## Use it

    cd ~/go/src/my/package
    typeshift -path . -format ts -output ts/go-types

## Dev it

These are the packages:

- **jsonschema** defines go structs for marshaling json-schema
- **jsontots** wraps the node module json-schema-to-typescript
- **readtypes** uses go's ast api to generate json-schema for the type declarations in a package
- **testpackages** are a bunch of packages with different types of declaration in them for testing
- **util** is what it sounds like
- **main.go** ties it all together

run `./snapshot-test` to make typescript declarations for testpackages in
testdata.

## Is it any good?

Meh, it's OK. It's pretty incomplete. And I still don't know very much about
go.

## Now I want to...

### ...validate that data at the boundaries of my TS system conform to these types

You'll want a "json schema validator". Bonus: make your validators _type
guards_.

### ...validate that data moving through my TS system conforms to these types

You'll want a "runtime contracts system". Bonus: only enable it in TEST.

### ...generate mock data, say, for a storybook, that conforms to these types

You'll want a "json schema mocker".

### ...test my functions by generating edgy-but-schema-conformant-parameters and then validating that the result conforms to schema

You'll want an implementation of "quickcheck" that can make "arb"s from json-schema.

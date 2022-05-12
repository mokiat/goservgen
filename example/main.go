package main

import (
	"github.com/mokiat/goservgen"
	"github.com/mokiat/goservgen/core/dsl"
	"github.com/mokiat/goservgen/grpc/grpcdsl"
)

func main() {
	goservgen.Run(
	// TODO: specify model serializer
	// TODO: specify http model serializer? (or just use anonymous?)
	// TODO: specify http client serializer
	// TODO: specify http server serializer
	// TODO: specify grpcserializer
	)
}

var Pkg = dsl.Package("../gen", "github.com/mokiat/goservgen/gen")

var Person = dsl.Type(Pkg, "Person", dsl.Struct, func() {
	dsl.Comment(`
		Person represents a human individual with all the
		relevant data for them.
	`)
	dsl.Field("Name", dsl.String, func() {
		dsl.Comment(`
			Name specifies the person's composite name, which could include
			for example their given name and family name.
		`)
		grpcdsl.FieldIndex(1)
	})
	dsl.Field("Age", dsl.Int, func() {
		dsl.Comment(`
			Age specifies the person's age.
		`)
		grpcdsl.FieldIndex(2)
	})
	dsl.Field("Location", dsl.Struct, func() {
		dsl.Comment(`
			Location is a composite type that represents a person's location
			of living.
		`)
		dsl.Field("Country", dsl.String, nil)
		dsl.Field("City", dsl.String, nil)
		dsl.Field("Address", dsl.String, nil)
	})
})

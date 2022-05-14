//go:build tools

package main

import (
	"github.com/mokiat/goservgen"
	"github.com/mokiat/goservgen/core/dsl"
	"github.com/mokiat/goservgen/grpc/grpcdsl"
)

// NOTE: the `tools` build constraint aims to reduce package overhead on
// external projects that depend on this one.
//
// In order to generate this file, you need to make sure to pass the `tools` tag
// to the `go generate` command, as follows:
//
//     go generate -tags tools ./...
//
// Otherwise this file will be skipped. Having this behavior can be a useful
// in some situations.

//go:generate go run -tags tools ./

func main() {
	goservgen.Run(
	// TODO: specify core model serializer
	// TODO: specify core model validator
	//       - It can be used to check for duplicate type names
	//       - Or it can warn if comments don't follow the Go naming convention
	// TODO: specify http model serializer? (or just use anonymous?)
	// TODO: specify http client serializer
	// TODO: specify http server serializer
	// TODO: specify grpcserializer
	)
}

var pkg = dsl.Package("../gen", "github.com/mokiat/goservgen/gen", func() {
	dsl.Comment(`
		Package gen holds the API resources for our project.
	`)
})

var _ = dsl.Service(pkg, "Users", func() {
	dsl.Comment(`
		Users represents a service that can manage users within
		our microservice.
	`)

	dsl.Method("Create", func() {
		dsl.Comment(`
			Create stores a new user.
		`)

	})

	dsl.Method("Get", func() {
		dsl.Comment(`
			Get retrieves a user by their ID.
		`)

	})

	dsl.Method("Delete", func() {
		dsl.Comment(`
			Delete removes the user with the specified ID.
		`)

	})

	dsl.Method("List", func() {
		dsl.Comment(`
			List returns a list of all users managed by this service.
		`)

	})
})

var typePerson = dsl.Type(pkg, "Person", dsl.Struct, func() {
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

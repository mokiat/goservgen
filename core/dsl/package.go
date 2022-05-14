package dsl

import "fmt"

func Package(dir, namespace string, fn func()) *PackageDef {
	fmt.Println("Package...")
	return &PackageDef{} // TODO
}

type PackageDef struct {
	// TODO
}

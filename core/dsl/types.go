package dsl

import "fmt"

func Type(pkg *PackageDef, name string, baseType *TypeDef, fn func()) *TypeDef {
	fmt.Println("Type...")
	return &TypeDef{} // TODO
}

type TypeDef struct {
	// TODO
}

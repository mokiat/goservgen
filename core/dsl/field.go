package dsl

import "github.com/mokiat/goservgen/core/exec"

func Field(name string, baseType *TypeDef, fn func()) {
	exec.Schedule(func(c *exec.Context) error {
		panic("TODO: Field processing logic")
	})
}

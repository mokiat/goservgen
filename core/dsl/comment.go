package dsl

import (
	"github.com/mokiat/goservgen/core/exec"
	"github.com/mokiat/goservgen/core/model"
)

func Comment(content string) {
	exec.Schedule(func(c *exec.Context) error {
		c.Owner.AttachMetadata(&model.Comment{
			Content: content,
		})
		return nil
	})
}

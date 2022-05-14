package dsl

import (
	"fmt"

	"github.com/mokiat/goservgen/core/exec"
)

// Comment attaches a comment to the owning entity.
func Comment(comment string) {
	exec.Schedule(func(c *exec.Context) error {
		commentable, ok := c.Owner.(interface {
			SetComment(string)
		})
		if !ok {
			return fmt.Errorf("cannot attach Comment to entity")
		}
		commentable.SetComment(comment)
		return nil
	})
}

package model

// CommentComponent is intended to be composed in entities that can accept
// a GoDoc comment.
type CommentComponent struct {
	comment string
}

// Comment returns the entity's comment.
func (h *CommentComponent) Comment() string {
	return h.comment
}

// SetComment sets the entity's comment to the specified value.
func (h *CommentComponent) SetComment(comment string) {
	h.comment = comment
}

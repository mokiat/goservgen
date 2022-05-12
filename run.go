package goservgen

// Serializer is a processor that can consume a root tree.Node and generate
// usable resources (e.g. HTTP clients, HTTP models, etc)
type Serializer interface {
	// TODO
}

// WithSerializer is a configuration Option that requests that the specified
// serializer be allowed to process the definitions tree.
func WithSerializer(serializer Serializer) Option {
	return func(c *config) {
		c.serializers = append(c.serializers, serializer)
	}
}

// Option represents a configuration option for the Run function.
type Option func(c *config)

type config struct {
	serializers []Serializer
}

// Run executes the goservgen processing.
func Run(opts ...Option) {
	// TODO
}

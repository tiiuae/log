package resource

import (
	"context"
)

// Resource is a source for a log entry
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/resource/semantic_conventions/README.md
type Resource map[string]interface{}

type key struct{}

var resourceKey key

// NewContext returns a context which contains the given resource.
func NewContext(ctx context.Context, resource Resource) context.Context {
	return context.WithValue(ctx, resourceKey, resource)
}

// FromContext returns a resource associated with this context
func FromContext(ctx context.Context) Resource {
	resource, _ := ctx.Value(resourceKey).(Resource)
	return resource
}

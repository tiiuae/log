package log

import (
	"time"

	"github.com/tiiuae/log/resource"
)

// TraceFlags describes the tracing status
type TraceFlags byte

const (
	// TraceFlagsNone describes empty set of tracing flags
	TraceFlagsNone TraceFlags = 0

	// TraceFlagsSampled is set when a log entry is part of sampled trace
	TraceFlagsSampled = 1 << iota
)

// Entry describes content and information of a single log line
//
// See OpenTelemetry Speficifcation at
// https://opentelemetry.io/docs/reference/specification/logs/data-model/
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/data-model.md
type Entry struct {
	Timestamp  time.Time
	Severity   Severity
	Body       interface{}
	TraceID    string
	SpanID     string
	TraceFlags TraceFlags
	Resource   resource.Resource
	Attributes Attributes
}

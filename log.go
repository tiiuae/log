package log

import (
	"context"
	"os"
	"time"

	"github.com/tiiuae/log/resource"
	"go.opentelemetry.io/otel/trace"
)

// DefaultSink is used for logging entries with package level logging functions
var DefaultSink Sink = &SimpleSink{}

// Print will log entries with Trace1 severity.
//
// Print log entries have no specific level of interest.
func Print(ctx context.Context, message string, attrs ...Attribute) {
	DefaultSink.LogEntry(ctx, createEntry(ctx, SeverityTrace1, message, attrs))
}

// Debug will log entries with Debug1 severity.
//
// Debug log entries are meant for tracing / debugging information.
func Debug(ctx context.Context, message string, attrs ...Attribute) {
	DefaultSink.LogEntry(ctx, createEntry(ctx, SeverityDebug1, message, attrs))
}

// Info will log entries with Info1 severity.
//
// Info log entries are meant for routine information, such as ongoing status or performance.
func Info(ctx context.Context, message string, attrs ...Attribute) {
	DefaultSink.LogEntry(ctx, createEntry(ctx, SeverityInfo1, message, attrs))
}

// Notice will log entries with Info2 severity.
//
// Notice log entries are meant for normal but significant events, such as start up, shut down or configuration.
func Notice(ctx context.Context, message string, attrs ...Attribute) {
	DefaultSink.LogEntry(ctx, createEntry(ctx, SeverityInfo2, message, attrs))
}

// Warning will log entries with Warn1 severity.
//
// Warning log entries describe events that might cause problems.
func Warning(ctx context.Context, message string, attrs ...Attribute) {
	DefaultSink.LogEntry(ctx, createEntry(ctx, SeverityWarn1, message, attrs))
}

// Error will log entries with Error1 severity.
//
// Error log entries are meant for events that are likely to cause problems.
func Error(ctx context.Context, message string, attrs ...Attribute) {
	DefaultSink.LogEntry(ctx, createEntry(ctx, SeverityError1, message, attrs))
}

// Critical will log entries with Error3 severity.
//
// Critical log entries are meant for events that cause more severe problems or brief outages.
func Critical(ctx context.Context, message string, attrs ...Attribute) {
	DefaultSink.LogEntry(ctx, createEntry(ctx, SeverityError3, message, attrs))
}

// Fatal will log entries with Fatal1 severity.
//
// Fatal log entries are Fatal to the application execution but not for the whole system.
//
// Fatal will log entry and exit the process with os.Exit(1).
// In case the logging the entry fails the function will panic.
func Fatal(ctx context.Context, message string, attrs ...Attribute) {
	DefaultSink.LogEntry(ctx, createEntry(ctx, SeverityFatal1, message, attrs))
	err := DefaultSink.Sync(ctx)
	if err != nil {
		panic(err)
	}
	os.Exit(1)
}

// Alert will log entries with Fatal2 severity.
//
// Alert log entries means a person must take an action immediately.
func Alert(ctx context.Context, message string, attrs ...Attribute) {
	DefaultSink.LogEntry(ctx, createEntry(ctx, SeverityFatal2, message, attrs))
}

// Emergency will log entries with Fatal4 severity.
//
// Emergency log events mean one or more systems are unusable.
func Emergency(ctx context.Context, message string, attrs ...Attribute) {
	DefaultSink.LogEntry(ctx, createEntry(ctx, SeverityFatal4, message, attrs))
}

// createEntry combines given parameters and adds tracing information.
func createEntry(ctx context.Context, severity Severity, message string, attrs []Attribute) Entry {
	spanCtx := trace.SpanContextFromContext(ctx)
	traceID := spanCtx.TraceID().String()
	spanID := spanCtx.SpanID().String()
	traceFlags := TraceFlagsNone
	if spanCtx.IsSampled() {
		traceFlags |= TraceFlagsSampled
	}

	return Entry{
		Timestamp:  time.Now().UTC(),
		Severity:   severity,
		Body:       message,
		TraceID:    traceID,
		SpanID:     spanID,
		TraceFlags: traceFlags,
		Resource:   resource.FromContext(ctx),
		Attributes: attrs,
	}
}

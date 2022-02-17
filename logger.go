package log

import (
	"context"
	"os"
)

// Logger can be used for logging entries into a given sink
type Logger struct {
	sink Sink
}

// NewLogger returns a new logger with given sink
func NewLogger(sink Sink) *Logger {
	return &Logger{
		sink: sink,
	}
}

// Print will log entries with Trace1 severity.
//
// Print log entries have no specific level of interest.
func (l *Logger) Print(ctx context.Context, message string, attrs ...Attribute) {
	l.sink.LogEntry(ctx, createEntry(ctx, SeverityTrace1, message, attrs))
}

// Debug will log entries with Debug1 severity.
//
// Debug log entries are meant for tracing / debugging information.
func (l *Logger) Debug(ctx context.Context, message string, attrs ...Attribute) {
	l.sink.LogEntry(ctx, createEntry(ctx, SeverityDebug1, message, attrs))
}

// Info will log entries with Info1 severity.
//
// Info log entries are meant for routine information, such as ongoing status or performance.
func (l *Logger) Info(ctx context.Context, message string, attrs ...Attribute) {
	l.sink.LogEntry(ctx, createEntry(ctx, SeverityInfo1, message, attrs))
}

// Notice will log entries with Info2 severity.
//
// Notice log entries are meant for normal but significant events, such as start up, shut down or configuration.
func (l *Logger) Notice(ctx context.Context, message string, attrs ...Attribute) {
	l.sink.LogEntry(ctx, createEntry(ctx, SeverityInfo2, message, attrs))
}

// Warning will log entries with Warn1 severity.
//
// Warning log entries describe events that might cause problems.
func (l *Logger) Warning(ctx context.Context, message string, attrs ...Attribute) {
	l.sink.LogEntry(ctx, createEntry(ctx, SeverityWarn1, message, attrs))
}

// Error will log entries with Error1 severity.
//
// Error log entries are meant for events that are likely to cause problems.
func (l *Logger) Error(ctx context.Context, message string, attrs ...Attribute) {
	l.sink.LogEntry(ctx, createEntry(ctx, SeverityError1, message, attrs))
}

// Critical will log entries with Error3 severity.
//
// Critical log entries are meant for events that cause more severe problems or brief outages.
func (l *Logger) Critical(ctx context.Context, message string, attrs ...Attribute) {
	l.sink.LogEntry(ctx, createEntry(ctx, SeverityError3, message, attrs))
}

// Fatal will log entries with Fatal1 severity.
//
// Fatal log entries are Fatal to the application execution but not for the whole system.
//
// Fatal will log entry and exit the process with os.Exit(1).
// In case the logging the entry fails the function will panic.
func (l *Logger) Fatal(ctx context.Context, message string, attrs ...Attribute) {
	l.sink.LogEntry(ctx, createEntry(ctx, SeverityFatal1, message, attrs))
	err := l.sink.Sync(ctx)
	if err != nil {
		panic(err)
	}
	os.Exit(1)
}

// Alert will log entries with Fatal2 severity.
//
// Alert log entries means a person must take an action immediately.
func (l *Logger) Alert(ctx context.Context, message string, attrs ...Attribute) {
	l.sink.LogEntry(ctx, createEntry(ctx, SeverityFatal2, message, attrs))
}

// Emergency will log entries with Fatal4 severity.
//
// Emergency log events mean one or more systems are unusable.
func (l *Logger) Emergency(ctx context.Context, message string, attrs ...Attribute) {
	l.sink.LogEntry(ctx, createEntry(ctx, SeverityFatal4, message, attrs))
}

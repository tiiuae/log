package log

import "context"

// Sink will log entries
type Sink interface {
	// LogEntry records log entries
	LogEntry(ctx context.Context, entry Entry)

	// Sync will make sure all entries recorded are persisted.
	Sync(ctx context.Context) error
}

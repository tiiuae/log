package log

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// SimpleLogger logs events to stdout
type SimpleLogger struct {
}

// LogEntry will do simple formatting and output given entry to stdout
func (s *SimpleLogger) LogEntry(ctx context.Context, entry Entry) {
	line := entry.Timestamp.Format(time.RFC3339)
	line += " "
	line += fmt.Sprintf("%s %v", entry.Severity, entry.Body)
	if entry.Attributes != nil {
		entries := make([]string, len(entry.Attributes))
		for i, a := range entry.Attributes {
			entries[i] = fmt.Sprintf("%s: %v", a.Name, a.Value)
		}
		line += fmt.Sprintf(" (%s)", strings.Join(entries, ", "))
	}
	fmt.Println(line)
}

// Sync does nothing with SimpleLogger as all entries are written immediately
func (s *SimpleLogger) Sync(ctx context.Context) error {
	return nil
}

package gcp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	"cloud.google.com/go/logging"
	"github.com/tiiuae/log"
)

// Sink is a log sink which will sink entries into Google Cloud Logging
type Sink struct {
	projectID string
	logger    *logging.Logger
}

// WriterSink is a log sink which will format log entries into Google Cloud Logging supported format and write into the writer
//
// https://cloud.google.com/logging/docs/structured-logging
type WriterSink struct {
	projectID string
	writer    io.Writer
	mu        sync.Mutex
}

var severityMap map[log.Severity]logging.Severity = map[log.Severity]logging.Severity{
	log.SeverityNone:   logging.Default,
	log.SeverityTrace1: logging.Default,
	log.SeverityTrace2: logging.Default,
	log.SeverityTrace3: logging.Default,
	log.SeverityTrace4: logging.Default,
	log.SeverityDebug1: logging.Debug,
	log.SeverityDebug2: logging.Debug,
	log.SeverityDebug3: logging.Debug,
	log.SeverityDebug4: logging.Debug,
	log.SeverityInfo1:  logging.Info,
	log.SeverityInfo2:  logging.Notice,
	log.SeverityInfo3:  logging.Notice,
	log.SeverityInfo4:  logging.Notice,
	log.SeverityWarn1:  logging.Warning,
	log.SeverityWarn2:  logging.Warning,
	log.SeverityWarn3:  logging.Warning,
	log.SeverityWarn4:  logging.Warning,
	log.SeverityError1: logging.Error,
	log.SeverityError2: logging.Error,
	log.SeverityError3: logging.Critical,
	log.SeverityError4: logging.Critical,
	log.SeverityFatal1: logging.Critical,
	log.SeverityFatal2: logging.Alert,
	log.SeverityFatal3: logging.Alert,
	log.SeverityFatal4: logging.Emergency,
}

// New creates a new sink with Google Cloud Logging as a backend
func New(ctx context.Context, projectID string, gcpLogger *logging.Logger) log.Sink {
	return &Sink{
		projectID: projectID,
		logger:    gcpLogger,
	}
}

// NewWriter creates a new sink which will output JSON in Google Cloud Logging supported format
func NewWriter(ctx context.Context, projectID string, w io.Writer) log.Sink {
	return &WriterSink{
		projectID: projectID,
		writer:    w,
	}
}

// LogEntry will record log entry into Google Cloud Logging
func (s *Sink) LogEntry(ctx context.Context, e log.Entry) {
	labels := make(map[string]string)
	for _, a := range e.Attributes {
		labels[a.Name] = fmt.Sprintf("%v", a.Value)
	}
	gcpEntry := logging.Entry{
		Timestamp: e.Timestamp,
		Severity:  severityMap[e.Severity],
		Payload:   e.Body,
		Labels:    labels,
		//InsertID: string
		//HTTPRequest: *HTTPRequest
		//Operation: *logpb.LogEntryOperation
		//Resource: *mrpb.MonitoredResource
	}

	if e.TraceFlags&log.TraceFlagsSampled != 0 {
		gcpEntry.Trace = fmt.Sprintf("projects/%s/traces/%s", s.projectID, e.TraceID)
		gcpEntry.SpanID = e.SpanID
		gcpEntry.TraceSampled = e.TraceFlags&log.TraceFlagsSampled != 0
	}

	s.logger.Log(gcpEntry)
}

// Sync will flush any pending log entries into Google Cloud Logging
func (s *Sink) Sync(ctx context.Context) error {
	return s.logger.Flush()
}

// LogEntry will record log entry in JSON format with given writer
func (s *WriterSink) LogEntry(ctx context.Context, e log.Entry) {
	labels := make(map[string]string)
	for _, a := range e.Attributes {
		labels[a.Name] = fmt.Sprintf("%v", a.Value)
	}

	entry := map[string]interface{}{
		"time":                          e.Timestamp.Format(time.RFC3339Nano),
		"severity":                      severityMap[e.Severity].String(),
		"message":                       fmt.Sprintf("%s", e.Body),
		"logging.googleapis.com/labels": labels,
	}
	if e.TraceFlags&log.TraceFlagsSampled != 0 {
		entry["logging.googleapis.com/spanId"] = e.SpanID
		entry["logging.googleapis.com/trace"] = fmt.Sprintf("projects/%s/traces/%s", s.projectID, e.TraceID)
		entry["logging.googleapis.com/trace_sampled"] = e.TraceFlags&log.TraceFlagsSampled != 0
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	err := json.NewEncoder(s.writer).Encode(entry)
	if err != nil {
		panic(err)
	}
}

// Sync with writer won't do anything as all entries are already written
func (s *WriterSink) Sync(ctx context.Context) error {
	return nil
}

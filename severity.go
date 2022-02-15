package log

import "strconv"

// Severity is the severity of the event described in a log entry. These
// guideline severity levels are ordered, with numerically smaller levels
// treated as less severe than numerically larger levels.
type Severity int

const (
	SeverityNone Severity = iota
	SeverityTrace1
	SeverityTrace2
	SeverityTrace3
	SeverityTrace4
	SeverityDebug1
	SeverityDebug2
	SeverityDebug3
	SeverityDebug4
	SeverityInfo1
	SeverityInfo2
	SeverityInfo3
	SeverityInfo4
	SeverityWarn1
	SeverityWarn2
	SeverityWarn3
	SeverityWarn4
	SeverityError1
	SeverityError2
	SeverityError3
	SeverityError4
	SeverityFatal1
	SeverityFatal2
	SeverityFatal3
	SeverityFatal4
)

var severityName = map[Severity]string{
	SeverityNone:   "NONE",
	SeverityTrace1: "TRACE1",
	SeverityTrace2: "TRACE2",
	SeverityTrace3: "TRACE3",
	SeverityTrace4: "TRACE4",
	SeverityDebug1: "DEBUG1",
	SeverityDebug2: "DEBUG2",
	SeverityDebug3: "DEBUG3",
	SeverityDebug4: "DEBUG4",
	SeverityInfo1:  "INFO1",
	SeverityInfo2:  "INFO2",
	SeverityInfo3:  "INFO3",
	SeverityInfo4:  "INFO4",
	SeverityWarn1:  "WARN1",
	SeverityWarn2:  "WARN2",
	SeverityWarn3:  "WARN3",
	SeverityWarn4:  "WARN4",
	SeverityError1: "ERROR1",
	SeverityError2: "ERROR2",
	SeverityError3: "ERROR3",
	SeverityError4: "ERROR4",
	SeverityFatal1: "FATAL1",
	SeverityFatal2: "FATAL2",
	SeverityFatal3: "FATAL3",
	SeverityFatal4: "FATAL4",
}

// String converts a severity level to a string.
func (v Severity) String() string {
	// same as proto.EnumName
	s, ok := severityName[v]
	if ok {
		return s
	}
	return strconv.Itoa(int(v))
}

package log

// Attribute contains a name and value
type Attribute struct {
	Name  string
	Value interface{}
}

// Attributes describe a set of fields associated to a log entry
//
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/semantic_conventions/README.md
type Attributes []Attribute

// A is a shorthand of creating an Attribute
func A(name string, value interface{}) Attribute {
	return Attribute{
		Name:  name,
		Value: value,
	}
}

// E is a shorthand of creating an Error attribute
func E(err error) Attribute {
	return Attribute{
		Name:  "error",
		Value: err,
	}
}

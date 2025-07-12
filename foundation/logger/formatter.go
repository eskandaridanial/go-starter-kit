package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
	"text/template"
	"time"
)

// variable 'bufPool' is a pool of bytes buffers to reduce allocations
var bufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

// type 'Formatter' represents a logging formatter,
// concrete implementations are 'JSONFormatter' and 'TextFormatter'
type Formatter interface {
	Format(r Record) []byte
}

// struct 'JSONFormatter' implements 'Formatter' interface
type JSONFormatter struct{}

// function 'NewJSONFormatter' creates a new 'JSONFormatter'
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

// function 'Format' formats the given record as JSON
func (f *JSONFormatter) Format(r Record) []byte {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	payload := make(map[string]any, len(r.Fields)+4)
	payload["level"] = r.Level.String()
	payload["message"] = r.Message
	payload["caller"] = r.Caller
	payload["timestamp"] = r.Timestamp.Format(time.RFC3339Nano)

	for _, field := range r.Fields {
		payload[field.Key] = field.Value
	}

	b, _ := json.Marshal(payload)
	buf.Write(b)
	buf.WriteByte('\n')
	return buf.Bytes()
}

// const 'DefaultPattern' is the default pattern for text formatter
const DefaultPattern = "{{.timestamp}} [{{.level}}] {{.caller}} {{.message}} - [{{.fields}}]"

// struct 'TextFormatter' implements 'Formatter' interface
type TextFormatter struct {
	tmpl *template.Template
}

// function 'NewTextFormatter' creates a new 'TextFormatter'
func NewTextFormatter(pattern string) *TextFormatter {
	if pattern == "" {
		pattern = DefaultPattern
	}
	tmpl := template.Must(template.New("log").Parse(pattern))
	return &TextFormatter{tmpl: tmpl}
}

// function 'Format' formats the given record as text
func (f *TextFormatter) Format(r Record) []byte {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	var sb bytes.Buffer
	for i, field := range r.Fields {
		sb.WriteString(fmt.Sprintf("%s=%v", field.Key, field.Value))
		if i < len(r.Fields)-1 {
			sb.WriteByte(' ')
		}
	}

	data := map[string]any{
		"level":     r.Level.String(),
		"message":   r.Message,
		"caller":    r.Caller,
		"fields":    sb.String(),
		"timestamp": r.Timestamp.Format(time.RFC3339Nano),
	}

	err := f.tmpl.Execute(buf, data)
	if err != nil {
		buf.WriteString(fmt.Sprintf("logger: template error: %v", err))
	}

	buf.WriteByte('\n')
	return buf.Bytes()
}

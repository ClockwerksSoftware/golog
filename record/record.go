package record

import (
	"runtime"
	"time"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

const (
	// normal operation means this is called through a log function
	// and thereby is a couple layers removed from the call that
	// needs to be trapped
	logRecordCallerDepth = 2
)

type logRecord struct {
	recordByteCache
	recordStringCache

	location    LogRecordLocation
	name        string
	level       interfaces.Level
	message     string
	messageArgs []any
	attrs       []interfaces.Attribute
	datetime    time.Time

}

func NewLogRecord(name string, level LogLevelInt, message string, args ...any) interfaces.Record {
	lr := &logRecord{
		name: name,
		level: &LogLevel{
			level: level,
		},
		message:     message,
		messageArgs: args,
		attrs:       make([]interfaces.Attribute, 0),
		datetime:    time.Now().UTC(),
	}
	_, lr.location.filename, lr.location.line, lr.location.ok = runtime.Caller(logRecordCallerDepth)
	// retrieve the stack associated with the log record
	// 2048 byte buffer; same as some functionality in Go itself
	var stackData [2 << 10]byte
	stackLength := runtime.Stack(stackData[:], false)
	lr.location.stack = string(stackData[:stackLength])
	return lr
}

func (lr *logRecord) Location() interfaces.RecordLocation {
	return &lr.location
}

func (lr *logRecord) Name() string {
	return lr.name
}

func (lr *logRecord) Level() interfaces.Level {
	return lr.level
}

func (lr *logRecord) RawMessage() string {
	return lr.message
}

func (lr *logRecord) RawMessageArgs() []any {
	return lr.messageArgs
}

func (lr *logRecord) Time() time.Time {
	return lr.datetime
}

func (lr *logRecord) AddAttributes(attr ...interfaces.Attribute) {
	lr.attrs = append(lr.attrs, attr...)
}

func (lr *logRecord) Attributes() []interfaces.Attribute {
	return lr.attrs
}

var _ interfaces.Record = &logRecord{}

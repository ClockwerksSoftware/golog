package record

import (
	"runtime"
	"time"

	"github.com/barkimedes/go-deepcopy"

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
		messageArgs: make([]any, 0),
		attrs:       make([]interfaces.Attribute, 0),
		datetime:    time.Now().UTC(),
	}
	if len(args) > 0 {
		vCopy := deepcopy.MustAnything(args).([]any)
		lr.messageArgs = append(lr.messageArgs, vCopy...)
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
	// it would have been great if `deepcopy.MustAnything` would work
	// on the `Attribute` type; however, when calling it `attr[n]` it
	// returns an uninitialized copy. This unfortunately means that
	// a little time needs to be wasted to properly copy the `Attributes`
	attr_copy := make([]interfaces.Attribute, 0)
	for _, v := range attr {
		attr_copy = append(
			attr_copy,
			NewLogAttr(
				v.Key(),
				deepcopy.MustAnything(v.Value()),
			),
		)
	}
	lr.attrs = append(lr.attrs, attr_copy...)
}

func (lr *logRecord) Attributes() []interfaces.Attribute {
	return lr.attrs
}

var _ interfaces.Record = &logRecord{}

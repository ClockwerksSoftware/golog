package record

import (
	"fmt"
	"runtime"
	"time"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

const (
	logRecordCallerDepth = 2
)

type logRecord struct {
	location    LogRecordLocation
	name        string
	level       interfaces.Level
	message     string
	messageArgs []any
	attrs       []interfaces.Attribute
	datetime    time.Time

	cacheBytes  map[string][]byte
	cacheString map[string]string
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

func (lr *logRecord) CacheFormat(formatName string, formattedRecord []byte) {
	lr.cacheBytes[formatName] = formattedRecord
}
func (lr *logRecord) GetCacheFormat(formatName string) ([]byte, error) {
	if v, ok := lr.cacheBytes[formatName]; ok {
		return v, nil
	}

	return nil, fmt.Errorf(
		"%w: no cached record with name %q",
		interfaces.ErrNoCachedRecord,
		formatName,
	)
}

func (lr *logRecord) CacheFormatString(formatName string, formattedRecord string) {
	lr.cacheString[formatName] = formattedRecord
}
func (lr *logRecord) GetCacheFormatString(formatName string) (string, error) {
	if v, ok := lr.cacheString[formatName]; ok {
		return v, nil
	}

	return "", fmt.Errorf(
		"%w: no cached record with name %q",
		interfaces.ErrNoCachedRecord,
		formatName,
	)
}

var _ interfaces.Record = &logRecord{}

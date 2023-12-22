package loggers

import (
	"github.com/ClockwerksSoftware/golog/interfaces"
	"github.com/ClockwerksSoftware/golog/record"
)

// NOTE: Child loggers send back up to the root logger
// they're just interface points
type coreLogger struct {
	name         string
	handlers     []interfaces.Handler
	childLoggers map[string]interfaces.Log
}

func (lc *coreLogger) Name() string {
	return lc.name
}
func (lc *coreLogger) Info(msg string, args ...any) {
	r := record.NewLogRecord(rootLoggerName, record.INFO, msg, args...)
	lc.sendRecord(r)
}
func (lc *coreLogger) Warning(msg string, args ...any) {
	r := record.NewLogRecord(rootLoggerName, record.WARNING, msg, args...)
	lc.sendRecord(r)
}
func (lc *coreLogger) Error(msg string, args ...any) {
	r := record.NewLogRecord(rootLoggerName, record.ERROR, msg, args...)
	lc.sendRecord(r)
}
func (lc *coreLogger) Debug(msg string, args ...any) {
	r := record.NewLogRecord(rootLoggerName, record.DEBUG, msg, args...)
	lc.sendRecord(r)
}
func (lc *coreLogger) Critical(msg string, args ...any) {
	r := record.NewLogRecord(rootLoggerName, record.CRITICAL, msg, args...)
	lc.sendRecord(r)
}
func (lc *coreLogger) Log(l interfaces.Level, msg string, args ...any) {
	r := record.NewLogRecord(rootLoggerName, record.LogLevelInt(l.Int()), msg, args...)
	lc.sendRecord(r)
}

func (lc *coreLogger) ProcessRecord(r interfaces.Record) {
	lc.sendRecord(r)
}

func (lc *coreLogger) sendRecord(r interfaces.Record) {
	for _, handler := range lc.handlers {
		handler.Handle(r)
	}
}

func (lc *coreLogger) AddHandler(h interfaces.Handler) {
	lc.handlers = append(lc.handlers, h)
}

var _ interfaces.Log = &coreLogger{}

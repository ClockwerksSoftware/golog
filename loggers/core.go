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
func (lc *coreLogger) Info(msg string) {
	r := record.NewLogRecord(rootLoggerName, record.INFO, msg)
	lc.sendRecord(r)
}
func (lc *coreLogger) Infof(format string, args ...any) {
	r := record.NewLogRecord(rootLoggerName, record.INFO, format, args...)
	lc.sendRecord(r)
}
func (lc *coreLogger) Warning(msg string) {
	r := record.NewLogRecord(rootLoggerName, record.WARNING, msg)
	lc.sendRecord(r)
}
func (lc *coreLogger) Warningf(format string, args ...any) {
	r := record.NewLogRecord(rootLoggerName, record.WARNING, format, args...)
	lc.sendRecord(r)
}
func (lc *coreLogger) Error(msg string) {
	r := record.NewLogRecord(rootLoggerName, record.ERROR, msg)
	lc.sendRecord(r)
}
func (lc *coreLogger) Errorf(format string, args ...any) {
	r := record.NewLogRecord(rootLoggerName, record.ERROR, format, args...)
	lc.sendRecord(r)
}
func (lc *coreLogger) Debug(msg string) {
	r := record.NewLogRecord(rootLoggerName, record.DEBUG, msg)
	lc.sendRecord(r)
}
func (lc *coreLogger) Debugf(format string, args ...any) {
	r := record.NewLogRecord(rootLoggerName, record.DEBUG, format, args...)
	lc.sendRecord(r)
}
func (lc *coreLogger) Critical(msg string) {
	r := record.NewLogRecord(rootLoggerName, record.CRITICAL, msg)
	lc.sendRecord(r)
}
func (lc *coreLogger) Criticalf(format string, args ...any) {
	r := record.NewLogRecord(rootLoggerName, record.CRITICAL, format, args...)
	lc.sendRecord(r)
}
func (lc *coreLogger) Log(l interfaces.Level, msg string) {
	r := record.NewLogRecord(rootLoggerName, record.LogLevelInt(l.Int()), msg)
	lc.sendRecord(r)
}
func (lc *coreLogger) Logf(l interfaces.Level, format string, args ...any) {
	r := record.NewLogRecord(rootLoggerName, record.LogLevelInt(l.Int()), format, args...)
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

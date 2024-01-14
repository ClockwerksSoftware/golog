package loggers

import (
	"github.com/ClockwerksSoftware/golog/interfaces"
	"github.com/ClockwerksSoftware/golog/record"
)

// childLogger is a lightweight log interface for embedding in
// various Golang Packages to allow for them to be easily filtered out
// by the application. The name should be set to the package level.
// For instance, this package would be "github.com/ClockwerksSoftware/golog/loggers".
type childLogger struct {
	name string
	root interfaces.Log
}

func NewChildLogger(name string, root interfaces.Log) interfaces.Log {
	return &childLogger{
		name: name,
		root: root,
	}
}

func (cl *childLogger) Name() string {
	return cl.name
}
func (cl *childLogger) Info(msg string) {
	r := record.NewLogRecord(cl.name, record.INFO, msg)
	cl.sendRecord(r)
}
func (cl *childLogger) Infof(format string, args ...any) {
	r := record.NewLogRecord(cl.name, record.INFO, format, args...)
	cl.sendRecord(r)
}
func (cl *childLogger) Warning(msg string) {
	r := record.NewLogRecord(cl.name, record.WARNING, msg)
	cl.sendRecord(r)
}
func (cl *childLogger) Warningf(format string, args ...any) {
	r := record.NewLogRecord(cl.name, record.WARNING, format, args...)
	cl.sendRecord(r)
}
func (cl *childLogger) Error(msg string) {
	r := record.NewLogRecord(cl.name, record.ERROR, msg)
	cl.sendRecord(r)
}
func (cl *childLogger) Errorf(format string, args ...any) {
	r := record.NewLogRecord(cl.name, record.ERROR, format, args...)
	cl.sendRecord(r)
}
func (cl *childLogger) Debug(msg string) {
	r := record.NewLogRecord(cl.name, record.DEBUG, msg)
	cl.sendRecord(r)
}
func (cl *childLogger) Debugf(format string, args ...any) {
	r := record.NewLogRecord(cl.name, record.DEBUG, format, args...)
	cl.sendRecord(r)
}
func (cl *childLogger) Critical(msg string) {
	r := record.NewLogRecord(cl.name, record.CRITICAL, msg)
	cl.sendRecord(r)
}
func (cl *childLogger) Criticalf(format string, args ...any) {
	r := record.NewLogRecord(cl.name, record.CRITICAL, format, args...)
	cl.sendRecord(r)
}
func (cl *childLogger) Log(l interfaces.Level, msg string) {
	r := record.NewLogRecord(cl.name, record.LogLevelInt(l.Int()), msg)
	cl.sendRecord(r)
}
func (cl *childLogger) Logf(l interfaces.Level, format string, args ...any) {
	r := record.NewLogRecord(cl.name, record.LogLevelInt(l.Int()), format, args...)
	cl.sendRecord(r)
}

func (cl *childLogger) ProcessRecord(r interfaces.Record) {
	cl.sendRecord(r)
}

func (cl *childLogger) sendRecord(r interfaces.Record) {
	cl.root.ProcessRecord(r)
}

func (cl *childLogger) AddHandler(h interfaces.Handler) {
	// add it to the root handler
	cl.root.AddHandler(h)
}

var _ interfaces.Log = &childLogger{}

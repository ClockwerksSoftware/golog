package interfaces

type Log interface {
	Name() string
	Info(msg string)
	Infof(format string, v ...any)
	Warning(msg string)
	Warningf(format string, v ...any)
	Debug(msg string)
	Debugf(format string, v ...any)
	Error(msg string)
	Errorf(format string, v ...any)
	Critical(msg string)
	Criticalf(format string, v ...any)
	Log(l Level, msg string)
	Logf(l Level, format string, v ...any)

	ProcessRecord(r Record)
	AddHandler(h Handler)
}

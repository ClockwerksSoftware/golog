package interfaces

type Log interface {
   Name() string
   Info(format string, v ...any)
   Warning(format string, v ...any)
   Debug(format string, v ...any)
   Error(format string, v ...any)
   Critical(format string, v ...any)
   Log(l Level, format string, v ...any)

   ProcessRecord(r Record)
   AddHandler(h Handler)
}

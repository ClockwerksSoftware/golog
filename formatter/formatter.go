package formatter

import (
	"fmt"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

/*
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
*/

type logFormatter struct {
	name string
}

func NewLogFormatter() interfaces.Formatter {
	return &logFormatter{
		name: "LogFormatterBasic",
	}
}

func (lf *logFormatter) Name() string {
	return lf.name
}

func (lf *logFormatter) Format(r interfaces.Record) []byte {
	// basic version - just get the string and format it
	// NOTE: Not going to cache this as it would just create duplicate data
	// Other formatters might want to do the opposite method if their primary is the []byte usage
	return []byte(lf.FormatString(r))
}

func (lf *logFormatter) FormatString(r interfaces.Record) string {
	f, err := r.GetCacheFormatString(lf.name)
	if err == nil {
		return f
	}

	formattedMessage := fmt.Sprintf(
		r.RawMessage(),
		r.RawMessageArgs(),
	)
	// This is a super simplistic version of the formatting. it would be great to be able
	// to add some additional format strings to specify where the various pieces of the Record
	datetime, _ := r.Time().MarshalText()
	f = fmt.Sprintf("%s %s %s", r.Level().String(), datetime, formattedMessage)
	r.CacheFormatString(lf.name, f)
	return f
}

var _ interfaces.Formatter = &logFormatter{}

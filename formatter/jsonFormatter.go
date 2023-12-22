package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

type logJsonFormatter struct {
	name string
}

func NewLogJsonFormatter() interfaces.Formatter {
	return &logJsonFormatter{
		name: "LogJsonFormatter",
	}
}

func (ljf *logJsonFormatter) Name() string {
	return ljf.name
}
func (ljf *logJsonFormatter) Format(r interfaces.Record) []byte {
	// basic version - just get the string and format it
	// NOTE: Not going to cache this as it would just create duplicate data
	// Other formatters might want to do the opposite method if their primary is the []byte usage
	return []byte(ljf.FormatString(r))
}

func (ljf *logJsonFormatter) FormatString(r interfaces.Record) string {
	f, err := r.GetCacheFormatString(ljf.name)
	if err == nil {
		return f
	}

	jsonAttributes := make(map[string]any)
	jsonAttributes["name"] = r.Name()
	jsonAttributes["message"] = fmt.Sprintf(
		r.RawMessage(),
		r.RawMessageArgs(),
	)
	jsonAttributes["level"] = r.Level()
	jsonAttributes["time"], _ = r.Time().MarshalText()

	for _, attr := range r.Attributes() {
		jsonAttributes[attr.Key()] = attr.Value()
	}

	jsonString := "{"

	doEscape := func(s string) string {
		b := bytes.NewBufferString(s)
		b2 := []byte{}
		json.HTMLEscape(b, b2)
		return string(b2)
	}

	for k, v := range jsonAttributes {
		escapedKey := doEscape(k)
		escapedValue := doEscape(fmt.Sprintf("%v", v))

		jsonString = fmt.Sprintf("%s \"%s\": \"%s\",", jsonString, escapedKey, escapedValue)
	}
	// cut off the last ","
	jsonString = jsonString[:len(jsonString)-1]
	jsonString = fmt.Sprintf("%s }", jsonString)

	r.CacheFormatString(ljf.name, jsonString)
	return jsonString
}

var _ interfaces.Formatter = &logJsonFormatter{}

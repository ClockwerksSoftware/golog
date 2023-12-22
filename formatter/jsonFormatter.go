package formatter

import (
	//"bytes"
	"encoding/json"
	"fmt"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

const (
	jsonFormatterName = "LogJsonFormatter"
)

var (
	reservedJsonKeys = map[string]bool{
		"name": true,
		"message": true,
		"level": true,
		"time": true,
	}
)

type logJsonFormatter struct {
	name string
}

func NewLogJsonFormatter() interfaces.Formatter {
	return &logJsonFormatter{
		name: jsonFormatterName,
	}
}

func (ljf *logJsonFormatter) Name() string {
	return ljf.name
}
func (ljf *logJsonFormatter) Format(r interfaces.Record) []byte {
	f, err := r.GetCacheFormat(ljf.name)
	if err == nil {
		return f
	}

	jsonAttributes := make(map[string]any)
	jsonAttributes["name"] = r.Name()
	jsonAttributes["message"] = fmt.Sprintf(
		r.RawMessage(),
		r.RawMessageArgs(),
	)
	jsonAttributes["level"] = r.Level().String()
	tempTime, _ := r.Time().MarshalText()
	jsonAttributes["time"] = string(tempTime)

	for _, attr := range r.Attributes() {
		if _, ok := reservedJsonKeys[attr.Key()]; !ok {
			jsonAttributes[attr.Key()] = attr.Value()
		}
	}

	// TODO: replace the JSON parser
	jsonByteData, err := json.Marshal(jsonAttributes)
	r.CacheFormat(ljf.name, jsonByteData)
	return jsonByteData
}

func (ljf *logJsonFormatter) FormatString(r interfaces.Record) string {
	jsonByteData := ljf.Format(r)
	return string(jsonByteData)
}

var _ interfaces.Formatter = &logJsonFormatter{}

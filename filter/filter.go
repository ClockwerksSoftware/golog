package filter

import (
	"github.com/ClockwerksSoftware/golog/interfaces"
	"github.com/ClockwerksSoftware/golog/record"
)

const (
	nameLevelFilter = "levelFilter"
)

type logFilter struct {
	name  string
	level interfaces.Level
}

func NewLogFilter() interfaces.Filter {
	return &logFilter{
		name:  nameLevelFilter,
		level: record.GetLogLevel(record.DEBUG),
	}
}

func (lf *logFilter) Name() string {
	return lf.name
}

func (lf *logFilter) SetLevel(l interfaces.Level) {
	lf.level = l
}

func (lf *logFilter) Filter(r interfaces.Record) (output interfaces.Record, valid bool) {
	if r.Level().Int() >= lf.level.Int() {
		output = r
		valid = true
	}
	return
}

var _ interfaces.Filter = &logFilter{}

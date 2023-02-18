package record

import (
    "github.com/ClockwerksSoftware/golog/interfaces"
)

type LogRecordLocation struct {
    filename string
    line int
    ok bool
    stack string
}

func NewLogRecordLocation() interfaces.RecordLocation {
    return &LogRecordLocation{}
}

func (lrl *LogRecordLocation) SetFilename(f string) {
    lrl.filename = f
}

func (lrl *LogRecordLocation) Filename() string {
    return lrl.filename
}

func (lrl *LogRecordLocation) SetLine(n int) {
    lrl.line = n
}
func (lrl *LogRecordLocation) Line() int {
    return lrl.line
}

func (lrl *LogRecordLocation) SetValid(v bool) {
    lrl.ok = v
}
func (lrl *LogRecordLocation) Valid() bool {
    return lrl.ok
}

func (lrl *LogRecordLocation) SetStack(s string) {
    lrl.stack = s
}
func (lrl *LogRecordLocation) Stack() string {
    return lrl.stack
}

var _ interfaces.RecordLocation = &LogRecordLocation{}

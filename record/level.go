package record

import (
	"fmt"

	"github.com/ClockwerksSoftware/golog/common"
	"github.com/ClockwerksSoftware/golog/interfaces"
)

type LogLevelInt int
type LogLevel struct {
	level LogLevelInt
}

const (
	CRITICAL LogLevelInt = 0
	DEBUG                = 30
	ERROR                = 50
	WARNING              = 70
	INFO                 = 90
)

var (
	levelMap = map[LogLevelInt]string{
		CRITICAL: "CRITICAL",
		DEBUG:    "DEBUG",
		ERROR:    "ERROR",
		WARNING:  "WARNING",
		INFO:     "INFO",
	}
)

func NewLogLevel() interfaces.Level {
	return &LogLevel{}
}

func GetLogLevel(lvl LogLevelInt) interfaces.Level {
	return &LogLevel{
		level: lvl,
	}
}

func (l *LogLevel) Int() (n int) {
	n = int(l.level)
	return
}

func (l *LogLevel) String() (s string) {
	if v, ok := levelMap[l.level]; ok {
		s = v
		return
	}

	s = "UNKNOWN"
	return
}

func AddLogLevel(level LogLevelInt, name string, overwrite bool) (err error) {
	if v, ok := levelMap[level]; ok && !overwrite {
		err = fmt.Errorf("%w: Unable to map level %d as it is already mapped to %s", common.ErrLogLevelAlreadyExists, level, v)
		return
	}

	levelMap[level] = name
	return
}

var _ interfaces.Level = &LogLevel{}

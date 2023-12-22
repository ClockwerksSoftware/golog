package record

import (
	"github.com/ClockwerksSoftware/golog/interfaces"
)

// TODO: Use the `LogAttr` from slog as it handles the Value a bit better
// especially regarding details around what doesn't work with any/interface{}
type LogAttr struct {
	key   string
	value any
}

func (la *LogAttr) Key() string {
	return la.key
}

func (la *LogAttr) Value() any {
	return la.value
}

var _ interfaces.Attribute = &LogAttr{}

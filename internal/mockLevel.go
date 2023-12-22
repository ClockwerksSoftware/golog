package internal

import (
	"github.com/ClockwerksSoftware/golog/interfaces"
)

type MockLevel struct {
	StrValue string
	IntValue int
}

func (ml *MockLevel) String() string {
	return ml.StrValue
}

func (ml *MockLevel) Int() int {
	return ml.IntValue
}

var _ interfaces.Level = &MockLevel{}

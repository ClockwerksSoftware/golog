package internal

import (
	"github.com/ClockwerksSoftware/golog/interfaces"
)

type MockLocation struct {
	MockFilename string
	MockLine     int
	MockValid    bool
	MockStack    string
}

func (ml *MockLocation) Filename() string {
	return ml.MockFilename
}
func (ml *MockLocation) Line() int {
	return ml.MockLine
}
func (ml *MockLocation) Valid() bool {
	return ml.MockValid
}
func (ml *MockLocation) Stack() string {
	return ml.MockStack
}

var _ interfaces.RecordLocation = &MockLocation{}

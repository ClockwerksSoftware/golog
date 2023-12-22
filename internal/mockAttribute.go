package internal

import (
	"github.com/ClockwerksSoftware/golog/interfaces"
)

type MockAttribute struct {
	MockKey   string
	MockValue any
}

func (ma *MockAttribute) Key() string {
	return ma.MockKey
}

func (ma *MockAttribute) Value() any {
	return ma.MockValue
}

var _ interfaces.Attribute = &MockAttribute{}

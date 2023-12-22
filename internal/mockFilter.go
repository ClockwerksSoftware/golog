package internal

import (
	"fmt"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

type MockFilter struct {
	AllowRecord bool
	Called      bool
	CalledWith  []interfaces.Record
}

func (mf *MockFilter) Name() string {
	return "mock-filter"
}

func (mf *MockFilter) Filter(r interfaces.Record) (interfaces.Record, bool) {
	mf.Called = true
	mf.CalledWith = append(mf.CalledWith, r)
	if mf.AllowRecord {
		return r, true
	}

	return nil, false
}

func (mf *MockFilter) GetCalledWith(index int) (error, interfaces.Record) {
	if len(mf.CalledWith) >= index {
		return nil, mf.CalledWith[index]
	}
	return fmt.Errorf(
		"%w: invalid index: %d",
		ErrInvalidCallIndex,
		index,
	), nil
}

var _ interfaces.Filter = &MockFilter{}

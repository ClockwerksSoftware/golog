package internal

import (
	"fmt"
	"time"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

type MockRecord struct {
	ByteFormatCache    map[string][]byte
	StrFormatCache     map[string]string
	MockLocation       *MockLocation
	MockLevel          *MockLevel
	MockName           string
	MockRawMessage     string
	MockRawMessageArgs []any
	MockTime           time.Time
	MockAttributes     []interfaces.Attribute
}

func NewMockRecord() *MockRecord {
	return &MockRecord{
		ByteFormatCache: make(map[string][]byte),
		StrFormatCache:  make(map[string]string),
		MockLocation: &MockLocation{
			MockFilename: "foo",
			MockLine:     23950,
			MockValid:    true,
			MockStack:    "some stack",
		},
		MockName: "mock record",
		MockLevel: &MockLevel{
			StrValue: "bar",
			IntValue: 99,
		},
		MockRawMessage:     "some message",
		MockRawMessageArgs: make([]any, 0),
		MockTime:           time.Now().UTC(),
		MockAttributes:     make([]interfaces.Attribute, 0),
	}
}

func (mr *MockRecord) CacheFormat(formatName string, formattedRecord []byte) {
	mr.ByteFormatCache[formatName] = formattedRecord
}

func (mr *MockRecord) CacheFormatString(formatName string, formattedRecord string) {
	mr.StrFormatCache[formatName] = formattedRecord
}

func (mr *MockRecord) GetCacheFormat(formatName string) ([]byte, error) {
	if r, ok := mr.ByteFormatCache[formatName]; ok {
		return r, nil
	}
	return nil, fmt.Errorf(
		"%w: no cached []byte record",
		interfaces.ErrNoCachedRecord,
	)
}

func (mr *MockRecord) GetCacheFormatString(formatName string) (string, error) {
	if r, ok := mr.StrFormatCache[formatName]; ok {
		return r, nil
	}
	return "", fmt.Errorf(
		"%w: no cached string record",
		interfaces.ErrNoCachedRecord,
	)
}

func (mr *MockRecord) Location() interfaces.RecordLocation {
	return mr.MockLocation
}

func (mr *MockRecord) Name() string {
	return mr.MockName
}

func (mr *MockRecord) Level() interfaces.Level {
	return mr.MockLevel
}

func (mr *MockRecord) RawMessage() string {
	return mr.MockRawMessage
}

func (mr *MockRecord) RawMessageArgs() []any {
	return mr.MockRawMessageArgs
}

func (mr *MockRecord) Time() time.Time {
	return mr.MockTime
}

func (mr *MockRecord) AddAttributes(attr ...interfaces.Attribute) {
	mr.MockAttributes = append(mr.MockAttributes, attr...)
}

func (mr *MockRecord) Attributes() []interfaces.Attribute {
	return mr.MockAttributes
}

var _ interfaces.Record = &MockRecord{}

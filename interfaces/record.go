package interfaces

import (
	"errors"
	"time"
)

var (
	ErrNoCachedRecord = errors.New("no cached log record")
)

type RecordFormatCache interface {
	CacheFormat(formatName string, formattedRecord []byte)
	CacheFormatString(formatName string, formattedRecord string)
	GetCacheFormat(formatName string) ([]byte, error)
	GetCacheFormatString(formatName string) (string, error)
}

type Record interface {
	RecordFormatCache

	Location() RecordLocation
	Name() string
	Level() Level
	RawMessage() string
	RawMessageArgs() []any
	Time() time.Time

	AddAttributes(attr ...Attribute)
	Attributes() []Attribute
}

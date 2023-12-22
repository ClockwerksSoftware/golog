package interfaces

import (
	"errors"
	"time"
)

var (
	ErrNoCachedRecord = errors.New("no cached log record")
)

type RecordFormatByteCache interface {
	CacheFormat(formatName string, formattedRecord []byte)
	GetCacheFormat(formatName string) ([]byte, error)
}

type RecordFormatStringCache interface {
	CacheFormatString(formatName string, formattedRecord string)
	GetCacheFormatString(formatName string) (string, error)
}

type RecordFormatCache interface {
	RecordFormatByteCache
	RecordFormatStringCache
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

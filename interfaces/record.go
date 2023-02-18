package interfaces

import (
    "time"
)

type RecordFormatCache interface {
    CacheFormat(formatName string, formattedRecord []byte)
    CacheFormatString(formatName string, formattedRecord string)
    GetCacheFormat(formatName string) []byte
    GetCacheFormatString(formatName string) string
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

package interfaces

type Formatter interface {
    Name() string
    Format(r Record) []byte
    FormatString(r Record) string
}

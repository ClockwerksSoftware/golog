package interfaces

type RecordLocation interface {
    Filename() string
    Line() int
    Valid() bool
    Stack() string
}

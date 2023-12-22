package interfaces

type Filter interface {
	Name() string
	Filter(r Record) (Record, bool)
}

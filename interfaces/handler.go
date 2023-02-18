package interfaces

import (
    "io"
)

type Handler interface {
    Name() string
    SetFilter(f Filter)
    SetFormatter(f Formatter)
    SetOutput(w io.Writer)
    Handle(r Record)
}

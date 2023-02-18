package handler

import (
    "io"
    "os"

    "github.com/ClockwerksSoftware/golog/interfaces"
    "github.com/ClockwerksSoftware/golog/formatter"
)

type logHandler struct {
    name string
    filter interfaces.Filter
    formatter interfaces.Formatter
    writer io.Writer
}

func NewLogHandler() interfaces.Handler {
    return &logHandler{
        name: "basicHandler",
        writer: os.Stderr,
        formatter: formatter.NewLogFormatter(),
    }
}

func (lh *logHandler) Name() string {
    return lh.name
}

func (lh *logHandler) SetFilter(f interfaces.Filter) {
    lh.filter = f
}
func (lh *logHandler) SetFormatter(f interfaces.Formatter) {
    lh.formatter = f
}
func (lh *logHandler) SetOutput(w io.Writer) {
    lh.writer = w
}

func (lh *logHandler) Handle(r interfaces.Record) {
    // validate that a record came through
    if r == nil {
        // nothing to do
        return
    }

    // if there is a filter then apply it and only continue if a record comes back
    if lh.filter != nil {
        if _, ok := lh.filter.Filter(r); !ok {
            // nothing more to do - filter says don't continue
            return
        }
    }

    // format the record for output
    formattedOutput := lh.formatter.Format(r)

    // send it to the output stream
    _, _ = lh.writer.Write(formattedOutput)
}

var _ interfaces.Handler = &logHandler{}

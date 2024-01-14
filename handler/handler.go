package handler

import (
	"fmt"
	"io"
	"os"

	"github.com/ClockwerksSoftware/golog/formatter"
	"github.com/ClockwerksSoftware/golog/interfaces"
)

const (
	basicHandlerName = "basicHandler"
)

type logHandler struct {
	name      string
	children  []interfaces.Handler
	filters   []interfaces.Filter
	formatter interfaces.Formatter
	writer    io.Writer
}

func New() interfaces.Handler {
	return &logHandler{
		name:      basicHandlerName,
		children:  make([]interfaces.Handler, 0),
		filters:   make([]interfaces.Filter, 0),
		writer:    os.Stderr,
		formatter: formatter.NewLogFormatter(),
	}
}

func (lh *logHandler) Name() string {
	return lh.name
}

func (lh *logHandler) ChildHandlers() []interfaces.Handler {
	return lh.children
}

func (lh *logHandler) AddChildHandler(h interfaces.Handler) {
	lh.children = append(lh.children, h)
}

func (lh *logHandler) Filters() []interfaces.Filter {
	return lh.filters
}

func (lh *logHandler) AddFilter(f interfaces.Filter) {
	lh.filters = append(lh.filters, f)
}

func (lh *logHandler) Formatter() interfaces.Formatter {
	return lh.formatter
}

func (lh *logHandler) SetFormatter(f interfaces.Formatter) {
	lh.formatter = f
}

func (lh *logHandler) Output() io.Writer {
	return lh.writer
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
	for _, filter := range lh.filters {
		if _, ok := filter.Filter(r); !ok {
			// filter blocks it, nothing more to do
			return
		}
	}

	// format the record for output
	var formattedOutput []byte
	if lh.formatter != nil {
		formattedOutput = lh.formatter.Format(r)
	} else {
		formattedOutput = []byte(
			fmt.Sprintf(
				r.RawMessage(),
				r.RawMessageArgs()...,
			),
		)
	}

	// send it to the output stream
	if lh.writer != nil {
		_, _ = lh.writer.Write(formattedOutput)
	}

	// finally send it to any child handlers
	for _, handler := range lh.children {
		handler.Handle(r)
	}
}

var _ interfaces.Handler = &logHandler{}

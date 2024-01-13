package internal

import (
	"bytes"
	"fmt"
	"io"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

type MockHandler struct {
	name         string
	children  []interfaces.Handler
	filters   []interfaces.Filter
	formatter interfaces.Formatter
	writer    io.Writer
	Buffer    *bytes.Buffer
}

func NewMockHandler(name string) *MockHandler {
	mh := &MockHandler{
		name: name,
		children: make([]interfaces.Handler, 0),
		filters: make([]interfaces.Filter, 0),
		formatter: nil,
		Buffer: &bytes.Buffer{},
	}
	mh.writer = mh.Buffer
	return mh
}

func (mh *MockHandler) Name() string {
	return mh.name
}
func (mh *MockHandler) ChildHandlers() []interfaces.Handler {
	return mh.children
}
func (mh *MockHandler) AddChildHandler(h interfaces.Handler) {
	mh.children = append(mh.children, h)
}
func (mh *MockHandler) Filters() []interfaces.Filter {
	return mh.filters
}
func (mh *MockHandler) AddFilter(f interfaces.Filter) {
	mh.filters = append(mh.filters, f)
}
func (mh *MockHandler) Formatter() interfaces.Formatter {
	return mh.formatter
}
func (mh *MockHandler) SetFormatter(f interfaces.Formatter) {
	mh.formatter = f
}
func (mh *MockHandler) Output() io.Writer {
	return mh.writer
}
func (mh *MockHandler) SetOutput(w io.Writer) {
	mh.writer = w
}
func (mh *MockHandler) Handle(r interfaces.Record) {
	if r == nil {
		return
	}

	for _, filter := range mh.filters {
		if _, ok := filter.Filter(r); !ok {
			return
		}
	}

	var formattedOutput []byte
	if mh.formatter != nil {
		formattedOutput = mh.formatter.Format(r)
	} else {
		formattedOutput = []byte(
			fmt.Sprintf(
				r.RawMessage(),
				r.RawMessageArgs()...,
			),
		)
	}

	if mh.writer != nil {
		_, _ = mh.writer.Write(formattedOutput)
	}
	for _, handler := range mh.children {
		handler.Handle(r)
	}
}

var _ interfaces.Handler = &MockHandler{}

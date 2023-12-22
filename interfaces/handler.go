package interfaces

import (
	"io"
)

// Handler provides a means of processing log records through a series of
//	filters and potentially multiple downstream handlers
type Handler interface {
	// Name returns a name for the handler
	Name() string
	// GetChildHandlers returns the list of downstream log handlers
	ChildHandlers() []Handler
	// AddChildHandler adds a handler to the list of downstream log handlers
	AddChildHandler(h Handler)
	// Filters return the list of filters that will be applied by the handler instance
	Filters() []Filter
	// Add a filter to the list of filters applied by the handler instance
	AddFilter(f Filter)
	// Formatter returns the formatter used by the handler instance
	Formatter() Formatter
	// SetFormatter set the formatter instance used to output to the handler instance's output
	SetFormatter(f Formatter)
	// Output returns the writer instance used for storing the log record
	Output() io.Writer
	// SetOutput sets the writer instance used for storing the log record
	SetOutput(w io.Writer)
	// Handle processes the log record and decides whether or not to record it
	//	based on the associated log filters and formatter along with any
	//	downstream handlers
	Handle(r Record)
}

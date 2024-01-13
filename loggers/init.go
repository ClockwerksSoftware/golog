package loggers

import (
	"fmt"
	"os"

	"github.com/ClockwerksSoftware/golog/handler"
	"github.com/ClockwerksSoftware/golog/interfaces"
)

const (
	rootLoggerName      = "root"
	alternateLoggerName = "main"
)

var root *coreLogger

func init() {
	// logger is central to a program so needs to be reliably
	// there with minimal work for the application
	root = &coreLogger{
		name:         rootLoggerName,
		handlers:     make([]interfaces.Handler, 0),
		childLoggers: make(map[string]interfaces.Log),
	}
	if defaultHandler := handler.New(); defaultHandler != nil {
		root.AddHandler(defaultHandler)
	} else {
		fmt.Fprintf(os.Stderr, "**********************************\nFailed to allocate default handler\n**********************************\n")
	}
}

// DefaultLogger is a quick way to get the default logger
// the same could be done using `main` or `root` as a parameters to `GetLogger`
func DefaultLogger() (cl interfaces.Log) {
	return root
}

// GetLogger provides a way to get any named logger
// if the name is `main` or `root` then the default loggers is returned
// any other name generates a new child logger of the default logger
func GetLogger(name string) (cl interfaces.Log) {
	if name == rootLoggerName || name == alternateLoggerName {
		cl = root
		return
	}

	// check if there is a logger already allocated
	if v, ok := root.childLoggers[name]; ok {
		cl = v
		return
	}

	// allocate a new logger
	cl = NewChildLogger(name, root)
	root.childLoggers[name] = cl
	return
}

// RemoveLoggers drops a child logger from the instances
func RemoveLogger(cl interfaces.Log) {
	if v, ok := root.childLoggers[cl.Name()]; ok {
		if v.Name() == cl.Name() {
			delete(root.childLoggers, cl.Name())
		}
	}
}

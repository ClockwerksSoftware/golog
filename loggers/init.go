package loggers

import (
    "github.com/ClockwerksSoftware/golog/interfaces"
)

const (
    rootLoggerName = "root"
    alternateLoggerName = "main"
)

var root *coreLogger

func init() {
    // logger is central to a program so needs to be reliably
    // there with minimal work for the application
    root = &coreLogger{
        name: rootLoggerName,
        handlers: make([]interfaces.Handler, 0),
        childLoggers: make(map[string]interfaces.Log),
    }
}

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

func RemoveLogger(cl interfaces.Log) {
    if v, ok := root.childLoggers[cl.Name()]; ok {
        if v.Name() == cl.Name() {
            delete(root.childLoggers, cl.Name()) 
        }
    }
}

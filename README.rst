golog
=====

Golog uses a design similar to that of the Python Logger module but implemented in how Golang works.
That is, the general Logger, Handler, Formatter, and Record design in Python Logger are excellent and
language agnostic; however the specifics of the interactions between those parts are very Go-oriented.

LogRecord
---------

The LogRecord object records the various attributes of any given log message.

Logger
------

The Logger provides the core functionality and is where applications will actually interface with the
system, providing interfaces for the various log levels, and building the Log Record to hand off to
the internals for actual processing.

The Logger comes in two forms:
- Standard Core Logger which provides the basic functionality and acts as the root of the logger.
- Package Logger which provides an interface for individual Golang Packages to namespace themselves

Handler
-------

The LogHandler decides where to send the LogRecord. This could be to a file, to an external service,
or anywhere desired. The Handler uses a LogFormatter and a LogFilter to help in its decision making.

NOTE: The LogHandler itself should only focus on sending the record to its destination.

Filter
------

The LogFilter applies various attributes to decide which LogRecords get selected for output.

Formatter
---------

The LogFormatter is responsible for taking the LogRecord and generating the output format.
Formatting should not effect the LogFormatter at all, and it should cache a copy of the formatted
record into the LogRecord so formatting only has to happen once.

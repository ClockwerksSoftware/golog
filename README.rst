golog
=====

`golog` is a Go-specific implementation based on the design used by the Python Log System.
While it implements a Python design it is very much a Go-based logging system.

Why use the Python Log System Design?
-------------------------------------

Log systems are hard to design and hard to design well. There are many various methods out there.
Go has attempted several designs including the `log` and `log/slog` native packages provided
by Go itself. However, there are still lacking features necessary for high quality logging
systems.

The question is actually a bit of a misnomer as it highlights that the person asking it is focused
more on the fact that the design is borrowed from Python than the design itself. It should not
matter where the general design was taken from, or what language implemented it previously.
What should matter is whether the design provides meets the needs of applications.

This logger provides several components:

1. A general Log Record
2. A Log Formatter specification for generating the output data
3. A Log Filter specification for determing what logs make it to the output
4. A Log Handler specification for creating log output streams

These concepts are all language and technology agnostic. Their implementations will
certainly be tied to the details of programming languages and the technology stacks
they are built upon, but the general concepts stand firm.

.. note:: `log/slog` does implement some of these concepts; however, the
   relationships are slightly wrong and key implementation details are left to
   being implemented by each `slog` handler. Therefore application using `slog`
   need to be a lot more aware of the logging implementation details than should
   be necessary.

The Log system defined here makes heavy use of Go Interfaces which will make it
easier to swap around components and extend it for more functionality.

General Log Record
------------------

The General Log Record is defined by the `Record` interface and defines some
basic attributes that should be associated with any log record. The basics of
this interface can be relied upon for anything in the other parts of the Log System.

Log Formatter
-------------

The Log Formatter is defined by the `Formatter` interface. Formatters must implement
two forms of output: `[]byte` data for binary outputs and `string` data for simpler
string formats. Callers should use the one most appropriate for where they are
sending the data.

Binary data is useful when transferring over the wire to other systems; while
String data is useful when recording to files for simple output.

Log Formatters themselves should be idempotent; however, they may modify the Log
Record only through either (a) calls to the Log Record that cause it to modify
itself or (b) adding a cached copy of the formatted record output in order to
enable one-time formatting for any given log record/format combination.

Log Filter
----------

The Log Filter is defined by the `Filter` interface. It provides a mechanism for
Log Handlers to limit what flows down to its child handlers and any output streams.
Log Filters must not modify the Log Record, be idempotent, and should be aim to
operate in constant time.

Log Handlers
------------

Log Handlers are the core of the Log System along with Log Records. Log Handlers
are responsible for making the decisions of where Log Records go. Log Handlers must
account for race conditions, copying, etc that are implied by their design and should
not burden applications, log records, filters, or formatters with those details.

Log Handlers should themselves be idempotent during logging operations.

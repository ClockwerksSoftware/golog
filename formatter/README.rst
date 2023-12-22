Formatter
=========

The `Formatter` interface provides the ability to reformat log messages
into a given specification for output.

The formatter should be called as soon as it is decided that the log
record will actually be recorded to disk with the result cached until
the log record is disposed. Using this rule means that the penalty for
formatting the record is only incurred a single time if at all.

.. block-quote:: Multiple formatters may be called against a single Log Record
    depending on how many outputs there are.

Log Formatters for a basic method and map-based JSON method are provided.

Log Filter
==========

The `Filter` interface provides the ability to for applications to limit
what goes into a given log stream. The default filter provided here
supplies the functionality to limit the log stream data based on the log level.

Handlers apply the filters to decide what logs they handle and what logs
are received by their downstream handlers. In other words, a handler will reduce
the logs going to any handlers attached to it.

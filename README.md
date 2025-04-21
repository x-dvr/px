# PX - Process eXtended
PX is a thin wrapper around [os.Process](https://pkg.go.dev/os#Process) with additional functionality to control lifetime and mange standard streams.

TODO:
 - [ ] implement basic os.Process functionality (kill, signal, ...)
 - [ ] implement killAll to kill process and all its children
 - [ ] implement signalAll to send signal to the process and all children
 - [ ] implement net.Connection for stdin/stdout pair

package px

import (
	"os"
	"path/filepath"
	"sync"
)

// Process represents system process, similar to the os.Process
// from the Go standard library.
type Process struct {
	internal os.Process

	// path represents path to the executable
	path string

	// args represents command line arguments passed to an executable on process start
	args []string

	// cwd represents current working directory
	cwd string

	// env represents list of environment variables set for executable.
	// Each item has format NAME=VALUE
	env []string

	*sync.RWMutex
	name string
}

type procOptions func(*Process)

// Start starts a new process with the path to executable and options
func Start(path string, opts ...procOptions) (*Process, error) {
	exe := Process{
		path:    path,
		RWMutex: &sync.RWMutex{},
		name:    filepath.Base(path),
	}

	// setNewProcessGroupAttr(&exe.sysProcAttr)

	for _, opt := range opts {
		opt(&exe)
	}

	return &exe, nil
}

// WithArgs sets the command line arguments of the executable
func WithArgs(args ...string) procOptions {
	return func(p *Process) {
		p.args = args
	}
}

// WithWD sets working directory of the executable
func WithWD(wd string) procOptions {
	return func(p *Process) {
		p.cwd = wd
	}
}

// WithEnv sets environment variables of the executable
func WithEnv(env []string) procOptions {
	return func(p *Process) {
		p.env = env
	}
}

// StdinPipe, StdoutPipe and StderrPipe

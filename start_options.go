package px

import "io"

// WithArgs sets the command line arguments of the executable
func WithArgs(args ...string) startOptFn {
	return func(s *startOptions) {
		s.Args = args
	}
}

// WithWD sets working directory of the executable
func WithWD(wd string) startOptFn {
	return func(s *startOptions) {
		s.Cwd = wd
	}
}

// WithEnv sets environment variables of the executable
func WithEnv(env []string) startOptFn {
	return func(s *startOptions) {
		s.Env = env
	}
}

// WithStdin sets the process's standard input
func WithStdin(stdin io.Reader) startOptFn {
	return func(s *startOptions) {
		s.Stdin = stdin
	}
}

// WithStdout sets the process's standard output
func WithStdout(stdout io.Writer) startOptFn {
	return func(s *startOptions) {
		s.Stdout = stdout
	}
}

// WithStderr sets the process's standard error output
func WithStderr(stderr io.Writer) startOptFn {
	return func(s *startOptions) {
		s.Stderr = stderr
	}
}

// WithStdio sets the process's standard input, output and error output
func WithStdio(stdin io.Reader, stdout, stderr io.Writer) startOptFn {
	return func(s *startOptions) {
		s.Stdin = stdin
		s.Stdout = stdout
		s.Stderr = stderr
	}
}

// WithStdinPipe creates pipe for the process's standard input
func WithStdinPipe() startOptFn {
	return func(s *startOptions) {
		s.PipeStdin = true
	}
}

// WithStdoutPipe creates pipe for the process's standard output
func WithStdoutPipe() startOptFn {
	return func(s *startOptions) {
		s.PipeStdout = true
	}
}

// WithStderrPipe creates pipe for the process's standard error
func WithStderrPipe() startOptFn {
	return func(s *startOptions) {
		s.PipeStderr = true
	}
}

// WithStdioPipe creates pipes for the process's standard input, output and error
func WithStdioPipe() startOptFn {
	return func(s *startOptions) {
		s.PipeStdin = true
		s.PipeStdout = true
		s.PipeStderr = true
	}
}

type startOptions struct {
	// Args represents command line arguments passed to an executable on process start
	Args []string

	// Cwd represents current working directory
	Cwd string

	// Env represents list of environment variables set for executable.
	// Each item has format NAME=VALUE
	Env []string

	// Stdin specifies the process's standard input.
	Stdin io.Reader
	// PipeStdin signals that instead of using Stdin field new pipe should be created.
	// This pipe will be connected to the process's standard input.
	// Writer of this pipe can be accessed as the field `Process.Stdin`
	PipeStdin bool

	// Stdout and Stderr specify the process's standard output and error.
	Stdout io.Writer
	Stderr io.Writer
	// PipeStdout and PipeStderr signal that new pipes should be created and connected to
	// process's stdout and stderr
	PipeStdout bool
	PipeStderr bool
}

type startOptFn func(*startOptions)

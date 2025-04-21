package px

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// Process represents system process, similar to the os.Process
// from the Go standard library.
type Process struct {
	internal *os.Process
	stdin    io.WriteCloser
	stdout   io.ReadCloser
	stderr   io.ReadCloser
	state    *os.ProcessState
	stateCh  chan *os.ProcessState
	doneCh   chan struct{}
}

// Find searches for a process by pid
func Find(pid int) (*Process, error) {
	p, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}

	return &Process{
		internal: p,
	}, nil
}

// Start starts a new process with the path to executable and options
func Start(path string, opts ...startOptFn) (*Process, error) {
	po := startOptions{}
	for _, opt := range opts {
		opt(&po)
	}

	cmd := exec.Command(path, po.Args...)
	if po.Cwd != "" {
		cmd.Dir = po.Cwd
	}
	if po.Env != nil {
		cmd.Env = po.Env
	}
	cmd.Stdin = po.Stdin
	cmd.Stdout = po.Stdout
	cmd.Stderr = po.Stderr

	p := Process{}
	var err error
	if po.PipeStdin {
		p.stdin, err = cmd.StdinPipe()
		if err != nil {
			p.closeIOStreams()
			return nil, fmt.Errorf("prepare stdin pipe: %w", err)
		}
	}

	if po.PipeStdout {
		p.stdout, err = cmd.StdoutPipe()
		if err != nil {
			p.closeIOStreams()
			return nil, fmt.Errorf("prepare stdout pipe: %w", err)
		}
	}

	if po.PipeStderr {
		p.stderr, err = cmd.StderrPipe()
		if err != nil {
			p.closeIOStreams()
			return nil, fmt.Errorf("prepare stderr pipe: %w", err)
		}
	}

	setNewProcessGroupAttr(cmd.SysProcAttr)

	if err := cmd.Start(); err != nil {
		p.closeIOStreams()
		return nil, fmt.Errorf("start process: %w", err)
	}

	p.internal = cmd.Process
	p.doneCh = make(chan struct{})
	p.stateCh = make(chan *os.ProcessState, 1)

	go waitFor(cmd, p.doneCh, p.stateCh)

	return &p, nil
}

// Done returns a channel that's closed when the process stops running.
func (p *Process) Done() <-chan struct{} {
	return p.doneCh
}

// Wait waits for the process to exit and waits for any copying to
// stdin or copying from stdout or stderr to complete.
//
// returns state of the process
func (p *Process) Wait() *os.ProcessState {
	if p.state == nil {
		<-p.doneCh
		p.state = <-p.stateCh
	}
	return p.state
}

func (p *Process) closeIOStreams() error {
	errs := make([]error, 0, 3)

	if p.stdin != nil {
		errs = append(errs, p.stdin.Close())
	}
	if p.stdout != nil {
		errs = append(errs, p.stdout.Close())
	}
	if p.stderr != nil {
		errs = append(errs, p.stderr.Close())
	}

	return errors.Join(errs...)
}

func waitFor(cmd *exec.Cmd, doneCh chan struct{}, stateCh chan<- *os.ProcessState) {
	defer func() {
		stateCh <- cmd.ProcessState
		close(doneCh)
	}()

	cmd.Wait()

	// if err := cmd.Wait(); err != nil {
	// 	var exitError *exec.ExitError
	// 	if errors.As(err, &exitError) {
	// 		exitError.Sys()
	// 	}
	// }
}

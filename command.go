package command

import (
	"io"
	"os/exec"
)

// Cli presents a cli API.
type Cli interface {

	// Command returns a Command instance which can be used to run a single command.
	Command(cmd string, args ...string) Command
}

// Command presents cli command
type Command interface {

	// CombinedOutput runs the command and returns its combined standard
	// output and standard error.
	CombinedOutput() ([]byte, error)

	//OutputWithStdin runs the command with user give reader and returns its combined standard
	// output and standard error.
	OutputWithStdin(io.Reader) ([]byte, error)
}

type command struct{}

// New returns a new Interface which will os/exec to run commands.
func New() Cli {
	return &command{}
}

// Command is part of the Interface interface.
func (command *command) Command(cmd string, args ...string) Command {
	cli := exec.Command(cmd, args...)
	return &cmdWrapper{cli}
}

type cmdWrapper struct {
	*exec.Cmd
}

var _ Command = &cmdWrapper{}

// CombinedOutput is part of the Command interface.
func (cmd *cmdWrapper) CombinedOutput() ([]byte, error) {
	return cmd.Cmd.CombinedOutput()
}

// OutputWithStdin is part of the Command interface.
// The provided reader is used to command Stdin
func (cmd *cmdWrapper) OutputWithStdin(reader io.Reader) ([]byte, error) {
	stdin, err := cmd.Cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	go func() {
		defer stdin.Close()
		io.Copy(stdin, reader)
	}()
	return cmd.Cmd.CombinedOutput()
}

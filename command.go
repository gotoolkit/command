package command

import (
	"io"
	"log"
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

	// OutputWithStdin runs the command and returns its combined standard
	// output and standard error.
	// The provided reader is used to the command Stdin
	OutputWithStdin(io.Reader) ([]byte, error)

	//WaitStdout runs the command and returns its standard error.
	// The provided writer is used to the command Stdout
	WaitStdout(io.Writer) error
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
func (cmd *cmdWrapper) OutputWithStdin(reader io.Reader) ([]byte, error) {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	go func() {
		defer stdin.Close()
		io.Copy(stdin, reader)
	}()
	return cmd.Cmd.CombinedOutput()
}

// WaitStdout is part of the Command interface.
func (cmd *cmdWrapper) WaitStdout(writer io.Writer) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(writer, stdout); err != nil {
		return err
	}
	return cmd.Wait()
}

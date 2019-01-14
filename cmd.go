package firebasecli

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// ErrUnknownCommand is returned when a given command is unknown/undefined.
var ErrUnknownCommand = errors.New("unknown command")

// Cmd executes command.
type Cmd struct {
	App    *App
	Stdout io.Writer
	Stderr io.Writer
}

// Print prints to stdout.
func (c *Cmd) Print(a ...interface{}) (int, error) {
	return fmt.Fprint(c.getStdout(), a...)
}

// Println prints to stdout.
func (c *Cmd) Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(c.getStdout(), a...)
}

// Printf prints to stdout.
func (c *Cmd) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(c.getStdout(), format, a...)
}

// Eprint prints to stderr.
func (c *Cmd) Eprint(a ...interface{}) (int, error) {
	return fmt.Fprint(c.getStderr(), a...)
}

// Eprintln prints to stderr.
func (c *Cmd) Eprintln(a ...interface{}) (int, error) {
	return fmt.Fprintln(c.getStderr(), a...)
}

// Eprintf prints to stderr.
func (c *Cmd) Eprintf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(c.getStderr(), format, a...)
}

func (c *Cmd) getStdout() io.Writer {
	if c.Stdout != nil {
		return c.Stdout
	}
	return os.Stdout
}

func (c *Cmd) getStderr() io.Writer {
	if c.Stderr != nil {
		return c.Stderr
	}
	return os.Stderr
}

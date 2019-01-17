package firebasecli

import (
	"fmt"
	"io"
	"os"
)

// Cmd executes command.
type Cmd struct {
	// Sub holds sub commands.
	Sub map[string]Runnable

	// App holds App.
	App *App

	// Stdout holds standard output where a command outputs.
	// Set nil to suppress output.
	Stdout io.Writer

	// Stderr holds standard error where a command outputs.
	// Set nil to suppress output.
	Stderr io.Writer
}

// NewCmd returns a new initialized Cmd.
func NewCmd() *Cmd {
	return &Cmd{
		make(map[string]Runnable),
		&App{},
		os.Stdout,
		os.Stderr,
	}
}

// Print prints to stdout.
func (c *Cmd) Print(a ...interface{}) (int, error) {
	if c.Stdout == nil {
		return 0, nil
	}
	return fmt.Fprint(c.Stdout, a...)
}

// Println prints to stdout.
func (c *Cmd) Println(a ...interface{}) (int, error) {
	if c.Stdout == nil {
		return 0, nil
	}
	return fmt.Fprintln(c.Stdout, a...)
}

// Printf prints to stdout.
func (c *Cmd) Printf(format string, a ...interface{}) (int, error) {
	if c.Stdout == nil {
		return 0, nil
	}
	return fmt.Fprintf(c.Stdout, format, a...)
}

// Eprint prints to stderr.
func (c *Cmd) Eprint(a ...interface{}) (int, error) {
	if c.Stderr == nil {
		return 0, nil
	}
	return fmt.Fprint(c.Stderr, a...)
}

// Eprintln prints to stderr.
func (c *Cmd) Eprintln(a ...interface{}) (int, error) {
	if c.Stderr == nil {
		return 0, nil
	}
	return fmt.Fprintln(c.Stderr, a...)
}

// Eprintf prints to stderr.
func (c *Cmd) Eprintf(format string, a ...interface{}) (int, error) {
	if c.Stderr == nil {
		return 0, nil
	}
	return fmt.Fprintf(c.Stderr, format, a...)
}

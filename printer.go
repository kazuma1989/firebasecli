package firebasecli

import (
	"fmt"
	"io"
)

// Printer prints to Stdout/Stderr.
type Printer struct {
	Stdout io.Writer
	Stderr io.Writer
}

// Print prints to stdout.
func (p *Printer) Print(a ...interface{}) (int, error) {
	return fmt.Fprint(p.Stdout, a...)
}

// Println prints to stdout.
func (p *Printer) Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(p.Stdout, a...)
}

// Printf prints to stdout.
func (p *Printer) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(p.Stdout, format, a...)
}

// Eprint prints to stderr.
func (p *Printer) Eprint(a ...interface{}) (int, error) {
	return fmt.Fprint(p.Stderr, a...)
}

// Eprintln prints to stderr.
func (p *Printer) Eprintln(a ...interface{}) (int, error) {
	return fmt.Fprintln(p.Stderr, a...)
}

// Eprintf prints to stderr.
func (p *Printer) Eprintf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(p.Stderr, format, a...)
}

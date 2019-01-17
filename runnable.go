package firebasecli

// Runnable runs a command.
type Runnable interface {
	Run(...string) error
}

// RunnableFunc is a func which implements Runnable.
type RunnableFunc func(...string) error

// Run implements Runnable.
func (f RunnableFunc) Run(args ...string) error {
	return f()
}

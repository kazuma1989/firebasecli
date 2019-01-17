package firebasecli_test

import (
	"fmt"

	"github.com/kazuma1989/firebasecli"
)

// Define myCmd, extending firebasecli.Cmd.
type myCmd firebasecli.Cmd

// Override Run method.
func (c *myCmd) Run(args ...string) error {
	if runnable, ok := c.Sub[args[1]]; ok {
		return runnable.Run(args[1:]...)
	}
	return fmt.Errorf("failure from Run")
}

func ExampleCmd_Run_impl_myCmd() {
	cmdFooBar := func(args ...string) error {
		fmt.Println(args)
		return fmt.Errorf("failure from foo bar")
	}

	cmdFoo := &myCmd{
		Sub: make(map[string]firebasecli.Runnable),
	}
	cmdFoo.Sub["bar"] = firebasecli.RunnableFunc(cmdFooBar)

	cmd := firebasecli.NewCmd()
	cmd.Sub["foo"] = cmdFoo

	cmd.Run("foo", "bar")
	// Output: [bar]
}

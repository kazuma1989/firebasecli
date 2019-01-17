package firebasecli_test

import (
	"fmt"

	"github.com/kazuma1989/firebasecli"
)

type myCmd firebasecli.Cmd

func (c *myCmd) Run(args ...string) error {
	if runnable, ok := c.Sub[args[1]]; ok {
		return runnable.Run(args[1:]...)
	}
	return fmt.Errorf("failure")
}

func Example_impl_myCmd() {
	cmdFoo := &myCmd{
		Sub: make(map[string]firebasecli.Runnable),
	}

	cmdFooBar := func(args ...string) error {
		fmt.Println(args)
		return fmt.Errorf("FOO BAR")
	}
	cmdFoo.Sub["bar"] = firebasecli.RunnableFunc(cmdFooBar)

	cmd := firebasecli.NewCmd()
	cmd.Sub["foo"] = cmdFoo

	cmd.Run("foo", "bar")
	// Output: [bar]
}

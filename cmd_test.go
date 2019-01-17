package firebasecli_test

import (
	"fmt"

	"github.com/kazuma1989/firebasecli"
)

func ExampleCmd_Run() {
	cmd := firebasecli.NewCmd()
	// For testing, avoid logging into Firebase.
	cmd.App = nil
	cmd.Sub["sub1"] = firebasecli.RunnableFunc(func(args ...string) error {
		fmt.Println(args)
		return nil
	})

	cmd.Run("-c", "opt1", "sub1", "arg1", "arg2")
	// Output: [sub1 arg1 arg2]
}

type myCmd firebasecli.Cmd

func (c *myCmd) Run(args ...string) error {
	if runnable, ok := c.Sub[args[1]]; ok {
		return runnable.Run(args[1:]...)
	}
	return fmt.Errorf("failure")
}

func ExampleCmd_Run_complex() {
	// type myCmd firebasecli.Cmd

	// func (c *myCmd) Run(args ...string) error {
	// 	if runnable, ok := c.Sub[args[1]]; ok {
	// 		return runnable.Run(args[1:]...)
	// 	}
	// 	return fmt.Errorf("failure")
	// }

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

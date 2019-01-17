package firebasecli_test

import (
	"context"
	"fmt"

	"github.com/kazuma1989/firebasecli"
)

func ExampleNewCmd() {
	// Construct a new Cmd.
	cmd := firebasecli.NewCmd()

	// Log in with default credential (e.g. an env var or hard-coded path).
	cmd.App.Login(context.Background(), "")

	// Run some command.
	cmd.Db("list")
}

func ExampleNewCmd_suppressStdout() {
	cmd := firebasecli.NewCmd()
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Run("--help")
	// Output:
}

func ExampleCmd_Run() {
	cmd := firebasecli.NewCmd()
	cmd.Sub["sub1"] = firebasecli.RunnableFunc(func(args ...string) error {
		fmt.Println(args)
		return nil
	})

	// A hack. For testing, avoid logging into Firebase.
	cmd.App = nil
	// "opt1" is consumed as a root option so not passed to the sub command.
	cmd.Run("-c", "opt1", "sub1", "arg1", "arg2")
	// Output: [sub1 arg1 arg2]
}

func ExampleCmd_Run_addYourCommands() {
	firebasecli.DefaultCmd.Sub["foo"] = firebasecli.RunnableFunc(func(args ...string) error {
		fmt.Println(args)
		return nil
	})
	firebasecli.DefaultCmd.Run("foo", "arg1", "arg2")
	// Equivalent to:
	//   firebasecli.Run("foo", "arg1", "arg2")

	// Output: [foo arg1 arg2]
}

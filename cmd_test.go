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

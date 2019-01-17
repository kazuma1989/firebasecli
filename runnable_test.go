package firebasecli_test

import (
	"fmt"

	"github.com/kazuma1989/firebasecli"
)

func ExampleRunnableFunc() {
	cmd := firebasecli.NewCmd()
	cmd.Sub["sub1"] = firebasecli.RunnableFunc(func(args ...string) error {
		fmt.Println(args)
		return nil
	})
	cmd.Run("sub1", "arg1", "arg2")
	// Output: [sub1 arg1 arg2]
}

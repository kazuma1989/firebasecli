package firebasecli_test

import (
	"fmt"
	"os"

	"github.com/kazuma1989/firebasecli"
)

func ExampleRun() {
	err := firebasecli.Run(os.Args[1:]...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

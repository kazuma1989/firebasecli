package firebasecli

import (
	"context"
	"strings"

	docopt "github.com/docopt/docopt-go"
)

// Storage manipulates Cloud Storge
func (c *Cmd) Storage(args ...string) error {
	opts, err := docopt.ParseArgs(`
Manipulate Cloud Storge.

Usage:
  firebasecli storage list

Options:
  list  Show the all objects.
`, args, "")
	if err != nil {
		// TODO add an error explanation.
		return err
	}

	var arg struct {
		Storage bool

		List bool
	}
	if err := opts.Bind(&arg); err != nil {
		// TODO add an error explanation.
		return err
	}

	switch {
	case arg.List:
		return c.storageList()

	default:
		return ErrUnknownCommand
	}
}

func (c *Cmd) storageList() error {
	ctx := context.Background()

	objects, err := c.App.StorageList(ctx)
	if err != nil {
		return err
	}

	c.Println(strings.Join(objects, "\n"))
	return nil
}

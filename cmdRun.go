package firebasecli

import (
	"context"
	"sort"
	"strings"

	docopt "github.com/docopt/docopt-go"
)

// Run starts a sub command.
func (c *Cmd) Run(args ...string) error {
	usage := `Manipulate Firebase as an admin.

Usage:
  firebasecli [-c FILE] [COMMAND ARGS...]
  firebasecli -h

Options:
  -c, --credential FILE  Service account secret key.
                         When omitted, environment variable GOOGLE_APPLICATION_CREDENTIALS will be used,
                         otherwise "serviceAccountKey.json" will be used.
  -h, --help             Show help (this screen).

Available commands:
  `
	var availableCommands []string
	for k := range c.Sub {
		if k == "" {
			continue
		}
		availableCommands = append(availableCommands, k)
	}
	sort.Strings(availableCommands)
	usage += strings.Join(availableCommands, "\n  ")

	allowUnknownFlagsAfterSubcommand := true
	opts, err := (&docopt.Parser{
		OptionsFirst: allowUnknownFlagsAfterSubcommand,
	}).ParseArgs(usage, args, Version)
	if err != nil {
		return ErrFailedToParseArgs
	}

	var arg struct {
		Credential string

		Command string
		Args    []string
	}
	if err := opts.Bind(&arg); err != nil {
		return ErrFailedToParseArgs
	}

	runnable, ok := c.Sub[arg.Command]
	if !ok {
		return ErrUnknownCommand
	}

	if c.App != nil {
		err := c.App.Login(context.Background(), arg.Credential)
		if err != nil {
			return err
		}
	}

	_args := append([]string{arg.Command}, arg.Args...)
	return runnable.Run(_args...)
}

package firebasecli

import (
	"context"
	"os"
	"sort"
	"strings"

	docopt "github.com/docopt/docopt-go"
)

// Commands holds commands.
type Commands map[string]func(*Cmd, ...string) error

// Run starts a sub command.
func (c *Commands) Run(args ...string) error {
	usage := `
Manipulate Firebase as an admin.

Usage:
  firebasecli [-c FILE] COMMAND [ARGS...]
  firebasecli -h

Options:
  -c, --credential FILE  Service account secret key.
                         When omitted, environment variable GOOGLE_APPLICATION_CREDENTIALS will be used,
                         otherwise "serviceAccountKey.json" will be used.
  -h, --help             Show help (this screen).

Available commands:
  `
	var availableCommands []string
	for k := range *c {
		availableCommands = append(availableCommands, k)
	}
	sort.Strings(availableCommands)
	usage += strings.Join(availableCommands, "\n  ")

	allowUnknownFlagsAfterSubcommand := true
	opts, err := (&docopt.Parser{
		OptionsFirst: allowUnknownFlagsAfterSubcommand,
	}).ParseArgs(usage, args, Version)
	if err != nil {
		// TODO add an error explanation.
		return err
	}

	var arg struct {
		Credential string

		Command string
		Args    []string
	}
	if err := opts.Bind(&arg); err != nil {
		// TODO add an error explanation.
		return err
	}

	run, ok := (*c)[arg.Command]
	if !ok {
		return ErrUnknownCommand
	}

	app, err := NewApp(context.Background(), arg.Credential)
	if err != nil {
		return err
	}

	cmd := &Cmd{app, os.Stdout, os.Stderr}
	_args := append([]string{arg.Command}, arg.Args...)
	return run(cmd, _args...)
}

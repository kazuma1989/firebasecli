package firebasecli

import (
	"context"
	"errors"
	"os"

	docopt "github.com/docopt/docopt-go"
)

// Version of the firebasecli.
const Version = "0.0.1"

// Run starts the app.
func Run(args ...string) error {
	opts, err := (&docopt.Parser{
		OptionsFirst: true,
	}).ParseArgs(`
Manipulate Firebase as an admin.

Usage:
  firebasecli [-c FILE] COMMAND [ARGS...]
  firebasecli -h

Options:
  -c, --credential FILE  Service account secret key.
                         When omitted, environment variable GOOGLE_APPLICATION_CREDENTIALS will be used,
                         otherwise "serviceAccountKey.json" will be used.
  -h, --help             Show help (this screen).

Commands:
  auth  Manipulate Firebase Authentication.
  db    Manipulate Cloud Firestore.
`, args, Version)
	if err != nil {
		// TODO add an error explanation.
		return err
	}

	var arg struct {
		Credential string
		Command    string
		Args       []string
	}
	if err := opts.Bind(&arg); err != nil {
		// TODO add an error explanation.
		return err
	}

	app, err := NewApp(context.Background(), arg.Credential)
	if err != nil {
		// TODO add an error explanation.
		return err
	}

	cmd := &Cmd{app, os.Stdout, os.Stderr}
	switch arg.Command {
	case "auth":
		return errors.New("no implementation yet")

	case "db":
		return cmd.Db(arg.Args...)

	default:
		return ErrUnknownCommand
	}
}

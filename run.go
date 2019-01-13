package firebasecli

import (
	"context"
	"encoding/json"
	"os"

	docopt "github.com/docopt/docopt-go"
)

// Run starts the app.
func Run(args ...string) error {
	app, err := NewApp(context.Background(), "")
	if err != nil {
		return err
	}

	opts, err := docopt.ParseArgs(`
Manipulate Cloud Firestore.

Usage:
  firebasecli db list
  firebasecli db export COLLECTIONS... [-f]

Options:
  list  List the all collections.

Options:
  export COLLECTIONS  Export collections as JSON.
  -f, --format        Output as formatted JSON.
`, args, "")
	if err != nil {
		return err
	}

	var arg struct {
		Db          bool
		List        bool
		Export      bool
		Collections []string
		Format      bool
	}
	if err := opts.Bind(&arg); err != nil {
		return err
	}

	p := &Printer{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	ctx := context.Background()

	switch {
	case arg.List:
		collections, err := app.DbList(ctx)
		if err != nil {
			return err
		}
		p.Println(collections)

	case arg.Export:
		data, err := app.DbExport(ctx, arg.Collections)
		if err != nil {
			return err
		}
		var jsonData []byte
		if arg.Format {
			jsonData, err = json.MarshalIndent(data, "", "  ")
		} else {
			jsonData, err = json.Marshal(data)
		}
		if err != nil {
			return err
		}
		p.Println(string(jsonData))
	}

	return nil
}

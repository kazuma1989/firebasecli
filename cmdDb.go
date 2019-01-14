package firebasecli

import (
	"context"
	"encoding/json"
	"strings"

	docopt "github.com/docopt/docopt-go"
)

// Db manipulates Cloud Firestore.
func (c *Cmd) Db(args ...string) error {
	opts, err := docopt.ParseArgs(`
Manipulate Cloud Firestore.

Usage:
  firebasecli db list
  firebasecli db export COLLECTIONS... [-f]

Options:
  list  Show the all collection paths.

Options:
  export COLLECTIONS  Print the all documents in given collections as JSON.
  -f, --format        Output as formatted JSON.
`, args, "")
	if err != nil {
		// TODO add an error explanation.
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
		// TODO add an error explanation.
		return err
	}

	switch {
	case arg.List:
		return c.dbList()

	case arg.Export:
		return c.dbExport(arg.Collections, arg.Format)

	default:
		return ErrUnknownCommand
	}
}

func (c *Cmd) dbList() error {
	ctx := context.Background()

	collections, err := c.App.DbList(ctx)
	if err != nil {
		return err
	}

	c.Println(strings.Join(collections, "\n"))
	return nil
}

func (c *Cmd) dbExport(collections []string, format bool) error {
	ctx := context.Background()

	data, err := c.App.DbExport(ctx, collections)
	if err != nil {
		return err
	}

	var jsonData []byte
	if format {
		jsonData, err = json.MarshalIndent(data, "", "  ")
	} else {
		jsonData, err = json.Marshal(data)
	}
	if err != nil {
		// TODO add an error explanation.
		return err
	}

	c.Println(string(jsonData))
	return nil
}

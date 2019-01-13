package firebasecli

import (
	"context"
	"encoding/json"
	"os"
)

// Run starts the app.
func Run(args ...string) error {
	app, err := NewApp(context.Background(), "")
	if err != nil {
		return err
	}

	p := &Printer{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	ctx := context.Background()
	switch {
	case len(args) <= 1:

	case args[0] == "db" && args[1] == "list":
		collections, err := app.DbList(ctx)
		if err != nil {
			return err
		}
		p.Println(collections)

	case args[0] == "db" && args[1] == "export":
		data, err := app.DbExport(ctx, []string{"meta"})
		if err != nil {
			return err
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		p.Println(string(jsonData))

	default:
	}

	return nil
}

package firebasecli

import (
	"context"
	"os"
)

// Run starts the app.
func Run(args ...string) error {
	app, err := NewApp(context.Background(), "")
	if err != nil {
		return err
	}

	ctx := context.Background()
	collections, err := app.DbList(ctx)

	p := &Printer{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	p.Println(collections)

	return nil
}

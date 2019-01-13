package firebasecli

import (
	"context"
	"os"

	"google.golang.org/api/iterator"
)

// Run starts the app.
func Run() error {
	app, err := NewApp(context.Background(), "")
	if err != nil {
		return err
	}

	ctx := context.Background()
	dbClient, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	defer dbClient.Close()

	var collections []string
	it := dbClient.Collections(ctx)
	for {
		collectionRef, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		collections = append(collections, collectionRef.ID)
	}
	p := &Printer{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	p.Println(collections)

	return nil
}

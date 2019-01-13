package firebasecli

import (
	"context"

	"google.golang.org/api/iterator"
)

// DbList lists up the all collections.
func (app *App) DbList(ctx context.Context) (collections []string, err error) {
	dbClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer dbClient.Close()

	for it := dbClient.Collections(ctx); ; {
		ref, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		collections = append(collections, ref.ID)
	}

	return collections, nil
}

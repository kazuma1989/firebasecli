package firebasecli

import (
	"context"
)

// DbList lists up the all collections.
func (app *App) DbList(ctx context.Context) (collections []string, err error) {
	dbClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer dbClient.Close()

	collectionRefs, err := dbClient.Collections(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	for _, ref := range collectionRefs {
		collections = append(collections, ref.ID)
	}

	return collections, nil
}

package firebasecli

import (
	"context"

	"google.golang.org/api/iterator"
)

// DbExport exports collections.
func (app *App) DbExport(ctx context.Context, collectionPaths []string) (collections map[string]map[string]map[string]interface{}, err error) {
	dbClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer dbClient.Close()

	collections = make(map[string]map[string]map[string]interface{})
	for _, key := range collectionPaths {
		allDocs := make(map[string]map[string]interface{})

		it := dbClient.Collection(key).Documents(ctx)
		for {
			docSnap, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}

			allDocs[docSnap.Ref.ID] = docSnap.Data()
		}

		collections[key] = allDocs
	}

	return collections, nil
}

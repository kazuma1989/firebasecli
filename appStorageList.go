package firebasecli

import (
	"context"

	"google.golang.org/api/iterator"
)

// StorageList lists up the all collections.
func (app *App) StorageList(ctx context.Context) (objects []string, err error) {
	storageClient, err := app.Storage(ctx)
	if err != nil {
		return nil, err
	}
	bucket, err := storageClient.DefaultBucket()
	if err != nil {
		return nil, err
	}

	it := bucket.Objects(ctx, nil)
	for {
		attr, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		objects = append(objects, attr.Name)
	}

	return objects, nil
}

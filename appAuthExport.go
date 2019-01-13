package firebasecli

import (
	"context"

	"firebase.google.com/go/auth"
	"google.golang.org/api/iterator"
)

// AuthExport exports the all users.
func (app *App) AuthExport(ctx context.Context) (users []*auth.ExportedUserRecord, err error) {
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	for it := authClient.Users(ctx, ""); ; {
		u, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

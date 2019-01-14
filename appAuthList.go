package firebasecli

import (
	"context"

	"firebase.google.com/go/auth"
)

// AuthList lists up the all users.
func (app *App) AuthList(ctx context.Context) (users []*auth.ExportedUserRecord, err error) {
	return app.AuthExport(ctx)
}

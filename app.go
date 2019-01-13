package firebasecli

import (
	"context"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// DefaultCredential is a path to a service account secret key
// which will be used when neither arguments nor environment variables are specified.
const DefaultCredential = "serviceAccountKey.json"

// App represents application.
type App struct {
	*firebase.App
}

// NewApp initializes an App.
func NewApp(ctx context.Context, credentialPath string) (*App, error) {
	app, err := func() (*firebase.App, error) {
		if credentialPath != "" {
			opt := option.WithCredentialsFile(credentialPath)
			return firebase.NewApp(ctx, nil, opt)
		}

		// When env var GOOGLE_APPLICATION_CREDENTIALS is specified.
		// e.g.) export GOOGLE_APPLICATION_CREDENTIALS=path/to/key.json
		app, err := firebase.NewApp(ctx, nil)
		if err == nil {
			return app, nil
		}

		exePath, err := os.Executable()
		if err != nil {
			return nil, err
		}
		defaultCredential := filepath.Join(filepath.Dir(exePath), DefaultCredential)

		opt := option.WithCredentialsFile(defaultCredential)
		return firebase.NewApp(ctx, nil, opt)
	}()

	return &App{App: app}, err
}

package firebasecli

import (
	"context"
	"fmt"
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
// If it failed to connect Firebase (e.g. invalid credential), it returns an error.
func NewApp(ctx context.Context, credentialPath string) (*App, error) {
	app, err := newFirebaseApp(ctx, credentialPath)
	if err != nil {
		err = fmt.Errorf("failed to connect Firebase: %v", err)
	}
	return &App{app}, err
}

// newFirebaseApp initializes Firebase app with given credentials.
// First, the argument will be taken, then an environmental variable will be taken when the first is not given.
// Otherwise DefaultCredential will be taken.
func newFirebaseApp(ctx context.Context, credentialPath string) (*firebase.App, error) {
	if credentialPath != "" {
		opt := option.WithCredentialsFile(credentialPath)
		return firebase.NewApp(ctx, nil, opt)
	}

	// In case of env var GOOGLE_APPLICATION_CREDENTIALS specified.
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
}

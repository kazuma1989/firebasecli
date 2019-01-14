package firebasecli

import (
	"context"
	"encoding/base64"

	"firebase.google.com/go/auth"
	"firebase.google.com/go/auth/hash"
)

// AuthImport imports users.
func (app *App) AuthImport(ctx context.Context, users []*auth.ExportedUserRecord, key, saltSeparator string, rounds, memoryCost int) (*auth.UserImportResult, error) {
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	bKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	bSaltSeparator, err := base64.StdEncoding.DecodeString(saltSeparator)
	if err != nil {
		return nil, err
	}
	h := hash.Scrypt{
		Key:           bKey,
		SaltSeparator: bSaltSeparator,
		Rounds:        rounds,
		MemoryCost:    memoryCost,
	}

	var usersToImport []*auth.UserToImport
	for _, ur := range users {
		u := &auth.UserToImport{}

		u.CustomClaims(ur.CustomClaims)
		u.Disabled(ur.Disabled)
		u.DisplayName(ur.DisplayName)
		u.Email(ur.Email)
		u.EmailVerified(ur.EmailVerified)
		u.Metadata(ur.UserMetadata)
		u.PasswordHash([]byte(ur.PasswordHash))
		u.PasswordSalt([]byte(ur.PasswordSalt))
		u.PhoneNumber(ur.PhoneNumber)
		u.PhotoURL(ur.PhotoURL)
		u.UID(ur.UID)

		var providers []*auth.UserProvider
		for _, p := range ur.ProviderUserInfo {
			providers = append(providers, &auth.UserProvider{
				UID:         p.UID,
				ProviderID:  p.ProviderID,
				Email:       p.Email,
				DisplayName: p.DisplayName,
				PhotoURL:    p.PhotoURL,
			})
		}
		u.ProviderData(providers)

		usersToImport = append(usersToImport, u)
	}

	return authClient.ImportUsers(ctx, usersToImport, auth.WithHash(h))
}

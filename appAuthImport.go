package firebasecli

import (
	"context"
	"encoding/base64"
	"regexp"
	"strconv"

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

// ParseHashConfig parses hash config following the format below:
// hash_config {
//   algorithm: SCRYPT,
//   base64_signer_key: <base64string>,
//   base64_salt_separator: <base64string>,
//   rounds: <integer>,
//   mem_cost: <integer>,
// }
func ParseHashConfig(hashConfig string) (hashKey, saltSeparator string, rounds, memCost int, err error) {
	re := regexp.MustCompile(`hash_config\s*{\s*
\s*algorithm\s*:\s*SCRYPT\s*,\s*
\s*base64_signer_key\s*:\s*([a-zA-Z0-9/+]+=*)\s*,\s*
\s*base64_salt_separator\s*:\s*([a-zA-Z0-9/+]+=*)\s*,\s*
\s*rounds\s*:\s*([1-9][0-9]*)\s*,\s*
\s*mem_cost\s*:\s*([1-9][0-9]*)\s*,\s*
\s*}`)
	values := re.FindStringSubmatch(hashConfig)
	if len(values) <= 4 {
		err = ErrFailedToParseHashConfig
		return
	}

	hashKey = values[1]
	saltSeparator = values[2]
	rounds, err = strconv.Atoi(values[3])
	if err != nil {
		err = ErrFailedToParseHashConfig
		return
	}
	memCost, err = strconv.Atoi(values[4])
	if err != nil {
		err = ErrFailedToParseHashConfig
		return
	}

	return
}

package firebasecli

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"strings"

	"firebase.google.com/go/auth"
	docopt "github.com/docopt/docopt-go"
)

// Auth manipulates Firebase Authentication.
func (c *Cmd) Auth(args ...string) error {
	opts, err := docopt.ParseArgs(`
Manipulate Firebase Authentication.

Usage:
  admin-tool auth list [--uid]
  admin-tool auth export [-f]
  admin-tool auth import ACCOUNT_FILE (-h FILE | -k KEY -s SEP -r N -m N)

Options:
  list   List the all users.
  --uid  Show UIDs only.

Options:
  export        Export users as JSON to stdout.
                Use redirect (>) to save results in a file.
  -f, --format  Output as formatted JSON.

Options:
  import ACCOUNT_FILE       Import users as JSON from an account file.
                            Password hash algorithm is SCRYPT modified by Firebase.
  -h, --hash-config FILE    Hash config file containing a content like:
                            hash_config {
                              algorithm: SCRYPT,
                              base64_signer_key: <base64string>,
                              base64_salt_separator: <base64string>,
                              rounds: <integer>,
                              mem_cost: <integer>,
                            }
  -k, --hash-key KEY        Hash key.
  -s, --salt-separator SEP  Salt separator.
  -r, --rounds N            Rounds.
  -m, --mem-cost N          Memory costs.
`, args, "")
	if err != nil {
		// TODO add an error explanation.
		return err
	}

	var arg struct {
		Auth bool

		List bool
		UID  bool `docopt:"--uid"`

		Export bool
		Format bool

		Import        bool
		AccountFile   string `docopt:"ACCOUNT_FILE"`
		HashConfig    string
		HashKey       string
		SaltSeparator string
		Rounds        int
		MemCost       int
	}
	if err := opts.Bind(&arg); err != nil {
		// TODO add an error explanation.
		return err
	}

	switch {
	case arg.List:
		return c.authList(arg.UID)

	case arg.Export:
		return c.authExport(arg.Format)

	case arg.Import:
		return c.authImport(arg.AccountFile, arg.HashConfig, arg.HashKey, arg.SaltSeparator, arg.Rounds, arg.MemCost)

	default:
		return ErrUnknownCommand
	}
}

func (c *Cmd) authList(uid bool) error {
	ctx := context.Background()

	userRecs, err := c.App.AuthList(ctx)
	if err != nil {
		return err
	}

	var users []string
	switch {
	case uid:
		for _, u := range userRecs {
			users = append(users, u.UID)
		}
	default:
		for _, u := range userRecs {
			users = append(users, u.UID+"\t"+u.Email)
		}
	}

	c.Println(strings.Join(users, "\n"))
	return nil
}

func (c *Cmd) authExport(format bool) error {
	ctx := context.Background()

	users, err := c.App.AuthExport(ctx)
	if err != nil {
		return err
	}

	var jsonData []byte
	if format {
		jsonData, err = json.MarshalIndent(users, "", "  ")
	} else {
		jsonData, err = json.Marshal(users)
	}
	if err != nil {
		// TODO add an error explanation.
		return err
	}

	c.Println(string(jsonData))
	return nil
}

func (c *Cmd) authImport(accountFile, hashConfig, hashKey, saltSeparator string, rounds, memCost int) error {
	if hashConfig != "" {
		config, err := ioutil.ReadFile(hashConfig)
		if err != nil {
			// TODO add an error explanation.
			return err
		}

		hashKey, saltSeparator, rounds, memCost, err = ParseHashConfig(string(config))
		if err != nil {
			return err
		}
	}

	var users []*auth.ExportedUserRecord
	accountData, err := ioutil.ReadFile(accountFile)
	if err != nil {
		// TODO add an error explanation.
		return err
	}
	err = json.Unmarshal(accountData, &users)
	if err != nil {
		// TODO add an error explanation.
		return err
	}

	ctx := context.Background()
	result, err := c.App.AuthImport(ctx, users, hashKey, saltSeparator, rounds, memCost)
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	c.Eprintln(string(jsonData))
	return nil
}

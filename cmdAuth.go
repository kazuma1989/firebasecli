package firebasecli

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"firebase.google.com/go/auth"
	docopt "github.com/docopt/docopt-go"
)

// ErrFailedToParseHashConfig is returned when hash config file does not follow the format.
var ErrFailedToParseHashConfig = errors.New("failed to parse hash_config")

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
	ctx := context.Background()

	if hashConfig != "" {
		config, err := ioutil.ReadFile(hashConfig)
		if err != nil {
			// TODO add an error explanation.
			return err
		}

		re := regexp.MustCompile(`hash_config\s*{\s*
\s*algorithm\s*:\s*SCRYPT\s*,\s*
\s*base64_signer_key\s*:\s*([a-zA-Z0-9/+]+=*)\s*,\s*
\s*base64_salt_separator\s*:\s*([a-zA-Z0-9/+]+=*)\s*,\s*
\s*rounds\s*:\s*([1-9][0-9]*)\s*,\s*
\s*mem_cost\s*:\s*([1-9][0-9]*)\s*,\s*
\s*}`)
		values := re.FindStringSubmatch(string(config))
		if len(values) <= 4 {
			return ErrFailedToParseHashConfig
		}

		hashKey = values[1]
		saltSeparator = values[2]
		rounds, err = strconv.Atoi(values[3])
		if err != nil {
			// TODO add an error explanation.
			return err
		}
		memCost, err = strconv.Atoi(values[4])
		if err != nil {
			// TODO add an error explanation.
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

	result, err := c.App.AuthImport(ctx, users, hashKey, saltSeparator, rounds, memCost)
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	c.Eprintln(string(jsonData))
	return nil
}

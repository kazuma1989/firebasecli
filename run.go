package firebasecli

// Version of the firebasecli.
const Version = "0.0.1"

// DefaultCommands holds the default sub commands.
var DefaultCommands = make(Commands)

// Run starts the app.
var Run = DefaultCommands.Run

func init() {
	DefaultCommands["auth"] = func(c *Cmd, args ...string) error {
		return c.Auth(args...)
	}
	DefaultCommands["db"] = func(c *Cmd, args ...string) error {
		return c.Db(args...)
	}
}

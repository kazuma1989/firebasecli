package firebasecli

// DefaultCommands holds the default sub commands.
var DefaultCommands = NewCmd()

// Run starts the app.
var Run = DefaultCommands.Run

func init() {
	DefaultCommands.Sub["auth"] = RunnableFunc(DefaultCommands.Auth)
	DefaultCommands.Sub["db"] = RunnableFunc(DefaultCommands.Db)
}

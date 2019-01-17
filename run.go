package firebasecli

// DefaultCmd holds the default sub commands.
var DefaultCmd = NewCmd()

// Run starts the app.
var Run = DefaultCmd.Run

func init() {
	DefaultCmd.Sub["auth"] = RunnableFunc(DefaultCmd.Auth)
	DefaultCmd.Sub["db"] = RunnableFunc(DefaultCmd.Db)
}

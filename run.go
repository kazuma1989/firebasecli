package firebasecli

// Run starts the app.
func Run(args ...string) error {
	return DefaultCmd.Run(args...)
}

// DefaultCmd holds the default sub commands.
var DefaultCmd = NewCmd()

func init() {
	DefaultCmd.Sub["auth"] = RunnableFunc(DefaultCmd.Auth)
	DefaultCmd.Sub["db"] = RunnableFunc(DefaultCmd.Db)
}

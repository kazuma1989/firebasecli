package firebasecli

// Run starts the app.
func Run(args ...string) error {
	return DefaultCmd.Run(args...)
}

// DefaultCmd holds the default sub commands.
var DefaultCmd = NewCmd()

func init() {
	DefaultCmd.Sub[""] = RunnableFunc(func(args ...string) error {
		args = append([]string{"--help"}, args...)
		return DefaultCmd.Run(args...)
	})
	DefaultCmd.Sub["auth"] = RunnableFunc(DefaultCmd.Auth)
	DefaultCmd.Sub["db"] = RunnableFunc(DefaultCmd.Db)
}

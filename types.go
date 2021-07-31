package ezcli

// Command Handler struct.
type CommandHandler struct {
	Commands []*Command
}

// Option (flag) struct.
type Option struct {
	Name, Description string
}

// Command Data struct (arguments, options...)
type CommandData struct {
	Arguments, Options []string
}

// Command struct.
type Command struct {
	Name, Description string
	Options           []*Option
	Usages            []string
	CommandData       *CommandData
	Handler           func()
}

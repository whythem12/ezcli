package ezcli

import (
	"log"
	"os"
	"strings"
)

// Add a new command to the handler.
func (ch *CommandHandler) AddCommand(c *Command) {
	// Add command name as usage if not added.
	if len(c.Usages) == 0 {
		c.Usages = append(c.Usages, c.Name)
	}

	ch.Commands = append(ch.Commands, c)
}

// Add more command to the handler.
func (ch *CommandHandler) AddCommands(cs []*Command) {
	for _, c := range cs {
		ch.AddCommand(c)
	}
}

// Set error function that will run when command not found.
func (ch *CommandHandler) SetNotFoundFunction(fn func()) {
	ch.CommandNotFoundFunc = fn
}

// Handle commands.
func (ch *CommandHandler) Handle() {
	args := os.Args[1:]

	// Check parameter length
	if len(args) == 0 {
		log.Fatal("You need to pass a command. Type 'help' for show all commands.")
	}

	commandData := CommandData{}

	// Get options and params
	for _, item := range args {
		if !strings.HasPrefix(item, "-") {
			commandData.Arguments = append(commandData.Arguments, item)
		}
	}

	// Check parameter length
	if len(commandData.Arguments) == 0 {
		log.Fatal("You need to pass a command. Type 'help' for show all commands.")
	}

	commandName := commandData.Arguments[0]
	commandData.Arguments = commandData.Arguments[1:]

	// Handle command
	err := ch.FindCommand(commandName, func(c *Command) error {
		// Add Options
		for _, item := range args {
			if strings.HasPrefix(item, "-") {
				optionName, optionValue := strings.TrimLeft(item, "-"), ""

				// If has a value
				if strings.Contains(optionName, "=") {
					splittedText := strings.SplitN(optionName, "=", 2)
					optionName = splittedText[0]
					optionValue = splittedText[1]
				}

				c.FindOption(optionName, func(o *CommandOption) {
					o.Value = optionValue
					commandData.Options = append(commandData.Options, o)
				})
			}
		}

		c.CommandData = &commandData
		c.Execute(c)

		return nil
	})

	if err != nil {
		ch.CommandNotFoundFunc()
		return
	}
}

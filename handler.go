package ezcli

import (
	"fmt"
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

// Handle commands.
func (ch *CommandHandler) Handle() {
	args := os.Args[1:]

	// Check parameter length
	if len(args) == 0 {
		log.Fatal("You need to pass a command. Type 'help' for show all commands.")
		os.Exit(1)
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
		os.Exit(1)
	}

	commandName := commandData.Arguments[0]
	commandData.Arguments = commandData.Arguments[1:]

	// Handle command
	ch.FindCommand(commandName, func(c *Command) error {
		// Add Options
		for _, item := range args {
			if strings.HasPrefix(item, "-") {
				optionName := strings.TrimPrefix(item, "-")

				c.FindOption(optionName, func(o *CommandOption) {
					commandData.Options = append(commandData.Options, o)
				})
			}
		}

		c.CommandData = &commandData
		c.Execute(c)

		return nil
	})
}

// Find a command from handler.
func (ch *CommandHandler) FindCommand(name string, fn func(c *Command) error) error {
	for _, item := range ch.Commands {
		if strings.EqualFold(item.Name, name) {
			return fn(item)
		}
	}

	return fmt.Errorf("Command not found! Please check your parameter")
}

// Find an option from command.
func (c *Command) FindOption(name string, fn func(o *CommandOption)) {
	for _, item := range c.Options {
		if strings.EqualFold(item.Name, name) {
			fn(item)
		}
	}
}

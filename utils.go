package ezcli

import (
	"fmt"
	"log"
	"strings"
)

// Create a command handler with built-in help command.
func CreateHandler() *CommandHandler {
	var handler CommandHandler

	handler.SetNotFoundFunction(func() {
		log.Fatal("Command not found! Please run command 'help' for list all commands.")
	})

	handler.AddCommand(&Command{
		Name:        "help",
		Description: "Built-in help command for application.",
		Usages:      []string{"help", "help <command name>"},
		Execute: func(c *Command) {
			data := c.CommandData

			if len(data.Arguments) == 0 {
				fmt.Println("List of all commands. For more information, add a command name parameter to command.")
				for _, command := range handler.Commands {
					fmt.Printf("  %s | %s\n", command.Name, command.Description)
				}
			} else {
				commandName := data.Arguments[0]

				err := handler.FindCommand(commandName, func(c *Command) error {
					fmt.Printf("Command %s:\n  Description: %s\n  Usages:\n    %s\n", c.Name, c.Description, strings.Join(c.Usages, "\n    "))

					// Add options
					if len(c.Options) > 0 {
						fmt.Println("  Options:")

						for _, item := range c.Options {
							fmt.Printf("    %s | %s\n", item.Name, item.Description)
						}
					}

					return nil
				})

				// Command not found
				if err != nil {
					log.Fatal(err)
				}
			}
		},
	})

	return &handler
}

// Find a command from handler.
func (ch *CommandHandler) FindCommand(name string, fn func(c *Command) error) error {
	for _, item := range ch.Commands {
		if strings.EqualFold(item.Name, name) || stringSliceContains(item.Aliases, name) {
			return fn(item)
		}
	}

	return fmt.Errorf("Command not found! Please check your parameter")
}

// Find an option template from command.
func (c *Command) FindOption(name string, fn func(o *CommandOption)) {
	for _, item := range c.Options {
		if strings.EqualFold(item.Name, name) || stringSliceContains(item.Aliases, name) {
			fn(item)
		}
	}
}

// Find an option from command data.
func (c *CommandData) FindOption(name string, fn func(o *CommandOption)) {
	for _, item := range c.Options {
		if strings.EqualFold(item.Name, name) || stringSliceContains(item.Aliases, name) {
			fn(item)
		}
	}
}

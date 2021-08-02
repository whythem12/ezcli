package ezcli

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Create a command handler with built-in help command.
func CreateHandler() *CommandHandler {
	var handler CommandHandler

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
					os.Exit(1)
				}
			}
		},
	})

	return &handler
}

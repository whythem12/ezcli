# Sub Commands

Sub-command is a mini command inside your command.

## Informations

- Sub-commands takes the options from main command.
- Sub-commands doesn't supports aliasing.

## Simple Example

We have an app with this code:

```go
package main

import (
    "fmt"

    "github.com/5elenay/ezcli"
)

func main() {
    handler := ezcli.NewApp("test") // Create handle function also gives built-in help command. So you dont need to write a help command yourself.

    // Adding new command(s).
    handler.AddCommand(&ezcli.Command{
        Name: "test", // Command name.
        SubCommands: []*ezcli.SubCommand{ // Subcommands
            {
                Name:   "sub",                // Sub-command name.
                Usages: []string{"test sub"}, // Sub-command usage.
                Execute: func(sc *ezcli.SubCommand) { // The sub-command function will run.
                    fmt.Println("Sub-Command Function.")
                },
            },
        },
        Execute: func(c *ezcli.Command) { // The function will run.
            fmt.Println("Main Function.")
        },
    })

    handler.Handle() // Handle commands.
}
```

Now lets compile the app with `go build .`.

Now lets run the app with `./<name> test`. You should get the following result:

```
Main Function.
```

Our main commands works perfect. Lets test the sub-command with `./<name> test sub`. You should get the following result:

````
Sub-Command Function.
```

And our sub-command also works perfect.
````

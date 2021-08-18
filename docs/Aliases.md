## Aliases

You can give aliases for both commands and options easily!

## Command Aliases

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
        Name:        "hi",              // Command name.
        Description: "Say hi world!",   // Command description.
        Aliases:     []string{"hello"}, // Command aliases.
        Execute: func(c *ezcli.Command) { // The function will run.
            fmt.Println("Hello!")
        },
    })

    handler.Handle() // Handle commands.
}
```

Now lets compile the app with `go build .`.

Time to test the command. Lets run the app with `./<name> hi`. You should get the following result:

```
Hello!
```

It works, but how we will use aliases? Lets just use `hello` instead of `hi`. Run the app with `./<name> hello` and you should get the following result:

```
Hello!
```

Our alias is also works!

## Option Aliases

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
        Name:        "hi",            // Command name.
        Description: "Say hi world!", // Command description.
        Options: []*ezcli.CommandOption{
            { // New option
                Name:        "test",
                Description: "Test option.",
                Aliases:     []string{"t"}, // Aliases for this option.
            },
        },
        Execute: func(c *ezcli.Command) { // The function will run.
            fmt.Println("Hello!")

            for _, i := range c.CommandData.Options { // List all of the Options with values.
                fmt.Printf("%s - %s", i.Name, i.Value)
            }
        },
    })

    handler.Handle() // Handle commands.
}
```

Now lets compile the app with `go build .`.

Time to test the aliases. Run the app with `./<name> hi -test=123`. You should get the following result:

```
Hello!
test - 123
```

It works, lets test the alias with `./<name> hi -t=456`. You should get the following result:

```
Hello!
test - 456
```

Our alias is also works!

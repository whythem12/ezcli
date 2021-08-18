# Ezcli

Create a CLI tool in <10 second without any performance decrease.

[![Size](https://img.shields.io/github/languages/code-size/5elenay/ezcli)]()
[![License](https://img.shields.io/github/license/5elenay/ezcli)]()
[![Stars](https://img.shields.io/github/stars/5elenay/ezcli)]()
[![Release](https://img.shields.io/github/v/release/5elenay/ezcli)]()

## What is Ezcli?

Ezcli is a Go package for building CLI tools easily.

## Installation

- Initialize your project with `go mod init <name>`
- Get Ezcli with `go get github.com/5elenay/ezcli`

## API Reference

Click [here](https://pkg.go.dev/github.com/5elenay/ezcli) for API reference. (pkg.go.dev sometime does not update package version, so use `https://pkg.go.dev/github.com/5elenay/ezcli@<latest version>` for get the API reference for latest version.)

## Documentations

- [Options / Flags](https://github.com/5elenay/ezcli/blob/main/docs/Options.md)
- [Arguments](https://github.com/5elenay/ezcli/blob/main/docs/Arguments.md)
- [Aliases](https://github.com/5elenay/ezcli/blob/main/docs/Aliases.md)
- [Questions](https://github.com/5elenay/ezcli/blob/main/docs/Questions.md)
- [Sub-Commands](https://github.com/5elenay/ezcli/blob/main/docs/SubCommands.md)

## Example

Simple example for Ezcli.

```go
package main

import (
    "fmt"

    "github.com/5elenay/ezcli"
)

func main() {
    handler := ezcli.NewApp("test") // Create handle function also gives built-in help command. So you don't need to write a help command yourself.

    // Adding a new command.
    handler.AddCommand(&ezcli.Command{
        Name:        "hello", // Command name.
        Description: "Say hello world!", // Command description.
        Options: []*ezcli.CommandOption{ // Command options (example: -force, -confirm etc...).
            {
                Name:        "test", // Option name.
                Description: "A test command", // Option description.
            },
        },
        Execute: func(c *ezcli.Command) { // The function will run.
            fmt.Println("Hello Command!")
        },
    })

    handler.Handle() // Handle commands.
}

```

Now lets compile our app with `go build .`

Time to test our app! Run the compiled app with `./<name> help`

You should get the following result:

```
List of all commands. For more information, add a command name parameter to command.
  help | Built-in help command for application.
  hello | Say hello world!
```

Lets test our `hello` command with `./<name> hello`:

```
Hello Command!
```

Congrats, you created your first app with Ezcli!

## License

This project is licensed under [MIT](https://opensource.org/licenses/MIT) license.

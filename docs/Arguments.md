# Arguments

If your app will use arguments, You can learn some basics here.

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
        Name:        "hi",            // Command name.
        Description: "Say hi world!", // Command description.
        Execute: func(c *ezcli.Command) { // The function will run.
            fmt.Println("Hello!")

            for i, v := range c.CommandData.Arguments { // List all of the Options with values.
                fmt.Println(i, v)
            }
        },
    })

    handler.Handle() // Handle commands.
}
```

Now lets compile the app with `go build .`.

Time to test the arguments. Lets run the app with `./<name> hi`. You should get the following result:

```
Hello!
```

Lets add an argument for our app. Run the app with `./<name> hi example`. You should get the following result:

```
Hello!
0 example
```

It works! what if we add more argument? Lets run the app with `./<name> hi example im bored.`. You should get the following result:

```
Hello!
0 example
1 im
2 bored.
```

It works, Since `<Command>.<CommandData>.Arguments` is a string slice, its really easy to use.

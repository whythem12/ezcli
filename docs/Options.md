# Options (Flags)

Options or flags are useful when you need to configure your command.

## Simple Example

We have an app with this code:

```go
package main

import (
    "fmt"

    "github.com/5elenay/ezcli"
)

func main() {
    handler := ezcli.CreateHandler() // Create handle function also gives built-in help command. So you dont need to write a help command yourself.

    // Adding new command(s).
    handler.AddCommand(&ezcli.Command{
        Name:        "hi",            // Command name.
        Description: "Say hi world!", // Command description.
        Options: []*ezcli.CommandOption{
            { // New option
                Name:        "test",
                Description: "Test option.",
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

Time to test the options, First run the command without any option with `./<name> hi`. You should get the following result:

```
Hello!
```

It works, Lets add the `test` option. For this, We will add `-` prefix. (The number of the `-` prefix doesn't matter.) Lets run with `./<name> hi -test`. You should get the following result:

```
Hello!
test -
```

It works, but how we will add a value to this option? We must use `=` for add a value to the option. Lets run with `./<name> hi -test=1337`. You should get the following result:

```
Hello!
test - 1337
```

Perfect, Now we also can get the option's value. But what about strings? What if we need to get something like `an example value`? For this, lets run with `./<name> hi -test="hello everyone!"`. You should get the following result:

```
Hello!
test - hello everyone!
```

Great! Now lets check the some tips about options / flags.

## Tips

- You can use infinite `-` for prefix. For example: `./<name> hi --test=hi` and `./<name> hi ---------test=hi` will also work.
- You can put options / flags anywhere. For example: `./<name> hi -test=1` and `./<name> -test=1 hi` is same.

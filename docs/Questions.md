# Questions

Ezcli now provides you to handle input easy and secure way.

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
        Name: "name", // Command name.
        Execute: func(c *ezcli.Command) { // The function will run.
            something := ezcli.Question{ // Ask a question.
                Input: "What is your name: ", // Input string.
            }

            something.Ask()               // Ask and wait for an answer.
            fmt.Println(something.Answer) // Get the answer.
        },
    })

    handler.Handle() // Handle commands.
}
```

Now lets compile the app with `go build .`.

Now time to test our app. Lets run with `./<name> name`. You should get the following result:

```
What is your name:
```

Congrats! Lets write a name, for example `John Doe`:

```
What is your name: John Doe
```

And press enter. Lets see the result:

```
What is your name: John Doe
John Doe
```

And we got the answer without any error.

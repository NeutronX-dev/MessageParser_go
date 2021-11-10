# MessageParser_go
Message Parser for Go

# Installation
```
go get github.com/NeutronX-dev/MessageParser_go
```

# Importing
```go
import (
    "..."

    "github.com/NeutronX-dev/MessageParser_go"
)
```

# Example
```go
var ToParse string = "!help moderation"
message := MessageParser.ParseMessage("!", ToParse)
if message.IsParsed {
	switch message.Command {
	case "help":
		if message.HasArguments(1) {
			fmt.Println(message.Arguments)
			fmt.Printf("Need help with \"%v\"?\n", strings.Join(message.Arguments, " "))
		} else {
			fmt.Println("Need at least one Argument")
		}
	}
}
```
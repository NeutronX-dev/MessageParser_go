package main

import (
	"fmt"
	"strings"
)

type ParsedMessage struct {
	Command   string
	Arguments []string
	IsParsed  bool
}

func (Message *ParsedMessage) HasArguments(amount int) bool {
	return len(Message.Arguments)+1 >= amount
}
func (Message *ParsedMessage) ArgumnetsContain(value string) bool {
	for i := 0; i < len(Message.Arguments); i++ {
		if strings.ToLower(Message.Arguments[i]) == strings.ToLower(value) {
			return true
		}
	}
	return false
}

func ParseMessage(prefix, message string) ParsedMessage {
	if strings.HasPrefix(message, prefix) {
		message = message[len(prefix):]
		var All []string = strings.Split(message, " ")
		fmt.Println(All)
		if len(All) >= 1 {
			return ParsedMessage{Command: All[0], Arguments: All[1:], IsParsed: true}
		}
		return ParsedMessage{Command: All[0], Arguments: []string{}, IsParsed: true}
	} else {
		return ParsedMessage{}
	}
}
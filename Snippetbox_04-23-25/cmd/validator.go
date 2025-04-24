/* validator.go */

package main

import (
	"strings"
)

type Message struct {
	Title   string
	Content string
	Errors  map[string]string
}

func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Title) == "" {
		msg.Errors["Title"] = "Please enter a title."
	}

	if strings.TrimSpace(msg.Content) == "" {
		msg.Errors["Content"] = "Please enter content."
	}

	return len(msg.Errors) == 0
}

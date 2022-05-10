package main

import "github.com/andersfylling/disgord"

type Command struct {
	Name        string
	CustomID    string
	Description string
	Options     []*disgord.ApplicationCommandOption
	Components  []*CommandMessageComponent
	Type        int
	Handler     ExecutionHandler
}

type CommandMessageComponent struct {
	Name        string
	CustomID    string
	Type        disgord.MessageComponentType
	Style       disgord.ButtonStyle
	Label       string
	Url         string
	Disabled    bool
	Options     []*disgord.SelectMenuOption
	Placeholder string
	MinValues   int
	MaxValues   int
	Handler     ExecutionHandlerComponent
}

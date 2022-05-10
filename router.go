package main

import (
	"github.com/andersfylling/disgord"
)

type Router struct {
	Commands             []*Command
	CommandComponents    []*CommandMessageComponent
	CommandComponentsMap map[string]*Command
	Client               *disgord.Client
}

func Create(client *disgord.Client) *Router {
	return &Router{
		Commands:             []*Command{},
		Client:               client,
		CommandComponentsMap: make(map[string]*Command),
	}
}

func (r *Router) GetCmd(name string) *Command {
	for _, cmd := range r.Commands {
		if cmd.Name == name {
			return cmd
		}
	}

	return nil
}

func (r *Router) InitializeCommands() disgord.HandlerReady {
	return func(s disgord.Session, h *disgord.Ready) {
		user, _ := r.Client.Cache().GetCurrentUser()

		for i := range commands {
			if err := r.Client.ApplicationCommand(user.ID).Global().Create(&disgord.CreateApplicationCommand{
				Type:        disgord.ApplicationCommandType(commands[i].Type),
				Name:        commands[i].Name,
				Description: commands[i].Description,
				Options:     commands[i].Options,
			}); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (r *Router) RegisterCommand(command *Command) {
	r.Commands = append(r.Commands, command)

	for _, cmdComponent := range command.Components {
		r.CommandComponents = append(r.CommandComponents, cmdComponent)
	}
}

func (r *Router) RegisterCMDList(commands []*Command) {
	r.Commands = append(r.Commands, commands...)

	for _, c := range commands {
		for _, cmdComponent := range c.Components {
			r.CommandComponents = append(r.CommandComponents, cmdComponent)
		}
	}
}

func (r *Router) Init() {
	r.Client.Gateway().Ready(r.InitializeCommands())
	r.Client.Gateway().InteractionCreate(r.Handler())
}

func (r *Router) Handler() disgord.HandlerInteractionCreate {
	return func(s disgord.Session, h *disgord.InteractionCreate) {
		ctx := &Ctx{
			Client:      r.Client,
			Session:     &s,
			Interaction: h,
			Router:      r,
		}
		for _, cmd := range r.Commands {
			if h.Data.Name != "" && h.Data.Name == cmd.CustomID {
				ctx.Command = cmd

				cmd.Handler(ctx)
				return
			}
		}

		for _, component := range r.CommandComponents {
			if h.Data.CustomID != "" && h.Data.CustomID == component.CustomID {
				component.Handler(&ComponentCtx{
					Client:      r.Client,
					Session:     &s,
					Interaction: h,
					Router:      r,
				})
				return
			}
		}
	}
}

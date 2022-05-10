package main

import (
	"context"
	"github.com/andersfylling/disgord"
)

type ComponentCtx struct {
	Client      *disgord.Client
	Session     *disgord.Session
	Interaction *disgord.InteractionCreate
	Router      *Router
}

type ExecutionHandlerComponent func(ctx *ComponentCtx)

func (ctx *ComponentCtx) ResponseText(text string) error {
	err := ctx.Interaction.Reply(context.Background(), *ctx.Session, &disgord.CreateInteractionResponse{
		Type: 4,
		Data: &disgord.CreateInteractionResponseData{
			Content: text,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (ctx *ComponentCtx) ResponseEmbed(embed *disgord.Embed) error {
	err := ctx.Interaction.Reply(context.Background(), *ctx.Session, &disgord.CreateInteractionResponse{
		Type: 4,
		Data: &disgord.CreateInteractionResponseData{
			Embeds: []*disgord.Embed{embed},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (ctx *ComponentCtx) ResponseTextEmbed(text string, embed *disgord.Embed) error {
	err := ctx.Interaction.Reply(context.Background(), *ctx.Session, &disgord.CreateInteractionResponse{
		Type: 4,
		Data: &disgord.CreateInteractionResponseData{
			Embeds:  []*disgord.Embed{embed},
			Content: text,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

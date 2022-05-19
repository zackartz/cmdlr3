package cmdlr3

import (
	"context"
	"github.com/andersfylling/disgord"
)

type Ctx struct {
	Client      *disgord.Client
	Session     *disgord.Session
	Interaction *disgord.InteractionCreate
	Command     *Command
	Router      *Router
}

type ExecutionHandler func(ctx *Ctx)

func (ctx *Ctx) ResponseText(text string) error {
	var components []*disgord.MessageComponent

	if GetComponentsFromCommand(ctx.Command) != nil {
		components = []*disgord.MessageComponent{
			{
				Type:       1,
				Components: GetComponentsFromCommand(ctx.Command),
			},
		}
	}

	err := ctx.Interaction.Reply(context.Background(), *ctx.Session, &disgord.CreateInteractionResponse{
		Type: 4,
		Data: &disgord.CreateInteractionResponseData{
			Content:    text,
			Components: components,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (ctx *Ctx) ResponseEmbed(embed *disgord.Embed) error {
	var components []*disgord.MessageComponent

	if GetComponentsFromCommand(ctx.Command) != nil {
		components = []*disgord.MessageComponent{
			{
				Type:       1,
				Components: GetComponentsFromCommand(ctx.Command),
			},
		}
	}

	err := ctx.Interaction.Reply(context.Background(), *ctx.Session, &disgord.CreateInteractionResponse{
		Type: 4,
		Data: &disgord.CreateInteractionResponseData{
			Embeds:     []*disgord.Embed{embed},
			Components: components,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (ctx *Ctx) ResponseTextEmbed(text string, embed *disgord.Embed) error {
	var components []*disgord.MessageComponent

	if GetComponentsFromCommand(ctx.Command) != nil {
		components = []*disgord.MessageComponent{
			{
				Type:       1,
				Components: GetComponentsFromCommand(ctx.Command),
			},
		}
	}

	err := ctx.Interaction.Reply(context.Background(), *ctx.Session, &disgord.CreateInteractionResponse{
		Type: 4,
		Data: &disgord.CreateInteractionResponseData{
			Embeds:     []*disgord.Embed{embed},
			Content:    text,
			Components: components,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

type SubCommandCtx struct {
	Client      *disgord.Client
	Session     *disgord.Session
	Interaction *disgord.InteractionCreate
	Command     *SubCommand
	Router      *Router
}

type SubCommandExecutionHandler func(ctx *SubCommandCtx)

func (ctx *SubCommandCtx) ResponseText(text string) error {
	var components []*disgord.MessageComponent

	if GetComponentsFromSubCommand(ctx.Command) != nil {
		components = []*disgord.MessageComponent{
			{
				Type:       1,
				Components: GetComponentsFromSubCommand(ctx.Command),
			},
		}
	}

	err := ctx.Interaction.Reply(context.Background(), *ctx.Session, &disgord.CreateInteractionResponse{
		Type: 4,
		Data: &disgord.CreateInteractionResponseData{
			Content:    text,
			Components: components,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (ctx *SubCommandCtx) ResponseEmbed(embed *disgord.Embed) error {
	var components []*disgord.MessageComponent

	if GetComponentsFromSubCommand(ctx.Command) != nil {
		components = []*disgord.MessageComponent{
			{
				Type:       1,
				Components: GetComponentsFromSubCommand(ctx.Command),
			},
		}
	}

	err := ctx.Interaction.Reply(context.Background(), *ctx.Session, &disgord.CreateInteractionResponse{
		Type: 4,
		Data: &disgord.CreateInteractionResponseData{
			Embeds:     []*disgord.Embed{embed},
			Components: components,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (ctx *SubCommandCtx) ResponseTextEmbed(text string, embed *disgord.Embed) error {
	var components []*disgord.MessageComponent

	if GetComponentsFromSubCommand(ctx.Command) != nil {
		components = []*disgord.MessageComponent{
			{
				Type:       1,
				Components: GetComponentsFromSubCommand(ctx.Command),
			},
		}
	}

	err := ctx.Interaction.Reply(context.Background(), *ctx.Session, &disgord.CreateInteractionResponse{
		Type: 4,
		Data: &disgord.CreateInteractionResponseData{
			Embeds:     []*disgord.Embed{embed},
			Content:    text,
			Components: components,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

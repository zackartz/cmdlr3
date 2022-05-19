package main

import (
	"github.com/andersfylling/disgord"
	"github.com/sirupsen/logrus"
	"github.com/zackartz/cmdlr3"
	"os"
)

var (
	token = os.Getenv("TOKEN")
	log   = logrus.New()

	commands = []*cmdlr3.Command{{
		Name:        "testing_command",
		CustomID:    "testing_command",
		Description: "This is a command for testing",
		SubCommands: []*cmdlr3.SubCommand{
			{
				Name:        "hi",
				Description: "says hi",
				SubCommands: []*cmdlr3.SubCommand{
					{
						Name:        "hi2",
						Description: "Really says hi",
						Handler: func(ctx *cmdlr3.SubCommandCtx) {
							ctx.ResponseText("hi2")
						},
					},
				},
			},
			{
				Name:        "not_hi",
				Description: "Does not say hi",
				Handler: func(ctx *cmdlr3.SubCommandCtx) {
					ctx.ResponseText("not hi")
				},
			},
		},
		Handler: func(ctx *cmdlr3.Ctx) {
			err := ctx.ResponseText("Do this work?")
			if err != nil {
				log.Printf("%v", err)
			}
		},
	}}
)

func main() {
	client := disgord.New(disgord.Config{
		Intents:  disgord.AllIntents(),
		BotToken: token,
		Logger:   log,
	})
	defer func(gateway disgord.GatewayQueryBuilder) {
		err := gateway.StayConnectedUntilInterrupted()
		if err != nil {
			panic(err)
		}
	}(client.Gateway())

	router := cmdlr3.Create(client)
	router.RegisterCMDList(commands)
	router.Init()
}

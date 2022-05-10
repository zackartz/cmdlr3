package cmdlr3

// EXAMPLE MAIN.GO

//
//import (
//	"fmt"
//	"github.com/andersfylling/disgord"
//	"github.com/sirupsen/logrus"
//	"os"
//	"strings"
//)
//
//var (
//	token = os.Getenv("TOKEN")
//	log   = logrus.New()
//
//	commands = []*Command{{
//		Name:        "testing_command",
//		CustomID:    "testing_command",
//		Description: "This is a command for testing",
//		Options:     nil,
//		Components: []*CommandMessageComponent{
//			//{
//			//	Type:     2,
//			//	CustomID: "button_thingy",
//			//	Label:    "Click me!",
//			//	Style:    1,
//			//	Handler: func(ctx *ComponentCtx) {
//			//		ctx.ResponseText("You clicked the button!")
//			//	},
//			//},
//			{
//				Type:     3,
//				Label:    "Pick a value",
//				CustomID: "select_thingy",
//				Options: []*disgord.SelectMenuOption{{
//					Label:       "Option A",
//					Value:       "valueA",
//					Description: "Description for Option A",
//				}, {
//					Label:       "Option B",
//					Value:       "valueB",
//					Description: "Description for Option B",
//				}},
//				MinValues: 1,
//				MaxValues: 2,
//				Handler: func(ctx *ComponentCtx) {
//					ctx.ResponseText(fmt.Sprintf("You picked `%s`", strings.Join(ctx.Interaction.Data.Values, ", ")))
//				},
//			},
//		},
//		Handler: func(ctx *Ctx) {
//			err := ctx.ResponseText("Do this work?")
//			if err != nil {
//				log.Printf("%v", err)
//			}
//		},
//	}}
//)
//
//func main() {
//	client := disgord.New(disgord.Config{
//		Intents:  disgord.AllIntents(),
//		BotToken: token,
//		Logger:   log,
//	})
//	defer func(gateway disgord.GatewayQueryBuilder) {
//		err := gateway.StayConnectedUntilInterrupted()
//		if err != nil {
//			panic(err)
//		}
//	}(client.Gateway())
//
//	router := Create(client)
//	router.RegisterCMDList(commands)
//	router.Init()
//}

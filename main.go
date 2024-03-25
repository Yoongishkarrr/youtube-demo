package main 

import (
    "context"
    "fmt"
    "log"
    "os"
    "encoding/json"

    "github.com/joho/godotenv"
    "github.com/shomali11/slacker"
	witai "github.com/wit-ai/wit-go/v2"
	"github.com/tidwall/gjson"
	"github.com/Krognol/go-wolfram"
)

var wolframClient *wolfram.Client 

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
    for event := range analyticsChannel {
        fmt.Println("Command Events")
        fmt.Println(event.Timestamp)
        fmt.Println(event.Command)
        fmt.Println(event.Parameters)
        fmt.Println(event.Event)
        fmt.Println()
    }
}

func main() {
    godotenv.Load(".env")
    
    bot := slacker.NewClient(os.Getenv("SLACK_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
    client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	wolframClient := &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}
    go printCommandEvents(bot.CommandEvents())
    
    bot.Command("query for bot - <message>", &slacker.CommandDefinition{
        Description: "send any question to wolfram",
        Handler: func(botctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
            query := request.Param("message")
            fmt.Println(query)
			msg, _ := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			data ,_ := json.MarshalIndent(msg, "", "    ")
			rough := string(data[:])
			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			answer := value.String()
			res, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.metric, 1000)
			if err != nil {
				fmt.Println("there is an  error with the API call")
			}
			fmt.Println(value)
            response.Reply(res)
        },
    })

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    err := bot.Listen(ctx)

    if err != nil {
        log.Fatalf("%v\n", err)
    }
}

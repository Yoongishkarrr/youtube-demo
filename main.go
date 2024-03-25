package main 

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/shomali11/slacker"
)

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
    
    go printCommandEvents(bot.CommandEvents())
    
    bot.Command("query for bot - <message>", &slacker.CommandDefinition{
        Description: "send any question to wolfram",
        Handler: func(botctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
            query := request.Param("message")
            fmt.Println(query)
            /*client.Parse*/
            response.Reply("received")
        },
    })

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    err := bot.Listen(ctx)

    if err != nil {
        log.Fatalf("%v\n", err)
    }
}

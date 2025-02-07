package app

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/urfave/cli"
)

func LangchainAgent(c *cli.Context)  {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	apiKey := viper.GetString("OPENAI_API_KEY")
	
	fmt.Println("====== LANGCHAIN AGENT/OPENAI ======")

	llm, err := openai.New(
		openai.WithModel("gpt-4o-mini"),
		openai.WithToken(apiKey),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	role := c.String("role")
	question := c.String("prompt")

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, role),
		llms.TextParts(llms.ChatMessageTypeHuman, question),	
	}

	completion, err := llm.GenerateContent(
		ctx,
		content,
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
		llms.WithSeed(1),
		llms.WithFrequencyPenalty(0.6),
		llms.WithTemperature(0.4),
	)

	if err != nil {
		log.Fatal(err)
	}

	_ = completion

}
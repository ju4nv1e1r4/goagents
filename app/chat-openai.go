package app

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func OpenAIAgent(c *cli.Context)  {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	apiKey := viper.GetString("OPENAI_API_KEY")

	fmt.Println("====== PURE OPENAI AGENT ======")

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	role := c.String("role")
	question := c.String("prompt")

	print("> ")
	println(question)
	println()

	stream := client.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(role),
				openai.UserMessage(question),
			}),
			Seed: openai.Int(1),
			Model: openai.F(openai.ChatModelGPT4oMini),
			FrequencyPenalty: openai.Float(0.6),
			Temperature: openai.Float(0.4),
	})

	for stream.Next() {
			evt := stream.Current()
			if len(evt.Choices) > 0 {
				print(evt.Choices[0].Delta.Content)
			}
	}
	println()

	if err := stream.Err(); err != nil {
		panic(err)
	}
}
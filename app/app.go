package app

import (
	"github.com/urfave/cli"
)


func RunAgents() *cli.App {
	app := cli.NewApp()
	app.Name = "LLM Agents"
	app.Usage = "Run LLM models"

	flags := []cli.Flag{
			cli.StringFlag{
					Name: "role",
					Value: "You are a fully customizable text-generative AI agent.",
			},
			cli.StringFlag{
					Name: "prompt",
					Value: "In twenty words, tell me who you are.",
			},
	}

	app.Commands = []cli.Command{
	{
		Name: "openai",
		Usage: "Run Pure OpenAI Agent",
		Flags: flags,
		Action: OpenAIAgent,
	},
	{
		Name: "langchain",
		Usage: "Run Langchain Agent With OpenAI model",
		Flags: flags,
		Action: LangchainAgent,
	},
}

	return app
}

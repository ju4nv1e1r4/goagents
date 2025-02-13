package careerApp

import (
	"github.com/urfave/cli"
)


func Match() *cli.App {
	careerApp := cli.NewApp()
	careerApp.Name = "LLM Agents"
	careerApp.Usage = "Run LLM models"

	flags := []cli.Flag{
			cli.StringFlag{
					Name: "role",
					Value: "Você é um assistente de Career Coach e Carreiras. Seu trabalho é analisar vagas. Você receberá um currículo e irá analisar o match entre ele e uma vaga.",

			},
			cli.StringFlag{
					Name: "prompt",
					Value: "Qual o match entre a vaga e o currículo",
			},
	}

	careerApp.Commands = []cli.Command{
	{
		Name: "openai",
		Usage: "Run Career Assistant",
		Flags: flags,
		Action: HRAssistant,
	},
	{
		Name: "langchain",
		Usage: "Run Career Coach",
		Flags: flags,
		Action: CareerCoach,
	},
}

	return careerApp
}

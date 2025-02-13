package careerApp

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/spf13/viper"
)

func loadAPIkey() (string, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return "", fmt.Errorf("erro ao ler o arquivo de configuração: %v", err)
	}
	return viper.GetString("OPENAI_API_KEY"), nil
}

func CallLLM(prompt string, role string) (string, error) {
	apiKey, err := loadAPIkey()
	if err != nil {
		return "", err
	}
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	

	completion, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(role),
			openai.UserMessage(prompt),
		}),
		Seed:  openai.Int(1),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		panic(err)
	}

	return completion.Choices[0].Message.Content, nil
}

func RolesAssistant(text string) (string, error) {
	role := "Você é um assistente de carreiras. Analise uma vaga de emprego e destaque os pontos mais importantes do ponto de vista do empregador."
	return CallLLM(text, role)
}

func CVAssistant(text string) (string, error) {
	role := "Você é um assistente de carreiras. Analise um currículo e faça um resumo da experiência e qualificações do candidato."
	return CallLLM(text, role)
}

func CareerCoach(roleAnalysis string, cvAnalysis string) (string, error) {
	role := "Você é um Coach de Carreiras. Você receberá a análise de uma vaga e um currículo e irá gerar um relatório."

	prompt := fmt.Sprintf(
		"Análise da Vaga:\n%s\n\nAnálise do Currículo:\n%s\n\nCom base nisso, gere um relatório contendo: \n- Comparação entre os requisitos da vaga e o currículo.\n- Recomendações para o candidato.\n- Conclusão geral.",
		roleAnalysis, cvAnalysis,
	)

	return CallLLM(prompt, role)
}
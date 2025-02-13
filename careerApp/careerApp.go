package careerApp

import (
	"sync"
)


func MatchRoleCV(role string, cv string) (string, string) {
	var wg sync.WaitGroup
	var roleResponse, cvResponse string
	var roleErr, cvErr error

	wg.Add(2)

	go func() {
		defer wg.Done()
		roleResponse, roleErr = RolesAssistant(role)
	}()

	go func() {
		defer wg.Done()
		cvResponse, cvErr = CVAssistant(cv)
	}()

	wg.Wait()
	
	if roleErr != nil || cvErr != nil {
		return "Erro na análise da vaga", "Erro na análise do currículo"
	}

	return roleResponse, cvResponse
}

func Report(role string, cv string) string {
    roleAnalysis, cvAnalysis := MatchRoleCV(role, cv)

    finalReport, err := CareerCoach(roleAnalysis, cvAnalysis)
    if err != nil {
        return `{"outputLLM": "Erro ao gerar relatório"}`
    }

    return finalReport
}

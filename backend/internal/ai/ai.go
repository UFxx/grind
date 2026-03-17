package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sunsetsavorer/grind/internal/enums"
	"github.com/sunsetsavorer/grind/internal/exceptions"
)

type AIService struct {
	Model     string
	AccessKey string
	client    *http.Client
}

func NewAIService(model, accessKey string) *AIService {

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	return &AIService{
		Model:     model,
		AccessKey: accessKey,
		client:    client,
	}
}

func (service *AIService) GetChatCompletion(
	systemMessage string,
	userMessage string,
) (string, error) {

	messages := []ChatMessage{
		{
			Role:    enums.AIMessageRoles.System,
			Content: systemMessage,
		},
		{
			Role:    enums.AIMessageRoles.User,
			Content: userMessage,
		},
	}

	requestBody, err := json.Marshal(ChatCompletionRequestBody{
		Model:    service.Model,
		Messages: messages,
	})

	if err != nil {
		return "", exceptions.NewServiceError(err)
	}

	request, err := http.NewRequest(http.MethodPost, "https://api.proxyapi.ru/openrouter/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", exceptions.NewServiceError(err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+service.AccessKey)

	response, err := service.client.Do(request)
	if err != nil {
		return "", exceptions.NewServiceError(err)
	}
	defer response.Body.Close()

	var responseBody ChatCompletionResponse

	if err = json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return "", exceptions.NewServiceError(err)
	}

	if len(responseBody.Choices) == 0 {
		return "", exceptions.NewServiceError(fmt.Errorf("no choices in ai response"))
	}

	return responseBody.Choices[0].Message.Content, nil
}

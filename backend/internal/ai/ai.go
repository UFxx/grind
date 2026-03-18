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
	model     string
	accessKey string
	client    *http.Client
	baseURL   string
}

func NewAIService(model, accessKey, baseURL string) *AIService {

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	return &AIService{
		model:     model,
		accessKey: accessKey,
		client:    client,
		baseURL:   baseURL,
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
		Model:    service.model,
		Messages: messages,
	})

	if err != nil {
		return "", exceptions.NewServiceError(err)
	}

	requestURL := fmt.Sprintf("%s/chat/completions", service.baseURL)

	request, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", exceptions.NewServiceError(err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+service.accessKey)

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

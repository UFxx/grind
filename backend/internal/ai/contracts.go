package ai

type (
	ChatMessage struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}

	ChatCompletionRequestBody struct {
		Model    string        `json:"model"`
		Messages []ChatMessage `json:"messages"`
	}

	ChatCompletionChoice struct {
		Message ChatMessage `json:"message"`
	}

	ChatCompletionResponse struct {
		ID      string                 `json:"id"`
		Choices []ChatCompletionChoice `json:"choices"`
	}
)

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/alvarojhr/tearate-api/internal/database/models"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatGPTResponse struct {
	ID      string `json:"id"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type ChatGPTResponseAnswer struct {
	Score          float32 `json:"score"`
	Feedback       string  `json:"feedback"`
	GeneratedByGTP bool    `json:"generated_by_gtp"`
	Reason         string  `json:"reason"`
}

func RateAnswer(questionText string, points float32, answer models.Answer) (models.Answer, error) {
	apiURL := "https://api.openai.com/v1/chat/completions"

	prompt := fmt.Sprintf("I want you to act as a university professor. Based on the following question:\n%s\n\nRate the following response between 0 and %f :\n\n%s \n\nValidate if response was generated using GTP and return content as a JSON with score, feedback keys and another key to determine if this answer was generated using Chat GPT and why should or shouldn't be possible.", questionText, points, answer.Response)

	requestBody := &ChatGPTRequest{
		Model: "gpt-4o",
		Messages: []ChatMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return models.Answer{}, err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return models.Answer{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_KEY")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.Answer{}, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Answer{}, err
	}

	log.Print(string(responseBody))

	var chatGPTResponse ChatGPTResponse
	err = json.Unmarshal(responseBody, &chatGPTResponse)
	if err != nil {
		return models.Answer{}, err
	}

	b, err := json.MarshalIndent(chatGPTResponse, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	log.Print(string(b))

	// Extract the required fields from the chatGPTResponse.Choices[0].Message.Content
	// assuming it is in the desired JSON format
	var responseAnswer ChatGPTResponseAnswer
	err = json.Unmarshal([]byte(chatGPTResponse.Choices[0].Message.Content), &responseAnswer)
	if err != nil {
		return models.Answer{}, err
	}

	answer.Rate = float32(responseAnswer.Score)
	answer.Feedback = responseAnswer.Feedback

	return answer, nil
}

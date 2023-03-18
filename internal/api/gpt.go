package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func RateAnswer(questionText string, points float32, answerResponse string) (string, error) {
	apiURL := "https://api.openai.com/v1/chat/completions"

	prompt := fmt.Sprintf("Based on the following question:\n%s\n\nRate the following response between 0 and %f :\n\n%s \n\nReturn content as a JSON", questionText, points, answerResponse)

	requestBody := &ChatGPTRequest{
		Model: "gpt-4",
		Messages: []ChatMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer ")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatGPTResponse ChatGPTResponse
	err = json.Unmarshal(responseBody, &chatGPTResponse)
	if err != nil {
		return "", err
	}

	return chatGPTResponse.Choices[0].Message.Content, nil
}

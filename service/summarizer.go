package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	URL = "http://localhost:9000/issue-resolver"
)

type Request struct {
	Content string `json:"content"`
}

type Response struct {
	Result string `json:"result"`
}

func GetResolvedMail(body string) (string, error) {
	request := Request{
		Content: fmt.Sprintf("%s - can GoFr solve the issue? If yes, tell me concisely", body),
	}
	//jsonBody := []byte(fmt.Sprintf(`{"content": "%s - can GoFr solve the issue? If yes, tell me concisely"}`, body))
	out, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("Cannot marshall json request - %v", err)
	}
	req, err := http.NewRequest("GET", URL, bytes.NewBuffer(out))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return "", fmt.Errorf("Cannot get response from LLM")
	}

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the JSON response
	var responseData Response
	if err := json.Unmarshal(respBody, &responseData); err != nil {
		panic(err)
	}

	// Extract the 'result' key
	return responseData.Result, nil
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"gofr.dev/pkg/gofr"
)

const (
	LINKEDIN_API_POST_URL     = "https://api.linkedin.com/v2/ugcPosts"
	LINKEDIN_API_GET_USER_URL = "https://api.linkedin.com/v2/userinfo"
	LINKEDIN_API_TOKEN        = "LINKEDIN_API_TOKEN"
)

type ShareCommentary struct {
	Text string `json:"text"`
}

type ShareContent struct {
	ShareCommentary    ShareCommentary `json:"shareCommentary"`
	ShareMediaCategory string          `json:"shareMediaCategory"`
}

type SpecificContent struct {
	ShareContent ShareContent `json:"com.linkedin.ugc.ShareContent"`
}

type Visibility struct {
	MemberNetworkVisibility string `json:"com.linkedin.ugc.MemberNetworkVisibility"`
}

type PostRequest struct {
	Author          string          `json:"author"`
	LifecycleState  string          `json:"lifecycleState"`
	SpecificContent SpecificContent `json:"specificContent"`
	Visibility      Visibility      `json:"visibility"`
}

func getURN() (string, error) {
	req, err := http.NewRequest("GET", LINKEDIN_API_GET_USER_URL, nil)
	if err != nil {
		return "", err
	}

	accessToken := os.Getenv(LINKEDIN_API_TOKEN)
	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	return response["sub"].(string), nil
}

func generatePostRequestJson(author, content string) ([]byte, error) {
	request := PostRequest{
		Author:         "urn:li:person:" + author,
		LifecycleState: "PUBLISHED",
		SpecificContent: SpecificContent{
			ShareContent: ShareContent{
				ShareCommentary: ShareCommentary{
					Text: content,
				},
				ShareMediaCategory: "NONE",
			},
		},
		Visibility: Visibility{
			MemberNetworkVisibility: "PUBLIC",
		},
	}
	return json.Marshal(request)
}

func createPost(urn, content string) (string, error) {
	jsonData, err := generatePostRequestJson(urn, content)
	if err != nil {
		return "", fmt.Errorf("cannot create the post request for LinkedIn - %w", err)
	}
	req, err := http.NewRequest("POST", LINKEDIN_API_POST_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Add("X-Restli-Protocol-Version", "2.0.0")
	req.Header.Add("Content-Type", "application/json")
	accessToken := os.Getenv(LINKEDIN_API_TOKEN)
	if accessToken == "" {
		return "", fmt.Errorf("no LinkedIn access token found")
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	if resp.StatusCode == 201 {
		// return "Successfully posted the content", nil
		return fmt.Sprintf("%v", response), nil
	}

	return "", fmt.Errorf("cannot post in LinkedIn - %v", response)
}

func PostInLinkedIn(ctx *gofr.Context) (string, error) {
	var request struct {
		Content string `json:"content"`
	}

	if err := ctx.Bind(&request); err != nil {
		return "", err
	}

	if request.Content == "" {
		return "", fmt.Errorf("post content cannot be empty")
	}
	urn, err := getURN()
	if err != nil {
		return "", err
	}
	return createPost(urn, request.Content)
}

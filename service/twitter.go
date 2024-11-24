package main

import (
	"fmt"
	"io"
	"os"

	"bytes"
	"encoding/json"
	"net/http"

	"github.com/dghubble/oauth1"
	"gofr.dev/pkg/gofr"
)

// Function to post a tweet
func postTweet(tweet string) error {
	// OAuth1 configuration
	// Twitter API credentials
	var consumerKey = os.Getenv("X_CONSUMER_KEY")
	var consumerSecret = os.Getenv("X_CONSUMER_SECRET")
	var accessToken = os.Getenv("X_ACCESS_TOKEN")
	var accessSecret = os.Getenv("X_ACCESS_SECRET")
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// Create HTTP client
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter API endpoint for posting a tweet
	twitterEndpoint := "https://api.x.com/2/tweets"

	// Prepare the JSON payload
	payload := map[string]string{"text": tweet}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON payload: %w", err)
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", twitterEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Execute the request
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error posting tweet: %w", err)
	}
	defer resp.Body.Close()

	// Check if the tweet was successfully posted
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		fmt.Printf("Response Body: %s\n", respBody)
		return fmt.Errorf("failed to post tweet. Status: %s", resp.Status)
	}

	return nil
}

// Handler for the POST endpoint
func TweetHandler(ctx *gofr.Context) (interface{}, error) {
	// Parse the JSON body for the "content" field
	var request struct {
		Content string `json:"content"`
	}

	if err := ctx.Bind(&request); err != nil {
		return nil, err
	}

	if request.Content == "" {
		return nil, fmt.Errorf("tweet content cannot be empty")
	}

	// Post the tweet
	if err := postTweet(request.Content); err != nil {
		return nil, err
	}

	return map[string]string{"message": "Tweet posted successfully!"}, nil
}

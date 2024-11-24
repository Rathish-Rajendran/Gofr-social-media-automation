package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"bytes"
	"encoding/json"
	"net/http"

	"github.com/dghubble/oauth1"
	"gofr.dev/pkg/gofr"
	"golang.org/x/exp/rand"
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



// TwitterRequest represents the data to be sent in request
type TwitterRequest struct {
    Content string `json:"content"`
}

// TwitterResponse represents the API response structure
type TwitterResponse struct {
    Result string `json:"result"`
}

// SendGETRequestWithBody sends a GET request with a body
func SendGETRequestWithBody(url string, body TwitterRequest) (*TwitterResponse, error) {
    // Convert body to JSON
    jsonBody, err := json.Marshal(body)
    if err != nil {
        return nil, fmt.Errorf("error marshaling body: %v", err)
    }

    // Create new request
    req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(jsonBody))
    if err != nil {
        return nil, fmt.Errorf("error creating request: %v", err)
    }

    // Set headers
    req.Header.Set("Content-Type", "application/json")

    // Create HTTP client and send request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error sending request: %v", err)
    }
    defer resp.Body.Close()

    // Read response body
    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response: %v", err)
    }

    // Parse response
    var response TwitterResponse
    if err := json.Unmarshal(respBody, &response); err != nil {
        return nil, fmt.Errorf("error parsing response: %v", err)
    }

    return &response, nil
}


func GetTweet(ctx *gofr.Context) (interface{}, error) {
	// Seed the random number generator
	rand.Seed(uint64(time.Now().UnixNano()))
    num := rand.Intn(5)

	buzzWords := []string{ "concurrency", "http api metrics", "observabillity", "ease of building", "GRPC" }
	body := TwitterRequest{
        Content: fmt.Sprintf("Generate a post:twitter on %s", buzzWords[num]),
    }

    response, err := SendGETRequestWithBody("http://127.0.0.1:9000/post-generator", body)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return nil, err
    }

    fmt.Printf("Response result: %s\n", response.Result)
	return string(response.Result), nil
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

package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/emersion/go-message"
)

// Extracts the plain text part of an email while ignoring headers and footers
func extractPlainText(body io.Reader) (string, error) {
	msg, err := message.Read(body)
	if err != nil {
		return "", fmt.Errorf("failed to read message: %v", err)
	}

	// Check if the message is multipart
	mr := msg.MultipartReader()
	if mr == nil {
		// Not multipart, read as single part
		plainText, err := io.ReadAll(msg.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read message body: %v", err)
		}
		return string(plainText), nil
	}

	// Iterate through the parts to find the plain text part
	var plainTextBody string
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to read part: %v", err)
		}

		// Check Content-Type for `text/plain`
		contentType := p.Header.Get("Content-Type")
		if strings.HasPrefix(contentType, "text/plain") {
			bodyBytes, err := io.ReadAll(p.Body)
			if err != nil {
				return "", fmt.Errorf("failed to read plain text part: %v", err)
			}
			plainTextBody = string(bodyBytes)
			break
		}
	}

	if plainTextBody == "" {
		return "", fmt.Errorf("no plain text part found")
	}

	// Clean the body (remove headers and footers)
	cleanedBody := cleanBody(plainTextBody)
	return cleanedBody, nil
}

// Removes common email footers
func cleanBody(body string) string {
	lines := strings.Split(body, "\n")
	var cleanedLines []string
	for _, line := range lines {
		// Skip footer lines (e.g., unsubscribe links)
		if strings.HasPrefix(line, "--") || strings.HasPrefix(line, "To unsubscribe") {
			break
		}
		cleanedLines = append(cleanedLines, line)
	}
	return strings.TrimSpace(strings.Join(cleanedLines, "\n"))
}

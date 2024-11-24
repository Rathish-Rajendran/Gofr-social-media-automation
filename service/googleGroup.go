package main

import (
	"fmt"
	"strings"

	// "log"
	"encoding/json"
	"os"
	"os/exec"

	"gofr.dev/pkg/gofr"
	"gopkg.in/gomail.v2"
)

// func generateReplyMail( sender string, header string, content string ) *imap.Message {
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", os.Getenv("GMAIL"))
// 	m.SetHeader("To", sender)
// 	m.SetHeader("Subject", header)
// 	m.SetBody("text/plain", "Thank you for your email. This is a reply.")

// 	return m
// }

type Email struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Output struct {
	Output []Email `json:"output"`
}

func GoogleGroupReplay(ctx *gofr.Context) (interface{}, error) {
	email := os.Getenv("GMAIL")
	password := os.Getenv("GMAIL_PASSWORD")
	message := gomail.NewMessage()
	message.SetHeader("From", email)
	// senderEmail := fmt.Sprintf("%v@%v", sender.MailboxName, sender.HostName)
	// message.SetHeader("To", senderEmail)
	// message.SetHeader("Subject", msg.Envelope.Subject)
	// message.SetBody("text/plain", replyContent)
	// Connect to Gmail SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, email, password)
	// Send the email
	if err := d.DialAndSend(message); err != nil {
		panic(err)
	}
	println("Email sent successfully!")
	return nil, nil
}

// Process google mail and reponsed
func GoogleGroupHandler(ctx *gofr.Context) (interface{}, error) {
	// Authenticate google client to process the received mail
	// Connect to the Gmail IMAP server

	// Simulate running a command that outputs the JSON (you can replace this with your actual command)
	cmd := exec.Command("./mail.py")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}
	fixedOutput := strings.ReplaceAll(string(output), "'", "\"")

	// // Fix the single quotes to double quotes for valid JSON
	// output = []byte(`"` + string(output)[1:len(output)-1] + `"`)

	// Parse the JSON output
	var result Output
	err = json.Unmarshal([]byte(fixedOutput), &result)
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	for _, email := range result.Output {
		respBody, err := GetResolvedMail(email.Body)
		if err != nil {
			continue
		}
		email.Body = respBody
	}
	// message := gomail.NewMessage()
	// message.SetHeader("From", os.Getenv("GMAIL"))
	// senderEmail := fmt.Sprintf("%v@%v", sender.MailboxName, sender.HostName)
	// message.SetHeader("To", senderEmail)
	// message.SetHeader("Subject", msg.Envelope.Subject)
	// message.SetBody("text/plain", replyContent)
	// // Connect to Gmail SMTP server
	// d := gomail.NewDialer("smtp.gmail.com", 587, email, password)
	// // Send the email
	// if err := d.DialAndSend(message); err != nil {
	// 	panic(err)
	// }
	println("Email sent successfully!")

	// Process each mails

	// if err := <-done; err != nil {
	// 	log.Fatalf("Error while fetching messages: %v", err)
	// }
	return fixedOutput, nil
}

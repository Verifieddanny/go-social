package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"time"

	"github.com/resend/resend-go/v3"
)

type ResendMailer struct {
	fromEmail string
	apiKey    string
	client    *resend.Client
}

func NewResend(apiKey, fromEmail string) *ResendMailer {
	client := resend.NewClient(apiKey)

	return &ResendMailer{
		fromEmail: fromEmail,
		apiKey:    apiKey,
		client:    client,
	}
}

func (m *ResendMailer) Send(templateFile, username, email string, data any, isSandbox bool) (string, error) {
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)

	if err != nil {
		return "", err
	}

	subject := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(subject, "subject", data); err != nil {
		return "", err
	}

	body := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(body, "body", data); err != nil {
		return "", err
	}

	if isSandbox {
		fmt.Printf("--- SANDBOX MODE ---\nTo: %s\nSubject: %s\nBody: %s\n--------------------\n", email, subject, body)
		return "testId", nil
	}
	params := &resend.SendEmailRequest{
		// From:    fmt.Sprintf("%s < %s >", FromName, m.fromEmail),
		From:    m.fromEmail,
		To:      []string{email},
		Html:    body.String(),
		Subject: subject.String(),
	}

	var retryErr error
	for i := 0; i < maxRetries; i++ {
		sent, retryErr := m.client.Emails.Send(params)
		if retryErr != nil {
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}

		return sent.Id, nil

	}

	return "", fmt.Errorf("failed to send email after %d attempts, error: %v", maxRetries, retryErr)
}

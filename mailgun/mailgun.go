package mailgun

import (
	mailgun "github.com/mailgun/mailgun-go"
)

type Email struct {
	From    string   `json:"from" yaml:"from"`
	To      []string `json:"to" yaml:"to"`
	Subject string   `json:"subject" yaml:"subject"`
	Body    string   `json:"body" yaml:"body"`
}

func SendEmail(domain string, apiKey string, email *Email) (string, string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)

	resp, id, err := SendMessage(mg, email.From, email.Subject, email.Body, email.To...)
	if err != nil {
		return "", "", err
	}
	return resp, id, nil
}

func SendMessage(mg mailgun.Mailgun, sender, subject, body string, recipient ...string) (string, string, error) {
	message := mg.NewMessage(sender, subject, body, recipient...)
	resp, id, err := mg.Send(message)

	if err != nil {
		return "", "", err
	}

	return resp, id, nil
}

func SendFromFile(apikey, domain, filePath string) error {

	return nil
}

package mailgun

type MailFile struct {
	From    string
	To      []string
	Subject string
	Body    string
}

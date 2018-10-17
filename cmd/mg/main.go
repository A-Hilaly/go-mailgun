package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/a-hilaly/go-mailgun/mailgun"
)

var (
	domain       *string
	apiKey       *string
	apiKeyEnv    *string
	from         *string
	to           *string
	subject      *string
	body         *string
	bodyFilePath *string
	useFile      *string
)

func init() {
	domain = flag.String("domain", "example.com", "domain name")
	apiKey = flag.String("apikey", "", "API Key")
	apiKeyEnv = flag.String("apikeyenv", "MAILGUN_API", "")
	from = flag.String("from", "", "from email")
	to = flag.String("to", "", "to emails format x@x.x,y@x.x")
	subject = flag.String("subject", "", "email subject")
	body = flag.String("body", "", "body subject")
	bodyFilePath = flag.String("", "", "")
	useFile = flag.String("file", "", "")
}

func sendFromArgs() {
	var emailBody string
	if *apiKey == "" {
		envkey := os.Getenv(*apiKeyEnv)
		if envkey == "" {
			log.Fatalln("no api key provided")
		}
		*apiKey = envkey
	}

	if *bodyFilePath != "" {
		bd, err := ioutil.ReadFile(*bodyFilePath)
		if err != nil {
			log.Fatalf("failed to read body from file: %v", err)
		}

		emailBody = string(bd)

	} else {
		emailBody = *body

	}

	receivers := strings.Split(*to, ",")
	resp, id, err := mailgun.SendEmail(*domain, *apiKey, &mailgun.Email{
		From:    *from,
		To:      receivers,
		Subject: *subject,
		Body:    emailBody,
	})

	if err != nil {
		log.Fatalf("failed to deliver emails: %v", err)
	}

	log.Printf("\nid: %s\nresponse: %s\n", id, resp)
}

func sendFromFile() {

}

func main() {
	flag.Parse()
	if *useFile != "" {
		sendFromFile()
	} else {
		sendFromArgs()
	}
}

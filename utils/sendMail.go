package utils

import (
	"encoding/json"
	"go-mail/model"
	"log"
	"net/smtp"
	"os"

	"github.com/adjust/rmq/v4"
)

func SendEmail(delivery rmq.Delivery) {
	// var task interface{}
	var T model.EmailTemplate
	json.Unmarshal([]byte(delivery.Payload()), &T)

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + T.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + T.Body)

	// Create authentication
	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), os.Getenv("SMTP_HOST"))
	log.Println("sending email...")
	// Send message
	err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, os.Getenv("EMAIL"), T.To, []byte(msg))
	if err != nil {
		log.Fatal("Message is not sent to smtp server!")
	}
	log.Println("Message sent to smtp server")
	log.Println("Message sent.")

	delivery.Ack()
}

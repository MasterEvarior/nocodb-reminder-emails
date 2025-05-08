package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(smtpAddress string, sender string, receiver string, data []Row) error {
	receivers := []string{receiver}
	msg := fmt.Sprintf("To: %s\r\n"+
		"Subject: NocoDB Reminder\r\n"+
		"\r\n", receiver)

	for _, row := range data {
		msg += fmt.Sprintf("- %s (%s)(%s) \r\n", row.Title, row.Subject, row.Status)
	}

	err := smtp.SendMail(smtpAddress, nil, sender, receivers, []byte(msg))
	if err != nil {
		log.Printf("Could not send email because of the following error: %v", err)
		return err
	}
	log.Printf("Sent a reminder email from %s to %s", sender, receiver)

	return nil
}

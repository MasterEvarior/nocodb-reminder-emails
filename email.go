package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(smtpAddress string, sender string, receiver string, data []Row) error {
	receivers := []string{receiver}
	msg := fmt.Sprintf(line("To: %s")+line("Subject: NocoDB Reminder")+line(""), receiver)

	for _, row := range data {
		msg += fmt.Sprintf(line("- %s (%s)(%s)"), row.Title, row.Subject, row.Status)
	}

	err := smtp.SendMail(smtpAddress, nil, sender, receivers, []byte(msg))
	if err != nil {
		log.Printf("Could not send email because of the following error: %v", err)
		return err
	}
	log.Printf("Sent a reminder email from %s to %s", sender, receiver)

	return nil
}

func line(content string) string {
	return content + "\r\n"
}

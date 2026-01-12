package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
)

func processRecipients(id int, RecipientChannel chan Recipient, wg *sync.WaitGroup) {
	for recipient := range RecipientChannel {
		smtpHost := "localhost"
		smtpPort := "1025"
		formattedMessage := "To: " + recipient.Email + "\r\n" +
			"Subject: Greetings " + recipient.Name + "\r\n" +
			"\r\n" +
			"Hello " + recipient.Name + ",\r\n" +
			"This is a test email.\r\n"

		err := smtp.SendMail(smtpHost+":"+smtpPort, nil, "kartikeywariyal706@gmail.com", []string{recipient.Email}, []byte(formattedMessage))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i, "Worker %d: Email sent to %s\n", id, recipient.Email)
		i++
	}
	defer wg.Done()
}

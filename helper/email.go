package helper

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(email, otp, subjectMail, emailBody string) {
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal(err1)
	}

	from := os.Getenv("Email")
	password := os.Getenv("Password")

	// Receiver email address.
	to := []string{email}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.

	// subject := "Subject: Activita Email Verification\n"
	subject := subjectMail

	messageBody := emailBody

	mainMessage := fmt.Sprintf(messageBody, otp)

	body := mainMessage
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

}

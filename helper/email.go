package helper

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(email, otp string) {
	err1:= godotenv.Load()
	if err1!=nil{
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

	subject := "Subject: Activita Email Verification\n"

	mainMessage := fmt.Sprintf("<body>Welcome to Activita, Your Profile verification code is <h2 style=\"text-align:center;\"><span style=\"font-size:40px;border:2px solid black;padding:10px\">%v</span></h2> \n</body>", otp)

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
	fmt.Println("Email Sent Successfully!")

}

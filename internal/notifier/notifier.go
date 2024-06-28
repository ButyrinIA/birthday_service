package notifier

import (
	"fmt"
	"log"
	"net/smtp"

	"rutube/internal/models"
)

func SendBirthdayNotification(user models.User) error {
	from := "your-email@example.com"
	password := "your-email-password"
	to := user.Email
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	msg := []byte(fmt.Sprintf("Subject: Happy Birthday!\n\nHappy Birthday, %s!", user.Username))

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		log.Printf("Failed to send birthday email to %s: %s", user.Email, err)
		return err
	}
	return nil
}

package service

import (
    "fmt"
    "net/smtp"
    "os"
)

func SendEmail(name, email, message string) error {
    // Ambil config dari .env
    smtpHost := os.Getenv("SMTP_HOST")  // smtp.gmail.com
    smtpPort := os.Getenv("SMTP_PORT")  // 587
    smtpUser := os.Getenv("SMTP_USER")  // email kamu
    smtpPass := os.Getenv("SMTP_PASS")  // app password

    auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

    body := fmt.Sprintf(
        "From: %s <%s>\nMessage:\n%s", name, email, message,
    )

    return smtp.SendMail(
        smtpHost+":"+smtpPort,
        auth,
        smtpUser,
        []string{smtpUser}, // kirim ke email kamu sendiri
        []byte(body),
    )
}
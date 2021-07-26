package main

import (
	"fmt"
	"net/smtp"
)

func sendEmail(email string, link string) {
	// sender data
	from := "mask4youresetpwnoreply@gmail.com" //ex: "John.Doe@gmail.com"
	password := "passw0rdret!"

	// receiver address
	toEmail := email // ex: "Jane.Smith@yahoo.com"

	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	// message
	message := []byte("To: " + email + "\r\n" +
		"Subject: Password Reset Request\r\n" +
		"\r\n" +
		"You have received this email because a password reset request for Crazy for Masks account was received. The reset link will only be valid for 30mins. Click the link to reset your password: \r\n" + link)

	// athentication data
	auth := smtp.PlainAuth("", from, password, host)

	// send mail
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}

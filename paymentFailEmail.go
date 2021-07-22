package main

import (
	"fmt"
	"net/smtp"
)

// https://morioh.com/p/d167bdd617c4
func sendPaymentFailEmail(email string, price string) {
	// sender data
	from := "mask4youpayment@gmail.com" //ex: "John.Doe@gmail.com"
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
		"Subject: Payment Failed\r\n" +
		"\r\n" +
		"Your payment of " + price + " for your order on crazy4masks has failed")

	// athentication data
	auth := smtp.PlainAuth("", from, password, host)

	// send mail
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}

// https://github.com/matcornic/hermes/blob/master/examples/reset.go
// This is a reset password email generation function

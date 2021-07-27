package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"
)

// MinTokenLength is the minimum allowed length of token string.
// It is useful for avoiding DoS attacks with very long tokens: before passing
// a token to VerifyToken function, check that it has length less than [the
// maximum login length allowed in your application] + MinTokenLength.
var (
	//MinTokenLength = authcookie.MinLength
	MinTokenLength = 8
)

var (
	ErrMalformedToken = errors.New("malformed token")
	ErrExpiredToken   = errors.New("token expired")
	ErrWrongSignature = errors.New("wrong token signature")
)

func generateSecretKey() []byte {
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error Generating Key")
	}
	return key
}

func generateVoucher3Off(email string) string {
	token, err := maker.CreateToken(email, time.Hour*24)
	fmt.Println(token)
	if err != nil {
		fmt.Println("Error creating token")
		return ""
	}
	resultToken := "3" + token
	return resultToken
}

func generateVoucher4Off(email string) string {
	token, err := maker.CreateToken(email, time.Hour*24)
	if err != nil {
		fmt.Println("Error creating token")
		return ""
	}
	resultToken := "4" + token
	return resultToken
}

func generateVoucher5Off(email string) string {
	token, err := maker.CreateToken(email, time.Hour*24)
	if err != nil {
		fmt.Println("Error creating token")
		return ""
	}
	resultToken := "5" + token
	return resultToken
}

func checkVoucher(resultToken string) {
	discount := resultToken[0]
	token := resultToken[1:]
	switch discount {
	case 3:
		result := verifyVoucher(token)
		if result {
			fmt.Println("3Off")
		}
	case 4:
		result := verifyVoucher(token)
		if result {
			fmt.Println("4Off")
		}
	case 5:
		result := verifyVoucher(token)
		if result {
			fmt.Println("5Off")
		}
	default:
		fmt.Println("Invalid Input")
	}
}

func verifyVoucher(token string) bool {
	// Verify whether token is valid
	_, err := maker.VerifyToken(token)
	if err != nil {
		fmt.Println("Invalid voucher")
		return false
	}
	return true
}

func pointsConverter(points int) (voucher string) {
	//email := "email" // user email
	discount := points / 100
	if discount >= 3 && discount < 4 {
		voucher = generateVoucher3Off("hello")
	} else if discount >= 4 && discount < 5 {
		voucher = generateVoucher4Off("hello")
	} else if discount >= 5 {
		voucher = generateVoucher5Off("hello")
	}
	return voucher
}

/*
func resetBuyerPassword(email string) {
	token, err := maker.CreateToken(email, time.Minute*30)
	if err != nil {
		fmt.Println("Error creating token")
		return
	}

	Link := "https://localhost:5221/buyertoken/" + token
	sendEmail(email, Link)
}

func resetBuyerPasswordLinkClicked(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	tokenStr := vars["token"]

	// Verify whether token is valid
	_, err := maker.VerifyToken(tokenStr)
	if err != nil {
		io.WriteString(res, `
 			<html>
			 <meta http-equiv='refresh' content='5; url=/ '/>
			 Password reset link has expired! <br>
			 You will be redirected shortly in 5 seconds...<br>
 			</html>
 		`)
		return
	}
	http.Redirect(res, req, "/buyerresetchangepassword", http.StatusSeeOther)
}
*/

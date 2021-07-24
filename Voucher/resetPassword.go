package main

import (
	"cookie/token"
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/dchest/authcookie"
)

// MinTokenLength is the minimum allowed length of token string.
// It is useful for avoiding DoS attacks with very long tokens: before passing
// a token to VerifyToken function, check that it has length less than [the
// maximum login length allowed in your application] + MinTokenLength.
var (
	MinTokenLength = authcookie.MinLength
)

var (
	ErrMalformedToken = errors.New("malformed token")
	ErrExpiredToken   = errors.New("token expired")
	ErrWrongSignature = errors.New("wrong token signature")
	maker             token.Maker
)

func generateSecretKey() []byte {
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error Generating Key")
	}
	return key
}

func generateVoucher(email string) string {
	token, err := maker.CreateToken(email, time.Hour*24)
	if err != nil {
		fmt.Println("Error creating token")
		return ""
	}
	return token
}

func verifyVoucher(token string) {
	// Verify whether token is valid
	_, err := maker.VerifyToken(token)
	if err != nil {
		fmt.Println("Invalid voucher")
		return
	}
	fmt.Println("Voucher Valid")
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

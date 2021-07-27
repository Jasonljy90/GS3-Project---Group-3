package main

import (
	"cookie/token"
	"fmt"
)

var (
	maker token.Maker
	err   error
)

func main() {
	email := "jasonljy90@gmail.com"
	// Generate key and convert to string for voucher generation feature
	key := generateSecretKey()
	keyStr := string(key)
	maker, err = token.NewJWTMaker(keyStr)
	if err != nil {
		fmt.Println("Error generating token maker!")
	}

	voucher := pointsConverter(700, email)
	checkVoucher(voucher)
}

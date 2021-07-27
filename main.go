package main

import (
	"cookie/token"
	"fmt"
	"sync"
)

var (
	wg    sync.WaitGroup
	maker token.Maker
	err   error
)

func main() {
	//email := "jasonljy90@gmail.com"
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Generate key and convert to string for voucher generation feature
		key := generateSecretKey()
		keyStr := string(key)
		maker, err = token.NewJWTMaker(keyStr)
		if err != nil {
			fmt.Println("Error generating token maker!")
		}
	}()

	//voucher := pointsConverter(300, email)
	voucher := pointsConverter(300)
	checkVoucher(voucher)
}

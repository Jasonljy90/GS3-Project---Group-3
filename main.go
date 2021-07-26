package main

import (
	"cookie/token"
	"fmt"
)

func main() {
	go func() {
		defer wg.Done()
		// Generate key and convert to string for voucher generation feature
		key := generateSecretKey()
		keyStr := string(key)
		maker, err := token.NewJWTMaker(keyStr)
		if err != nil {
			fmt.Println("Error generating token maker!")
		}
	}()
}

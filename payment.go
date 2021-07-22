package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Payment struct {
	Amount      float64
	ReceiptMail string
	ProductName string
}

func sendPayment(amount float64, quantity int, email string) bool {
	// Convert dollar and cents to cents
	amountCents := amount * 100

	// Convert int to string
	quantityStr := strconv.Itoa(quantity)

	// Create a struct object to be sent to the stripe server to process payment
	p := Payment{
		Amount:      amountCents,
		ReceiptMail: email,
		ProductName: quantityStr,
	}

	//Encode the data
	postBody, _ := json.Marshal(p) // Convert struct to json string
	responseBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("http://localhost:8000/stripe/payment", "application/json", responseBody)

	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	// Checks whether payment to stripe is successful. Returns true if successful else return false.
	if sb != "Payment Unsuccessful" {
		return true
	} else {
		return false
	}
	//log.Printf(sb)
}

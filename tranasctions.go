package main

// Transaction represet amount transfer
// made by user in favour of service
// validated by the subscription server
type Transaction struct {
	ID             int    //
	PaymentTransID string // ID assinged by payment service provider
	PaymentData    string // complete struct of payment
}

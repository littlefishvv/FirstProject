package strategy
func ExamplePayByCash() {
	payment := NewPayment("Ada", "", 123, &CashStrategy{})
	payment.Pay()
	// Output:
	// Pay $123 to Ada by cash
}

func ExamplePayByZfb() {
	payment := NewPayment("Bob", "0002", 888, &ZfbStrategy{})
	payment.Pay()
	// Output:
	// Pay $888 to Bob by zfb
}
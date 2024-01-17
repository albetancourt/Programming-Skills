package input

import "fmt"

func WelcomeUser() {
	fmt.Print("Welcome to the currency converter!\n\n")
}

func GetText(message string) string {
	fmt.Print(message)
	var text string
	fmt.Scan(&text)

	return text
}

func GetAmount(message string) float64 {
	fmt.Print(message)
	var amount float64
	fmt.Scan(&amount)

	return amount
}

package main

import (
	"fmt"
	"strings"

	"example.com/converter/currency"
	"example.com/converter/input"
)

func main() {
	input.WelcomeUser()

	for {
		initialCurrency := input.GetText("Choose your initial currency: ")
		targetCurrency := input.GetText("Choose the currency they want to exchange to: ")
		amountToExchange := input.GetAmount("Enter the amount to exchange: ")

		if !currency.IsValidAmount(amountToExchange, targetCurrency) {
			fmt.Println("Invalid amount. Please choose an amount within the allowed limits.")
			continue
		}

		exchangedAmount := currency.Convert(initialCurrency, targetCurrency, amountToExchange)

		withdrawAnswer := input.GetText("Do you want to withdraw funds? (yes/no): ")
		userWantsToWithDraw := strings.ToLower(withdrawAnswer) == "yes"

		if userWantsToWithDraw {
			withdrawalFee := currency.GetWithdrawalFee(exchangedAmount)
			withdrawnAmount := exchangedAmount - withdrawalFee
			fmt.Printf("Withdrawal successful. Amount withdrawn: %g %s\n", withdrawnAmount, targetCurrency)
		} else {
			fmt.Printf("Exchange successful. Amount received: %g %s\n", exchangedAmount, targetCurrency)
		}

		anotherOperationAnswer := input.GetText("Do you want to perform another operation? (yes/no): ")
		userWantsAnotherOperation := strings.ToLower(anotherOperationAnswer) == "yes"

		if !userWantsAnotherOperation {
			fmt.Println("Goodbye!")
			return
		}
	}

}

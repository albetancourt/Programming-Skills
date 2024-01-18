package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var availableTickets map[string]string
var purchasedTickets map[string]string
var selectedTickets []string
var winningTicket string

const costPerTicket int = 1
const maxTicketsPerUser = 2

func main() {
	initializeAvailableTickets()
	initializePurchasedTickets()
	selectTickets()

	paymentMethod, _ := getTextFromUser("What will be your payment method? (cash/card): ")
	err := processPayment(paymentMethod)
	if err != nil {
		fmt.Println(err)
		return
	}

	sellTickets()
	displayWinningTicket()
}

func initializeAvailableTickets() {
	availableTickets = map[string]string{
		"5678B": "", "9876C": "", "2345D": "", "6789E": "", "3456F": "",
		"8765G": "", "4321H": "", "7890J": "", "5432K": "", "2109L": "",
		"8765M": "", "1357N": "", "2468P": "", "6543Q": "", "7891R": "",
		"3579S": "", "9821T": "", "4682U": "", "5763V": "", "1234A": "",
	}
}

func initializePurchasedTickets() {
	purchasedTickets = make(map[string]string)
}

func ticketIsAvailable(ticketCode string) bool {
	_, ticketIsAvailable := availableTickets[ticketCode]

	return ticketIsAvailable
}

func getTextFromUser(message string) (string, error) {
	fmt.Print(message)

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	line = strings.TrimSuffix(line, "\n")

	return line, nil
}

func getIntFromUser(message string) int {
	fmt.Print(message)
	var num int
	fmt.Scan(&num)

	return num
}

func selectTickets() {
	for {
		ticketCode, _ := getTextFromUser("Enter the ticket you want to buy: ")
		if !ticketIsAvailable(ticketCode) {
			fmt.Println("Ticket is not available")
			continue
		}

		selectedTickets = append(selectedTickets, ticketCode)

		if len(selectedTickets) >= maxTicketsPerUser {
			return
		}

		answer, _ := getTextFromUser("Do you want to buy another ticket? (yes/no): ")
		wantToBuyAnotherTicket := strings.ToLower(answer) == "yes"

		if !wantToBuyAnotherTicket {
			return
		}
	}
}

func sellTickets() {
	for _, ticket := range selectedTickets {
		purchasedTickets[ticket] = ""
		delete(availableTickets, ticket)
	}
}

func processPayment(method string) error {
	switch method {
	case "cash":
		return processCashPayment()

	case "card":
		return processCardPayment()

	default:
		return errors.New("Invalid payment method selected.")
	}
}

func processCashPayment() error {
	one_usd_bills := getIntFromUser("Please enter the number of 1 USD bills: ")
	five_usd_bills := getIntFromUser("Please enter the number of 5 USD bills: ")

	userPayment := one_usd_bills*1 + five_usd_bills*5
	totalCost := calculateTotalCost()

	if userPayment < totalCost {
		return errors.New("The payment is insufficient")
	}

	if userPayment > totalCost {
		change := userPayment - totalCost
		fmt.Printf("Here is your change: %d USD\n", change)
	}

	return nil
}

func processCardPayment() error {
	displayTotalCost()
	getTextFromUser("Please enter your 4-digit PIN:")
	fmt.Println("Thank you for your payment! Your purchase is complete.")

	return nil
}

func calculateTotalCost() int {
	return len(selectedTickets) * costPerTicket
}

func displayTotalCost() {
	totalCost := calculateTotalCost()
	fmt.Printf("The total cost is %d USD\n", totalCost)
}

func displayWinningTicket() {
	generateWinningTicket()
	fmt.Printf("The winning ticket is: %s\n", winningTicket)

	if isInSlice(winningTicket, selectedTickets) {
		fmt.Println("Congratulations! You won the lottery!")
	} else {
		fmt.Println("Sorry, better luck next time.")
	}
}

func generateWinningTicket() {
	rand.Seed(time.Now().UnixNano())

	digits := fmt.Sprintf("%04d", rand.Intn(10000))
	letter := string('A' + rand.Intn(26))

	ticketCode := digits + letter

	winningTicket = ticketCode
}

func isInSlice(target string, slice []string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

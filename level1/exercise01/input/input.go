package input

import "fmt"

func DisplayMenu() {
	fmt.Println("")
	fmt.Println("1. Player Review")
	fmt.Println("2. Compare two players")
	fmt.Println("3. Identify the fastest player")
	fmt.Println("4. Identify the top goal scorer")
	fmt.Println("5. Identify the player with the most assists")
	fmt.Println("6. Identify the player with the highest passing accuracy")
	fmt.Println("7. Identify the player with the most defensive involvements")
	fmt.Println("8. Exit")
	fmt.Println("")
}

func GetInt(message string) int64 {
	var num int64
	fmt.Print(message)
	fmt.Scan(&num)

	return num
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

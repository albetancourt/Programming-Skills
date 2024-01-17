package main

import (
	"fmt"

	"example.com/manchester/db"
	"example.com/manchester/input"
	"example.com/manchester/player"
)

func main() {
	err := db.InitDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		input.DisplayMenu()
		choice := input.GetInt("Your choice: ")

		switch choice {

		case 1:
			jerseyNumber := input.GetInt("Enter the player's jersey number: ")
			p, err := player.GetByJersey(jerseyNumber)

			if err != nil {
				fmt.Printf("An error occurred: %s\n", err)
				continue
			}

			fmt.Print(player.GetCharacteristics(p))

		case 2:
			jerseyNumber1 := input.GetInt("Enter player 1's jersey number: ")
			jerseyNumber2 := input.GetInt("Enter player 2's jersey number: ")

			p1, err := player.GetByJersey(jerseyNumber1)
			if err != nil {
				fmt.Printf("An error occurred while trying to get player 1's data: %s\n", err)
				continue
			}

			p2, err := player.GetByJersey(jerseyNumber2)
			if err != nil {
				fmt.Printf("An error occurred while trying to get player 2's data: %s\n", err)
				continue
			}

			fmt.Print(player.GetComparison(p1, p2))

		case 3:
			p, err := player.GetFastest()

			if err != nil {
				fmt.Printf("An error occurred: %s\n", err)
				continue
			}

			fmt.Print(player.GetCharacteristics(p))

		case 4:
			p, err := player.GetTopScorer()

			if err != nil {
				fmt.Printf("An error occurred: %s\n", err)
				continue
			}

			fmt.Print(player.GetCharacteristics(p))

		case 5:
			p, err := player.GetTopAssistProvider()

			if err != nil {
				fmt.Printf("An error occurred: %s\n", err)
				continue
			}

			fmt.Print(player.GetCharacteristics(p))

		case 6:
			p, err := player.GetBestPasser()

			if err != nil {
				fmt.Printf("An error occurred: %s\n", err)
				continue
			}

			fmt.Print(player.GetCharacteristics(p))

		case 7:
			p, err := player.GetBestDefender()

			if err != nil {
				fmt.Printf("An error occurred: %s\n", err)
				continue
			}

			fmt.Print(player.GetCharacteristics(p))

		default:
			fmt.Println("Goodbye!")
			return
		}
	}

}

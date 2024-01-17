package main

import (
	"fmt"
	"strings"

	"example.com/travel/db"
	"example.com/travel/input"
	"example.com/travel/travel"
)

func main() {
	err := db.InitDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	input.WelcomeUser()

	for {
		budget := input.GetBudget()

		preferredSeason, err := input.GetPreferredSeason()
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			continue
		}

		preferredActivities, err := input.GetPreferredActivities()
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			continue
		}

		suggestedDestination, err := travel.GetSuggestedDestination(budget, preferredSeason, preferredActivities)
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			continue
		}

		if suggestedDestination == nil {
			fmt.Println("Unfortunately, we couldn't find a destination that perfectly aligns with your budget, preferred activities, and the selected season.\n")
		} else {
			fmt.Printf("\nBased on your preferences, we recommend %s as your ideal destination!\n\n", suggestedDestination.Name)
		}

		anotherTripAnswer, _ := input.GetText("Do you want to plan another trip? (yes/no): ")
		userWantsToContinue := strings.ToLower(anotherTripAnswer) == "yes"

		if !userWantsToContinue {
			fmt.Println("Goodbye!")
			return
		}
	}

}

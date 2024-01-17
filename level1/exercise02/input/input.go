package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/travel/travel"
)

func WelcomeUser() {
	fmt.Print("Welcome to the travel agency system!\n\n")
}

func GetBudget() int64 {
	budget := GetInt("What is your budget for the trip in USD?: ")

	return budget
}

func GetPreferredSeason() (*travel.Season, error) {
	PresentSeasonOptions()
	userChoice := GetInt("Your choice: ")
	season, err := travel.GetSeasonByID(userChoice)

	if err != nil {
		return nil, err
	}

	return season, nil
}

func PresentSeasonOptions() error {
	seasons, err := travel.GetFormattedSeasons()

	if err != nil {
		return err
	}

	fmt.Printf("\nWhat season are you planning to travel in?\n\n")
	fmt.Print(*seasons)
	fmt.Println("")

	return nil
}

func GetPreferredActivities() ([]*travel.Activity, error) {
	PresentActivityOptions()
	userChoices, err := GetText("Your choices: ")

	if err != nil {
		return nil, err
	}

	activityIDs := strings.Split(userChoices, ",")
	var activities []*travel.Activity

	for _, item := range activityIDs {
		id, err := strconv.ParseInt(item, 10, 64)

		if err != nil {
			return nil, err
		}

		a, err := travel.GetActivityByID(id)

		if err != nil {
			return nil, err
		}

		activities = append(activities, a)
	}

	return activities, nil

}

func PresentActivityOptions() error {
	activities, err := travel.GetFormattedActivities()

	if err != nil {
		return err
	}

	fmt.Printf("\nPlease select three activities you would like to do? (write them separated by commas):\n\n")
	fmt.Print(*activities)
	fmt.Println("")

	return nil
}

func GetInt(message string) int64 {
	fmt.Print(message)
	var num int64
	fmt.Scan(&num)

	return num
}

func GetFloat(message string) float64 {
	fmt.Print(message)
	var num float64
	fmt.Scan(&num)

	return num
}

func GetText(message string) (string, error) {
	fmt.Print(message)

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	line = strings.TrimSuffix(line, "\n")

	return line, nil
}

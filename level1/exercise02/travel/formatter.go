package travel

import (
	"fmt"
)

func GetFormattedActivities() (*string, error) {
	var s string

	activities, err := GetAllActivities()

	if err != nil {
		return nil, err
	}

	for _, activity := range activities {
		s += fmt.Sprintf("%d. %s\n", activity.ID, activity.Name)
	}

	return &s, nil
}

func GetFormattedSeasons() (*string, error) {
	var s string

	seasons, err := GetAllSeasons()

	if err != nil {
		return nil, err
	}

	for _, season := range seasons {
		s += fmt.Sprintf("%d. %s\n", season.ID, season.Name)
	}

	return &s, nil
}

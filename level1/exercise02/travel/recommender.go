package travel

import (
	"database/sql"

	"example.com/travel/db"
)

func GetSuggestedDestination(
	budget int64,
	preferredSeason *Season,
	preferredActivities []*Activity) (*Destination, error) {

	query := `
		SELECT d.id, d.name
		FROM destinations d
		INNER JOIN destination_activities da ON d.id = da.destination_id
		INNER JOIN season_destinations sd ON d.id = sd.destination_id
		INNER JOIN seasons s ON s.id = sd.season_id
		INNER JOIN season_cost sc ON s.id = sc.season_id
		WHERE s.id = ? AND sc.cost <= ? AND da.activity_id IN (?, ?, ?)
		GROUP BY d.id
		ORDER BY COUNT(da.activity_id) DESC
		LIMIT 1
	`

	args := getQueryArgs(budget, preferredSeason, preferredActivities)
	row := db.DB.QueryRow(query, args...)

	var d Destination
	err := row.Scan(&d.ID, &d.Name)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &d, nil
}

func getActivityIDs(activities []*Activity) []int64 {
	var ids []int64

	for _, a := range activities {
		ids = append(ids, a.ID)
	}

	return ids
}

func getQueryArgs(
	budget int64,
	preferredSeason *Season,
	preferredActivities []*Activity) []interface{} {

	preferredActivityIDs := getActivityIDs(preferredActivities)

	args := make([]interface{}, 0, len(preferredActivityIDs)+2)
	args = append(args, preferredSeason.ID, budget)

	for _, id := range preferredActivityIDs {
		args = append(args, id)
	}

	return args

}

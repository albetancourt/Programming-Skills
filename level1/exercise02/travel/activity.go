package travel

import (
	"example.com/travel/db"
)

type Activity struct {
	ID   int64
	Name string
}

func GetActivityByID(id int64) (*Activity, error) {
	sql := `
		SELECT id, name
		FROM activities
		WHERE id = ?
	`

	row := db.DB.QueryRow(sql, id)

	var a Activity
	err := row.Scan(&a.ID, &a.Name)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func GetAllActivities() ([]*Activity, error) {
	sql := "SELECT * FROM activities ORDER by id"
	rows, err := db.DB.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var activities []*Activity

	for rows.Next() {
		var a Activity
		err := rows.Scan(&a.ID, &a.Name)

		if err != nil {
			return nil, err
		}

		activities = append(activities, &a)
	}

	return activities, nil
}

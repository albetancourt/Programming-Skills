package travel

import "example.com/travel/db"

type Season struct {
	ID   int64
	Name string
	Cost int64
}

func GetSeasonByID(id int64) (*Season, error) {
	sql := `
		SELECT s.id, s.name, sc.cost
		FROM seasons s, season_cost sc
		WHERE s.id = sc.season_id AND s.id = ?
	`

	row := db.DB.QueryRow(sql, id)

	var s Season
	err := row.Scan(&s.ID, &s.Name, &s.Cost)

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func GetAllSeasons() ([]*Season, error) {
	sql := `
		SELECT s.id, s.name, sc.cost
		FROM seasons s, season_cost sc
		WHERE s.id = sc.season_id
		ORDER BY s.id
	`

	rows, err := db.DB.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var seasons []*Season

	for rows.Next() {
		var s Season
		err := rows.Scan(&s.ID, &s.Name, &s.Cost)

		if err != nil {
			return nil, err
		}

		seasons = append(seasons, &s)
	}

	return seasons, nil
}

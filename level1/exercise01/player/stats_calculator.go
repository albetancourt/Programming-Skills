package player

import (
	"example.com/manchester/db"
)

func GetFastest() (*Player, error) {
	sql := `
	SELECT p.id, p.name, p.jersey_number, c.goals, c.speed_points, c.assist_points, c.passing_accuracy_points, c.defensive_involvements
	FROM players p, player_characteristics c
	WHERE p.id = c.id
	ORDER BY speed_points DESC
	LIMIT 1
`

	row := db.DB.QueryRow(sql)

	var player Player
	player.setCharacteristics(NewPlayerCharacteristics())

	err := row.Scan(
		&player.ID,
		&player.Name,
		&player.JerseyNumber,
		&player.Characteristics.Goals,
		&player.Characteristics.SpeedPoints,
		&player.Characteristics.AssistPoints,
		&player.Characteristics.PassingAccuracyPoints,
		&player.Characteristics.DefensiveInvolvements)

	if err != nil {
		return nil, err
	}

	return &player, nil
}

func GetTopScorer() (*Player, error) {
	sql := `
	SELECT p.id, p.name, p.jersey_number, c.goals, c.speed_points, c.assist_points, c.passing_accuracy_points, c.defensive_involvements
	FROM players p, player_characteristics c
	WHERE p.id = c.id
	ORDER BY goals DESC
	LIMIT 1
`

	row := db.DB.QueryRow(sql)

	var player Player
	player.setCharacteristics(NewPlayerCharacteristics())

	err := row.Scan(
		&player.ID,
		&player.Name,
		&player.JerseyNumber,
		&player.Characteristics.Goals,
		&player.Characteristics.SpeedPoints,
		&player.Characteristics.AssistPoints,
		&player.Characteristics.PassingAccuracyPoints,
		&player.Characteristics.DefensiveInvolvements)

	if err != nil {
		return nil, err
	}

	return &player, nil
}

func GetTopAssistProvider() (*Player, error) {
	sql := `
	SELECT p.id, p.name, p.jersey_number, c.goals, c.speed_points, c.assist_points, c.passing_accuracy_points, c.defensive_involvements
	FROM players p, player_characteristics c
	WHERE p.id = c.id
	ORDER BY assist_points DESC
	LIMIT 1
`

	row := db.DB.QueryRow(sql)

	var player Player
	player.setCharacteristics(NewPlayerCharacteristics())

	err := row.Scan(
		&player.ID,
		&player.Name,
		&player.JerseyNumber,
		&player.Characteristics.Goals,
		&player.Characteristics.SpeedPoints,
		&player.Characteristics.AssistPoints,
		&player.Characteristics.PassingAccuracyPoints,
		&player.Characteristics.DefensiveInvolvements)

	if err != nil {
		return nil, err
	}

	return &player, nil
}

func GetBestPasser() (*Player, error) {
	sql := `
	SELECT p.id, p.name, p.jersey_number, c.goals, c.speed_points, c.assist_points, c.passing_accuracy_points, c.defensive_involvements
	FROM players p, player_characteristics c
	WHERE p.id = c.id
	ORDER BY passing_accuracy_points DESC
	LIMIT 1
`

	row := db.DB.QueryRow(sql)

	var player Player
	player.setCharacteristics(NewPlayerCharacteristics())

	err := row.Scan(
		&player.ID,
		&player.Name,
		&player.JerseyNumber,
		&player.Characteristics.Goals,
		&player.Characteristics.SpeedPoints,
		&player.Characteristics.AssistPoints,
		&player.Characteristics.PassingAccuracyPoints,
		&player.Characteristics.DefensiveInvolvements)

	if err != nil {
		return nil, err
	}

	return &player, nil
}

func GetBestDefender() (*Player, error) {
	sql := `
	SELECT p.id, p.name, p.jersey_number, c.goals, c.speed_points, c.assist_points, c.passing_accuracy_points, c.defensive_involvements
	FROM players p, player_characteristics c
	WHERE p.id = c.id
	ORDER BY defensive_involvements DESC
	LIMIT 1
`

	row := db.DB.QueryRow(sql)

	var player Player
	player.setCharacteristics(NewPlayerCharacteristics())

	err := row.Scan(
		&player.ID,
		&player.Name,
		&player.JerseyNumber,
		&player.Characteristics.Goals,
		&player.Characteristics.SpeedPoints,
		&player.Characteristics.AssistPoints,
		&player.Characteristics.PassingAccuracyPoints,
		&player.Characteristics.DefensiveInvolvements)

	if err != nil {
		return nil, err
	}

	return &player, nil
}

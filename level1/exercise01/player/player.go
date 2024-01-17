package player

import (
	"example.com/manchester/db"
)

type Player struct {
	ID              int64
	Name            string
	JerseyNumber    int64
	Characteristics *PlayerCharacteristics
}

func (p *Player) setCharacteristics(c *PlayerCharacteristics) {
	p.Characteristics = c
}

func GetByJersey(jerseyNumber int64) (*Player, error) {
	sql := `
		SELECT p.id, p.name, p.jersey_number, c.goals, c.speed_points, c.assist_points, c.passing_accuracy_points, c.defensive_involvements 
		FROM players p, player_characteristics c 
		WHERE p.jersey_number = ? and p.id = c.id
	`

	row := db.DB.QueryRow(sql, jerseyNumber)

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

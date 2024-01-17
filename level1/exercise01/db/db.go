package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "manchester.db")

	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = createTables()
	if err != nil {
		return err
	}

	err = insertInitialData()
	if err != nil {
		return err
	}

	return nil
}

func createTables() error {
	err := createPlayerTable()

	if err != nil {
		return err
	}

	err = createCharacteristicsTable()

	if err != nil {
		return err
	}

	return nil
}

func createPlayerTable() error {
	sql := `
	CREATE TABLE IF NOT EXISTS players (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		jersey_number INTEGER NOT NULL		
	)
	`

	_, err := DB.Exec(sql)

	if err != nil {
		return err
	}

	return nil

}

func createCharacteristicsTable() error {
	sql := `
	CREATE TABLE IF NOT EXISTS player_characteristics (
		id INTEGER PRIMARY KEY,
		goals INTEGER,
		speed_points INTEGER,
		assist_points INTEGER,
		passing_accuracy_points INTEGER,
		defensive_involvements INTEGER
	)
	`

	_, err := DB.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}

func insertInitialData() error {
	sqlStatements := []string{
		"DELETE FROM players;",
		"INSERT INTO players (id, name, jersey_number) VALUES (1, 'Bruno Fernandes', 8);",
		"INSERT INTO players (id, name, jersey_number) VALUES (2, 'Rasmus Hojlund', 11);",
		"INSERT INTO players (id, name, jersey_number) VALUES (3, 'Harry Maguire', 5);",
		"INSERT INTO players (id, name, jersey_number) VALUES (4, 'Alejandro Garnacho', 17);",
		"INSERT INTO players (id, name, jersey_number) VALUES (5, 'Mason Mount', 7);",
		"DELETE FROM player_characteristics;",
		"INSERT INTO player_characteristics (id, goals, speed_points, assist_points, passing_accuracy_points, defensive_involvements) VALUES (1, 5, 6, 9, 10, 3);",
		"INSERT INTO player_characteristics (id, goals, speed_points, assist_points, passing_accuracy_points, defensive_involvements) VALUES (2, 12, 8, 2, 6, 2);",
		"INSERT INTO player_characteristics (id, goals, speed_points, assist_points, passing_accuracy_points, defensive_involvements) VALUES (3, 1, 5, 1, 7, 9);",
		"INSERT INTO player_characteristics (id, goals, speed_points, assist_points, passing_accuracy_points, defensive_involvements) VALUES (4, 8, 7, 8, 6, 0);",
		"INSERT INTO player_characteristics (id, goals, speed_points, assist_points, passing_accuracy_points, defensive_involvements) VALUES (5, 2, 6, 4, 8, 1);",
	}
	
	for _, statement := range sqlStatements {
		_, err := DB.Exec(statement)
		if err != nil {
			return err
		}
	}

	return nil
}
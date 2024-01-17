package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "travel.db")

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
	sql := `
	CREATE TABLE IF NOT EXISTS destinations (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS seasons (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS season_cost (
		season_id INTEGER PRIMARY KEY,
		cost INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS activities (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS destination_activities (
		destination_id INTEGER,
		activity_id INTEGER
	);

	CREATE TABLE IF NOT EXISTS season_destinations (
		season_id INTEGER,
		destination_id INTEGER
	);
	`

	_, err := DB.Exec(sql)

	if err != nil {
		return err
	}

	return nil

}

func insertInitialData() error {
	sqlStatements := []string{
		"DELETE FROM destinations;",
		"INSERT INTO destinations (id, name) VALUES (1, 'Andorra');",
		"INSERT INTO destinations (id, name) VALUES (2, 'Switzerland');",
		"INSERT INTO destinations (id, name) VALUES (3, 'Spain');",
		"INSERT INTO destinations (id, name) VALUES (4, 'Portugal');",
		"INSERT INTO destinations (id, name) VALUES (5, 'France');",
		"INSERT INTO destinations (id, name) VALUES (6, 'Italy');",
		"INSERT INTO destinations (id, name) VALUES (7, 'Belgium');",
		"INSERT INTO destinations (id, name) VALUES (8, 'Austria');",
		"DELETE FROM seasons;",
		"INSERT INTO seasons (id, name) VALUES (1, 'Winter');",
		"INSERT INTO seasons (id, name) VALUES (2, 'Summer');",
		"INSERT INTO seasons (id, name) VALUES (3, 'Spring');",
		"INSERT INTO seasons (id, name) VALUES (4, 'Autumn');",
		"DELETE FROM season_cost;",
		"INSERT INTO season_cost (season_id, cost) VALUES (1, 100);",
		"INSERT INTO season_cost (season_id, cost) VALUES (2, 400);",
		"INSERT INTO season_cost (season_id, cost) VALUES (3, 300);",
		"INSERT INTO season_cost (season_id, cost) VALUES (4, 200);",
		"DELETE FROM activities;",
		"INSERT INTO activities (id, name) VALUES (1, 'Skiing');",
		"INSERT INTO activities (id, name) VALUES (2, 'Hiking');",
		"INSERT INTO activities (id, name) VALUES (3, 'Extreme sports');",
		"INSERT INTO activities (id, name) VALUES (4, 'Beach');",
		"INSERT INTO activities (id, name) VALUES (5, 'Nature tours');",
		"INSERT INTO activities (id, name) VALUES (6, 'Cultural tours');",
		"INSERT INTO activities (id, name) VALUES (7, 'Historical tours');",
		"DELETE FROM destination_activities;",
		"INSERT INTO destination_activities (destination_id, activity_id) VALUES (1, 1);",
		"INSERT INTO destination_activities (destination_id, activity_id) VALUES (2, 5);",
		"INSERT INTO destination_activities (destination_id, activity_id) VALUES (3, 2);",
		"INSERT INTO destination_activities (destination_id, activity_id) VALUES (3, 3);",
		"INSERT INTO destination_activities (destination_id, activity_id) VALUES (4, 4);",
		"INSERT INTO destination_activities (destination_id, activity_id) VALUES (5, 3);",
		"INSERT INTO destination_activities (destination_id, activity_id) VALUES (6, 6);",
		"INSERT INTO destination_activities (destination_id, activity_id) VALUES (6, 7);",
		"INSERT INTO destination_activities (destination_id, activity_id) VALUES (7, 2);",
		"INSERT INTO destination_activities (destination_id, activity_id) VALUES (7, 3);",
		"INSERT INTO destination_activities (destination_id, activity_id) VALUES (8, 6);",
		"INSERT INTO destination_activities (destination_id, activity_id) VALUES (8, 7);",
		"DELETE FROM season_destinations;",
		"INSERT INTO season_destinations (season_id, destination_id) VALUES (1, 1);",
		"INSERT INTO season_destinations (season_id, destination_id) VALUES (1, 2);",
		"INSERT INTO season_destinations (season_id, destination_id) VALUES (2, 3);",
		"INSERT INTO season_destinations (season_id, destination_id) VALUES (2, 4);",
		"INSERT INTO season_destinations (season_id, destination_id) VALUES (3, 5);",
		"INSERT INTO season_destinations (season_id, destination_id) VALUES (3, 6);",
		"INSERT INTO season_destinations (season_id, destination_id) VALUES (4, 7);",
		"INSERT INTO season_destinations (season_id, destination_id) VALUES (4, 8);",
	}

	for _, statement := range sqlStatements {
		_, err := DB.Exec(statement)
		if err != nil {
			return err
		}
	}

	return nil
}

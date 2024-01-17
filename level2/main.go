package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var teams map[string]int
var matchupCounts map[string]map[string]int

type TeamPoints struct {
	Team   string
	Points int
}

type MatchResult struct {
	Team1   string
	Team2   string
	Score1  int
	Score2  int
	Winner  string
	IsDraw  bool
	Points1 int
	Points2 int
}

func main() {
	InitializeTeamData()
	InitializeMatchupCounts()
	SeedNumberGenerator()
	PlayMatches()
	DisplayStandings()
}

func InitializeTeamData() {
	teams = make(map[string]int)

	teams["Manchester United"] = 0
	teams["Liverpool"] = 0
	teams["Arsenal"] = 0
	teams["Chelsea"] = 0
	teams["Manchester City"] = 0
	teams["Tottenham Hotspur"] = 0
}

func InitializeMatchupCounts() {
	matchupCounts = make(map[string]map[string]int)

	for team := range teams {
		matchupCounts[team] = make(map[string]int)
	}
}

func GetTeamsOrderedByPoints() []TeamPoints {
	var teamPoints []TeamPoints

	for team, points := range teams {
		teamPoints = append(teamPoints, TeamPoints{team, points})
	}

	sort.Slice(teamPoints, func(i, j int) bool {
		return teamPoints[i].Points > teamPoints[j].Points
	})

	return teamPoints
}

func DisplayStandings() {
	orderedTeams := GetTeamsOrderedByPoints()

	fmt.Printf("\nStandings:\n\n")

	for _, t := range orderedTeams {
		fmt.Printf("%s: %d\n", t.Team, t.Points)
	}
}

func SeedNumberGenerator() {
	rand.Seed(time.Now().UnixNano())
}

func GetMatchResult(team1 string, team2 string) *MatchResult {
	team1Score := rand.Intn(5)
	team2Score := rand.Intn(5)
	isDraw := team1Score == team2Score

	var winner string
	var points1, points2 int

	if isDraw {
		winner = ""
		points1 = 1
		points2 = 1
	} else if team1Score > team2Score {
		winner = team1
		points1 = 3
		points2 = 0
	} else {
		winner = team2
		points1 = 0
		points2 = 3
	}

	return &MatchResult{
		Team1:   team1,
		Team2:   team2,
		Score1:  team1Score,
		Score2:  team2Score,
		Winner:  winner,
		IsDraw:  isDraw,
		Points1: points1,
		Points2: points2,
	}
}

func PlayMatches() {
	fmt.Printf("Results:\n\n")

	for team1, _ := range teams {
		for team2, _ := range teams {
			if team1 != team2 {
				for i := 0; i < 3; i++ {
					if matchupCounts[team1][team2] < 3 {

						result := GetMatchResult(team1, team2)

						teams[team1] += result.Points1
						teams[team2] += result.Points2

						matchupCounts[team1][team2] += 1
						matchupCounts[team2][team1] += 1

						fmt.Printf("%s %d - %d %s\n", team1, result.Score1, result.Score2, team2)
					}
				}
			}
		}
	}
}

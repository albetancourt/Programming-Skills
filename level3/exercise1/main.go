package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type player struct {
	name       string
	powerLevel int
}

const goalkeeper = "goalkeeper"
const defender = "defender"
const midfielder = "midfielder"
const forward = "forward"

var positions = []string{
	goalkeeper,
	defender,
	midfielder,
	forward,
}

var playersByPosition = map[string]int{
	goalkeeper: 1,
	defender:   4,
	midfielder: 3,
	forward:    3,
}

var tottenhamPlayers = map[string][]player{
	goalkeeper: {
		{name: "Hugo Lloris", powerLevel: 85},
		{name: "Guglielmo Vicario", powerLevel: 79},
		{name: "Fraser Forster", powerLevel: 79},
		{name: "Brandon Austin", powerLevel: 79},
	},
	defender: {
		{name: "Eric Dier", powerLevel: 80},
		{name: "Cristian Romero", powerLevel: 80},
		{name: "Davinson Sánchez", powerLevel: 74},
		{name: "Japhet Tanganga", powerLevel: 70},
		{name: "Matt Doherty", powerLevel: 70},
		{name: "Djed Spence", powerLevel: 70},
		{name: "Sergio Reguilón", powerLevel: 74},
		{name: "Ben Davies", powerLevel: 76},
		{name: "Joe Rodon", powerLevel: 70},
		{name: "Mislav Orsic", powerLevel: 71},
	},
	midfielder: {
		{name: "Oliver Skipp", powerLevel: 70},
		{name: "Pierre-Emile Højbjerg", powerLevel: 70},
		{name: "Yves Bissouma", powerLevel: 72},
		{name: "James Maddison", powerLevel: 74},
		{name: "Giovani Lo Celso", powerLevel: 78},
		{name: "Ryan Sessegnon", powerLevel: 80},
		{name: "Dejan Kulusevski", powerLevel: 60},
		{name: "Pape Sarr", powerLevel: 65},
		{name: "Rodrigo Bentancur", powerLevel: 65},
		{name: "Oliver Skipp", powerLevel: 65},
	},
	forward: {
		{name: "Son Heung-Min", powerLevel: 78},
		{name: "Richarlison", powerLevel: 82},
		{name: "Bryan Gil", powerLevel: 80},
		{name: "Timo Werner", powerLevel: 82},
		{name: "Brennan Johnson", powerLevel: 70},
		{name: "Manor Solomon", powerLevel: 70},
		{name: "Alejo Veliz", powerLevel: 75},
		{name: "Dane Scarlett", powerLevel: 75},
	},
}

var manchesterPlayers = map[string][]player{
	goalkeeper: {
		{name: "Andre Onana", powerLevel: 80},
		{name: "Tom Heaton", powerLevel: 75},
		{name: "Altay Bayindir", powerLevel: 69},
	},
	defender: {
		{name: "Victor Lindelöf", powerLevel: 80},
		{name: "Harry Maguire", powerLevel: 82},
		{name: "Lisandro Martinez", powerLevel: 82},
		{name: "Tyrell Malacia", powerLevel: 67},
		{name: "Raphaël Varane", powerLevel: 80},
		{name: "Diogo Dalot", powerLevel: 89},
		{name: "Luke Shaw", powerLevel: 89},
		{name: "Aaron Wan-Bissaka", powerLevel: 70},
	},
	midfielder: {
		{name: "Sofyan Amrabat", powerLevel: 76},
		{name: "Scott McTominay", powerLevel: 80},
		{name: "Bruno Fernandes", powerLevel: 88},
		{name: "Christian Eriksen", powerLevel: 67},
		{name: "Mason Mount", powerLevel: 77},
		{name: "Kobbie Mainoo", powerLevel: 65},
		{name: "Daniel Gore", powerLevel: 60},
	},
	forward: {
		{name: "Anthony Martial", powerLevel: 50},
		{name: "Marcus Rashford", powerLevel: 76},
		{name: "Antony", powerLevel: 75},
		{name: "Rasmus Højlund", powerLevel: 80},
		{name: "Alejandro Garnacho", powerLevel: 85},
		{name: "Facundo Pellistri", powerLevel: 75},
	},
}

func main() {
	SetRandomSeed()

	selectedTottenhamPlayers := selectTottenhamPlayers()
	selectedManchesterPlayers := selectManchesterPlayers()

	showPlayers(selectedTottenhamPlayers, "Selected Tottenham players:")
	showPlayers(selectedManchesterPlayers, "Selected Manchester players:")
}

func showPlayers(players map[string][]player, header string) {
	fmt.Println(header)

	for _, position := range positions {
		for _, player := range players[position] {
			fmt.Printf("%s - Power Level: %d - Position: %s\n", player.name, player.powerLevel, position)
		}
	}

	fmt.Printf("\n\n")
}

func selectTottenhamPlayers() map[string][]player {
	players := make(map[string][]player)

	for _, p := range positions {
		players[p] = getNRandomPlayers(tottenhamPlayers[p], playersByPosition[p])
	}

	return players
}

func selectManchesterPlayers() map[string][]player {
	players := make(map[string][]player)

	for _, p := range positions {
		players[p] = getBestNPlayers(manchesterPlayers[p], playersByPosition[p])
	}

	return players
}

func sortByPowerLevel(players []player) {
	sort.Slice(players, func(i, j int) bool {
		return players[i].powerLevel > players[j].powerLevel
	})
}

func getBestNPlayers(players []player, n int) []player {
	sortByPowerLevel(players)

	return players[:n]
}

func getNRandomPlayers(players []player, n int) []player {
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	return players[:n]
}

func SetRandomSeed() {
	rand.Seed(time.Now().UnixNano())
}

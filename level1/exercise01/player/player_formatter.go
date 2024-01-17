package player

import (
	"fmt"
)

func GetCharacteristics(player *Player) string {
	s := fmt.Sprintf("\nName: %s\nJersey number: %d\nGoals: %d\nSpeed points: %d\nAssist points: %d\nPassing accuracy points: %d\nDefensive involvements: %d\n\n",
		player.Name,
		player.JerseyNumber,
		player.Characteristics.Goals,
		player.Characteristics.SpeedPoints,
		player.Characteristics.AssistPoints,
		player.Characteristics.PassingAccuracyPoints,
		player.Characteristics.DefensiveInvolvements)

	return s
}

func GetComparison(p1 *Player, p2 *Player) string {
	s := fmt.Sprintf("\nName: %s %s\nGoals: %d %d\nSpeed points: %d %d\nAssist points: %d %d\nPassing accuracy points: %d %d\nDefensive involvements: %d %d\n\n",
		p1.Name, p2.Name,
		p1.Characteristics.Goals, p2.Characteristics.Goals,
		p1.Characteristics.SpeedPoints, p2.Characteristics.SpeedPoints,
		p1.Characteristics.AssistPoints, p2.Characteristics.AssistPoints,
		p1.Characteristics.PassingAccuracyPoints, p2.Characteristics.PassingAccuracyPoints,
		p1.Characteristics.DefensiveInvolvements, p2.Characteristics.DefensiveInvolvements)

	return s
}

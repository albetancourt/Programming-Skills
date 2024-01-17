package player

type PlayerCharacteristics struct {
	Goals                 int64
	SpeedPoints           int64
	AssistPoints          int64
	PassingAccuracyPoints int64
	DefensiveInvolvements int64
}

func NewPlayerCharacteristics() *PlayerCharacteristics {
	return &PlayerCharacteristics{
		Goals: 0,
		SpeedPoints: 0,
		AssistPoints: 0,
		PassingAccuracyPoints: 0,
		DefensiveInvolvements: 0,
	}
}

package player

// PlayerDetails defines playername and playermark
type PlayerDetails struct {
	playerName string
	playerMark string
}

// SetPlayerName used to set playername
func (p *PlayerDetails) SetPlayerName(playerName string) {
	p.playerName = playerName
}

// SetPlayerName used to set playermark
func (p *PlayerDetails) SetPlayerMark(playerMark string) {
	p.playerMark = playerMark
}

// PlayerName returns playername
func (p *PlayerDetails) PlayerName() string {
	return p.playerName
}

// PlayerMark return player mark
func (p *PlayerDetails) PlayerMark() string {
	return p.playerMark
}

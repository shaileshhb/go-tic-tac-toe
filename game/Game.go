package game

import (
	"github.com/shaileshhb/go-tic-tac-toe/board"
	"github.com/shaileshhb/go-tic-tac-toe/player"
	"github.com/shaileshhb/go-tic-tac-toe/resultanalyzer"
)

var currentPlayer = 0
var nextPlayer = 1
var flag = 0

// Play starts playing the game
func (p *GameStructure) Play(location int) (bool, error) {

	inserted, err := p.board.AddMarkFromBoard(p.CurrentPlayer().PlayerMark(), location)
	if err != nil {
		return false, err
	}
	if inserted == true {
		if flag == 0 {
			currentPlayer++
			nextPlayer--
			flag++
		} else {
			currentPlayer--
			nextPlayer++
			flag--
		}
	}

	return true, nil
}

func (p *GameStructure) CurrentPlayer() *player.PlayerDetails {
	return &(*p.players)[currentPlayer]
}

func (p *GameStructure) NextPlayer() *player.PlayerDetails {
	return &(*p.players)[nextPlayer]
}

// StartGame initailizes the contents required to play the game
func StartGame(player *[]player.PlayerDetails, board *board.GameBoard, resultanalyzer *resultanalyzer.AnalyzeResult) *GameStructure {

	return &GameStructure{
		players:        player,
		board:          board,
		resultAnalyzer: resultanalyzer,
	}

}

// GetGameStatus returns current result as string value
func (p *GameStructure) GetGameStatus() string {
	return p.resultAnalyzer.CheckWinner()
}

// GameStructure declares a struct required to intitalize the game
type GameStructure struct {
	players        *[]player.PlayerDetails
	board          *board.GameBoard
	resultAnalyzer *resultanalyzer.AnalyzeResult
}

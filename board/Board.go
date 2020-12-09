package board

import (
	"strconv"

	"github.com/shaileshhb/go-tic-tac-toe/cell"
)

//CreateBoard creates the board according to the given size
func CreateBoard(boardSize int) *GameBoard {

	board := make([]string, boardSize*boardSize)
	for cell := range board {
		board[cell] = "-"
	}

	return &GameBoard{
		board:     board,
		boardSize: boardSize,
	}
}

// GameBoard returns the created board
func (g *GameBoard) GameBoard() []string {
	return g.board
}

// BoardSize returns the board size
func (g *GameBoard) BoardSize() int {
	return g.boardSize
}

// IsBoardFull returns false if it is not empty i.e. there exists "-" in board and return true if board is full
func (g *GameBoard) IsBoardFull() bool {

	for i := 0; i < len(g.board); i++ {
		if g.board[i] == "-" {
			return false
		}
	}
	return true
}

// AddMarkFromBoard checks if board is full and if not put the mark at said location
func (g *GameBoard) AddMarkFromBoard(playerMark string, location int) (bool, error) {

	if location >= g.boardSize*g.boardSize || location < 0 {
		str := "Please enter a cell number between 0 and" + strconv.Itoa(g.boardSize*g.boardSize-1)
		return false, &outOfCellError{cellOutOfBound: str}
	}

	if g.IsBoardFull() == false {
		err := cell.MarkCellIfEmpty(&g.board, location, playerMark)
		if err != nil {
			return false, err
		}
	}

	return true, nil

}

func (o *outOfCellError) Error() string {
	return o.cellOutOfBound
}

// GameBoard defines the board
type GameBoard struct {
	boardSize int
	board     []string
}

type outOfCellError struct {
	cellOutOfBound string
}

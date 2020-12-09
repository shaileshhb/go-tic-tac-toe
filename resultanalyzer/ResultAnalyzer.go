package resultanalyzer

import (
	"github.com/shaileshhb/go-tic-tac-toe/board"
	"github.com/shaileshhb/go-tic-tac-toe/result"
)

// SetCurrentGameBoard sets the current result and current board
func SetCurrentGameBoard(board *board.GameBoard) *AnalyzeResult {

	return &AnalyzeResult{
		CurrentResult: result.Progress,
		GameBoard:     board,
	}

}

// CheckWinner returns the winner, draw or progress state
func (a *AnalyzeResult) CheckWinner() string {

	if a.checkRow() == result.Win {
		return result.Win
	}
	if a.checkColumn() == result.Win {
		return result.Win
	}
	if a.checkDiagonal() == result.Win {
		return result.Win
	}
	if a.GameBoard.IsBoardFull() == true {
		return result.Draw
	}

	return result.Progress
}

func (a *AnalyzeResult) checkRow() string {

	j := 0
	resultCount := 1
	a.CurrentResult = result.Progress
	boardSize := a.GameBoard.BoardSize()
	currentGameBoard := a.GameBoard.GameBoard()

	for i := 1; i < len(currentGameBoard); i++ {
		if currentGameBoard[j] != "-" && currentGameBoard[j] == currentGameBoard[i] {
			resultCount++
			if resultCount == boardSize {
				a.CurrentResult = result.Win
				break
			}
		} else {
			i = j + boardSize
			j = j + boardSize
			resultCount = 1
			a.CurrentResult = result.Progress
		}
	}

	return a.CurrentResult

}

func (a *AnalyzeResult) checkColumn() string {

	j := 0
	resultCount := 1
	a.CurrentResult = result.Progress
	boardSize := a.GameBoard.BoardSize()
	currentGameBoard := a.GameBoard.GameBoard()

	for i := boardSize; j < len(currentGameBoard)/boardSize; i += boardSize {

		if currentGameBoard[j] != "-" && currentGameBoard[j] == currentGameBoard[i] {
			resultCount++
			if resultCount == boardSize {
				a.CurrentResult = result.Win
				break
			}
		} else {
			j++
			i = j
			resultCount = 1
			a.CurrentResult = result.Progress
		}
	}

	return a.CurrentResult

}

func (a *AnalyzeResult) checkDiagonal() string {

	j := 0
	resultCount := 1
	resultFound := false
	a.CurrentResult = result.Progress
	boardSize := a.GameBoard.BoardSize()
	currentGameBoard := a.GameBoard.GameBoard()

	for i := (boardSize + 1); i < len(currentGameBoard); i += boardSize + 1 {

		if currentGameBoard[j] != "-" && currentGameBoard[j] == currentGameBoard[i] {
			resultCount++
			if resultCount == boardSize {
				a.CurrentResult = result.Win
				resultFound = true
				break
			}
		} else {
			resultCount = 1
			break
		}
	}

	j = boardSize - 1

	for i := (j * 2); i < len(currentGameBoard)-j; i += j {

		if resultFound == true {
			break
		}

		if currentGameBoard[j] != "-" && currentGameBoard[j] == currentGameBoard[i] {
			resultCount++
			if resultCount == boardSize {
				a.CurrentResult = result.Win
				break
			}
		}
	}

	return a.CurrentResult

}

func (a *AnalyzeResult) GetStatus() string {
	return a.CurrentResult
}

// AnalyzeResult consists of CurrentResult and board
type AnalyzeResult struct {
	CurrentResult string
	*board.GameBoard
}

package main

import (
	"fmt"

	"github.com/shaileshhb/go-tic-tac-toe/board"
	"github.com/shaileshhb/go-tic-tac-toe/game"
	"github.com/shaileshhb/go-tic-tac-toe/mark"
	"github.com/shaileshhb/go-tic-tac-toe/player"
	"github.com/shaileshhb/go-tic-tac-toe/resultanalyzer"
)

func main() {
	fmt.Println("Main.go")

	var location int
	var userBoardSize int
	var playerName string

	fmt.Print("Enter Board size:")
	fmt.Scan(&userBoardSize)

	var players []player.PlayerDetails
	gameboard := board.CreateBoard(userBoardSize)
	resultanalyzer := resultanalyzer.SetCurrentGameBoard(gameboard)

	playGame := game.StartGame(&players, gameboard, resultanalyzer)

	playerName = getFirstDataFromUser(&players)
	for true {
		fmt.Print(playerName, " enter cell number:")
		fmt.Scan(&location)
		inserted, err := playGame.Play(location)
		if inserted {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
	}
	printGameBoard(gameboard)

	playerName = getFirstDataFromUser(&players)
	for true {
		fmt.Print(playerName, " enter cell number:")
		fmt.Scan(&location)
		inserted, err := playGame.Play(location)
		if inserted {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
	}
	printGameBoard(gameboard)

	for playGame.GetGameStatus() == "Progress" {

		fmt.Println(playGame.GetGameStatus())

		fmt.Print(playGame.CurrentPlayer().PlayerName(), " enter cell number:")
		fmt.Scan(&location)

		playGame.Play(location)

		printGameBoard(gameboard)
	}

	if playGame.GetGameStatus() == "Win" {
		fmt.Println(playGame.NextPlayer().PlayerName(), "is Winner!!")
	} else if playGame.GetGameStatus() == "Draw" {
		fmt.Println("Its a Draw")
	}

}

func getFirstDataFromUser(players *[]player.PlayerDetails) string {

	var player player.PlayerDetails
	var name string

	fmt.Print("Enter Your Name:")
	fmt.Scan(&name)

	player.SetPlayerName(name)

	if len(*players) == 0 {
		player.SetPlayerMark(mark.X)
	} else {
		player.SetPlayerMark(mark.O)
	}

	*players = append(*players, player)

	return name
}

func printGameBoard(g *board.GameBoard) {

	board := g.GameBoard()
	// fmt.Println(g.BoardSize())

	j := 1
	for i := 0; i < len(board); i++ {
		if j <= g.BoardSize() {
			fmt.Print("| " + board[i] + " ")
			j++
		} else {
			j = 2
			fmt.Println("|")
			fmt.Print("| " + board[i] + " ")
		}
	}
	fmt.Println("|")
}

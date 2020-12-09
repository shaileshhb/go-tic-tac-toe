package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print("Enter Board size:")
		value, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Please enter a integer for board size")
		} else {
			value = strings.TrimSpace(value)
			userBoardSize, err = strconv.Atoi(value)
			if err != nil {
				fmt.Println("Please enter a integer for board size")
			} else if userBoardSize <= 2 {
				fmt.Println("Board size should be greater than 2")
			} else {
				break
			}

		}
	}

	var players []player.PlayerDetails
	gameboard := board.CreateBoard(userBoardSize)
	resultanalyzer := resultanalyzer.SetCurrentGameBoard(gameboard)

	playGame := game.StartGame(&players, gameboard, resultanalyzer)

	playerName = getFirstDataFromUser(&players)
	for true {
		fmt.Print(playerName, " enter cell number:")

		value, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Please enter a valid number")
		}
		value = strings.TrimSpace(value)
		location, err = strconv.Atoi(value)
		if err != nil {
			fmt.Println("Please enter a valid number")
		} else {
			inserted, err := playGame.Play(location)
			if err != nil {
				fmt.Println(err)
			}
			if inserted {
				break
			}

		}

	}
	printGameBoard(gameboard)

	playerName = getFirstDataFromUser(&players)
	for true {
		fmt.Print(playerName, " enter cell number:")

		value, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Please enter a valid number")
		}
		value = strings.TrimSpace(value)
		location, err = strconv.Atoi(value)
		if err != nil {
			fmt.Println("Please enter a valid number")
		} else {
			inserted, err := playGame.Play(location)
			if err != nil {
				fmt.Println(err)
			}
			if inserted {
				break
			}
		}

	}
	printGameBoard(gameboard)

	for playGame.GetGameStatus() == "Progress" {

		fmt.Print(playGame.CurrentPlayer().PlayerName(), " enter cell number:")

		value, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Please enter a valid number")
		}
		value = strings.TrimSpace(value)
		location, err = strconv.Atoi(value)
		if err != nil {
			fmt.Println("Please enter a valid number")
		} else {
			_, err = playGame.Play(location)
			if err != nil {
				fmt.Println(err)
			} else {
				printGameBoard(gameboard)
			}
		}

	}

	if playGame.GetGameStatus() == "Win" {
		fmt.Println(playGame.NextPlayer().PlayerName(), "is Winner!!")
	} else if playGame.GetGameStatus() == "Draw" {
		fmt.Println("Its a Draw")
	}

}

func checkName(name string) bool {

	for _, str := range name {
		if (str < 'a' || str > 'z') && (str < 'A' || str > 'Z') {
			return false
		}
	}
	return true
}

func getFirstDataFromUser(players *[]player.PlayerDetails) string {

	var player player.PlayerDetails
	var name string
	reader := bufio.NewReader(os.Stdin)

	for true {

		fmt.Print("Enter Your Name:")
		value, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Please enter a valid name")
		}
		name = strings.TrimSpace(value)
		validName := checkName(name)
		if !validName {
			fmt.Println("Name should only contain alphabets")
		} else {
			if len(*players) == 0 {
				player.SetPlayerName(name)
				player.SetPlayerMark(mark.X)
				break
			} else {
				if strings.ToLower((*players)[0].PlayerName()) != strings.ToLower(name) {
					player.SetPlayerName(name)
					player.SetPlayerMark(mark.O)
					break
				} else {
					fmt.Println("Both players cannot have same name")
				}
			}
		}
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

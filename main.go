package main

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

func promptField(board [9]int) [9]int {
	var field int
	fmt.Print("Choose field: ")
	fmt.Scanln(&field)

	if field < 1 || field > 9 {
		fmt.Println(Red("Error"), "Choose a valid field")
		board = promptField(board)
	} else {
		if board[field-1] == 0 {
			board[field-1] = 1
		} else {
			fmt.Println(Red("Error"), "Cannot place on occupied field")
			board = promptField(board)
		}
	}

	return board
}

func showBoard(board [9]int, message ...string) {
	if message != nil {
		fmt.Println(message[0])
	}

	for i, value := range board {
		if value == 0 {
			fmt.Print(Faint(i + 1))
		} else if value == 1 {
			fmt.Print(Green("X"))
		} else {
			fmt.Print(Red("O"))
		}

		if i > 0 && (i+1)%3 == 0 {
			fmt.Print("\n")
		} else {
			fmt.Print("|")
		}
	}
}

func checkBoard(board [9]int) int {
	sums := [8]int{}
	sums[0] = board[0] + board[1] + board[2]
	sums[1] = board[3] + board[4] + board[5]
	sums[2] = board[6] + board[7] + board[8]
	sums[3] = board[0] + board[3] + board[6]
	sums[4] = board[1] + board[4] + board[7]
	sums[5] = board[2] + board[5] + board[8]
	sums[6] = board[0] + board[4] + board[8]
	sums[7] = board[2] + board[4] + board[6]

	for _, value := range sums {
		if value == 3 {
			return 1
		} else if value == 30 {
			return 2
		}
	}

	return 0
}

func predict(board [9]int, computer bool, predicting bool, field int) (int, int) {
	if predicting {
		if checkBoard(board) == 1 {
			return field, -10
		} else if checkBoard(board) == 2 {
			return field, 10
		}
	}

	empty := make([]int, 0)

	for i, value := range board {
		if value == 0 {
			empty = append(empty, i)
		}
	}

	newboard := board

	var bestField int
	var bestScore int

	for _, field := range empty {
		if computer {
			newboard[field] = 10
		} else {
			newboard[field] = 1
		}

		predField, predScore := predict(newboard, !computer, true, field)

		if predScore >= bestScore {
			bestField = predField
			bestScore = predScore
		}

		newboard[field] = 0
	}

	return bestField, bestScore
}

func bestField(board [9]int) int {
	field, _ := predict(board, true, false, 0)
	return field
}

func main() {
	turn := 1
	board := [9]int{}
	score := 0
	fmt.Println(Bold("Tic tac toe"))

	for score == 0 {
		showBoard(board)

		if turn%2 == 0 {
			fmt.Println("It's", Bold(Red("computer's")), "turn")
			board[bestField(board)] = 10
		} else {
			fmt.Println("It's", Bold(Green("player's")), "turn")
			board = promptField(board)
		}

		score = checkBoard(board)

		switch score {
		case 1:
			showBoard(board)
			fmt.Println(Bold(Green("Player")), "won!")
		case 2:
			showBoard(board)
			fmt.Println(Bold(Red("Computer")), "won!")
		}
		turn++
	}
}

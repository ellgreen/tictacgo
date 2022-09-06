package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const AIBattle = false

func main() {
	if !AIBattle {
		PrintExample()
	}

	var board Board
	turn := 'X'

	for i := 0; i <= 9; i++ {
		if result := board.Check(); result != Incomplete {
			fmt.Printf("Result: %s\n", HumanResult[result])
			break
		}

		move := board.BestMove(turn)

		if !AIBattle && turn == 'X' {
			move = GetInput(board)
		}

		board.Play(move, turn)

		board.Render()

		turn = NextTurn(turn)
	}
}

func GetInput(board Board) int {
	reader := bufio.NewReader(os.Stdin)
	moves := board.AvailablesMoves()

	for {
		fmt.Print("Enter move: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			panic("Error reading console input")
		}

		input = strings.Replace(input, "\n", "", -1)

		inputMove, err := strconv.Atoi(input)

		if err != nil {
			panic("Error turning input move into integer")
		}

		for _, move := range moves {
			if inputMove == move {
				return move
			}
		}

		fmt.Printf("Invalid move entered: %d\n", inputMove)
	}
}

func PrintExample() {
	board := Board{
		State: [9]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8'},
	}

	board.Render()
}

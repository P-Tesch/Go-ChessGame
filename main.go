package main

import (
	"fmt"
)

func main() {
	board := new(Board)
	board.BuildBoard()

	fmt.Println(board.BoardArray[0])
	fmt.Println(board.BoardArray[1])
	fmt.Println(board.BoardArray[2])
	fmt.Println(board.BoardArray[3])
	fmt.Println(board.BoardArray[4])
	fmt.Println(board.BoardArray[5])
	fmt.Println(board.BoardArray[6])
	fmt.Println(board.BoardArray[7])

	fmt.Println()

	possibleMoves1 := board.BoardArray[0][0].PossibleMoves()
	fmt.Println(possibleMoves1[0])
	fmt.Println(possibleMoves1[1])
	fmt.Println(possibleMoves1[2])
	fmt.Println(possibleMoves1[3])
	fmt.Println(possibleMoves1[4])
	fmt.Println(possibleMoves1[5])
	fmt.Println(possibleMoves1[6])
	fmt.Println(possibleMoves1[7])

	fmt.Println()

	possibleMoves2 := board.BoardArray[7][0].PossibleMoves()
	fmt.Println(possibleMoves2[0])
	fmt.Println(possibleMoves2[1])
	fmt.Println(possibleMoves2[2])
	fmt.Println(possibleMoves2[3])
	fmt.Println(possibleMoves2[4])
	fmt.Println(possibleMoves2[5])
	fmt.Println(possibleMoves2[6])
	fmt.Println(possibleMoves2[7])
}

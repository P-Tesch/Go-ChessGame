package main

import (
	"fmt"
)

func main() {
	board := new(Board)
	board.BuildBoard()

	fmt.Println(board.ToString())
}

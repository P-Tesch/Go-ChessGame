package main

import (
	"fmt"
)

func main() {
	position := Position{0, 0, 0, "a"}

	position.SetBoardPosition(3, 3)
	fmt.Println(position.GetBoardPosition())
	fmt.Println(position.GetChessPostion())

	position.SetChessPosition(8, "h")
	fmt.Println(position.GetBoardPosition())
	fmt.Println(position.GetChessPostion())
}

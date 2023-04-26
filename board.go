package main

type Board struct {
	BoardArray [8][8]Piece
}

func (board *Board) BuildBoard() {
	board.BoardArray[0][0] = NewRook(board, NewPosition(0, 0), NewColor(false))
	board.BoardArray[7][0] = NewRook(board, NewPosition(7, 0), NewColor(true))
	board.BoardArray[0][7] = NewRook(board, NewPosition(0, 0), NewColor(false))
	board.BoardArray[7][7] = NewRook(board, NewPosition(0, 0), NewColor(true))
}

func (board *Board) ToString() string {
	var boardString string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			piece := board.BoardArray[i][j]
			if piece == nil {
				boardString += "- "
			} else {
				boardString += piece.ToString() + " "
			}
		}
		boardString += "\n"
	}
	return boardString
}

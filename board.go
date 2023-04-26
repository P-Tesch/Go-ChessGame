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

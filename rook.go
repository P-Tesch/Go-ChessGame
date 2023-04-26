package main

type Rook struct {
	board   *Board
	positon *Position
	color   *Color
}

func NewRook(board *Board, position *Position, color *Color) *Rook {
	rook := new(Rook)
	rook.board = board
	rook.positon = position
	rook.color = color
	return rook
}

func (rook *Rook) PossibleMoves() [8][8]bool {
	currentRow, currentColumn := rook.positon.row, rook.positon.column
	var possibleMoves [8][8]bool
	boardArray := rook.board.BoardArray

	// Up
	for i := currentRow - 1; i >= 0; i-- {
		positionContent := boardArray[i][currentColumn]
		if positionContent == nil {
			possibleMoves[i][currentColumn] = true
		} else if !positionContent.IsSameTeam(rook) {
			possibleMoves[i][currentColumn] = true
			break
		} else {
			possibleMoves[i][currentColumn] = false
			break
		}
	}

	// Down
	for i := currentRow + 1; i < 8; i++ {
		positionContent := boardArray[i][currentColumn]
		if positionContent == nil {
			possibleMoves[i][currentColumn] = true
		} else if !positionContent.IsSameTeam(rook) {
			possibleMoves[i][currentColumn] = true
			break
		} else {
			possibleMoves[i][currentColumn] = false
			break
		}
	}

	// Right
	for i := currentColumn + 1; i < 8; i++ {
		positionContent := boardArray[currentRow][i]
		if positionContent == nil {
			possibleMoves[currentRow][i] = true
		} else if !positionContent.IsSameTeam(rook) {
			possibleMoves[currentRow][i] = true
			break
		} else {
			possibleMoves[currentRow][i] = false
			break
		}
	}

	// Left
	for i := currentColumn - 1; i >= 0; i-- {
		positionContent := boardArray[currentRow][i]
		if positionContent == nil {
			possibleMoves[currentRow][i] = true
		} else if !positionContent.IsSameTeam(rook) {
			possibleMoves[currentRow][i] = true
			break
		} else {
			possibleMoves[currentRow][i] = false
			break
		}
	}

	return possibleMoves
}

func (rook *Rook) MakeMove(moveTo Position) bool {
	//currentRow, currentColumn, moveToRow, moveToColumn := rook.positon.row, rook.positon.column, moveTo.row, moveTo.column
	return true
}

func (rook *Rook) Position() Position {
	return *rook.positon
}

func (rook *Rook) IsSameTeam(piece Piece) bool {
	return piece.Color().isWhite() == rook.Color().isWhite()
}

func (rook *Rook) Color() Color {
	return *rook.color
}

func (rook *Rook) ToString() string {
	if rook.Color().isWhite() {
		return "R"
	} else {
		return "r"
	}
}

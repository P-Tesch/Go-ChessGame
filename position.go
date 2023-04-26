package main

type Position struct {
	// Board Positions
	row    int
	column int

	// Chess Positions
	chessRow    int
	chessColumn string
}

func NewPosition(row int, column int) *Position {
	position := new(Position)
	position.SetBoardPosition(row, column)
	return position
}

func (position *Position) SetBoardPosition(row int, column int) {
	position.row = row
	position.column = column
	position.chessRow = 8 - row
	position.chessColumn = convertBoardColumnToChessColumn(column)
}

func (position *Position) SetChessPosition(row int, column string) {
	position.row = 8 - row
	position.column = convertChessColumnToBoardColumn(column)
	position.chessRow = row
	position.chessColumn = column
}

func (position *Position) GetBoardPosition() (int, int) {
	return position.row, position.column
}

func (position *Position) GetChessPostion() (int, string) {
	return position.chessRow, position.chessColumn
}

func convertBoardColumnToChessColumn(column int) string {
	possibleColumns := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return possibleColumns[column]
}

func convertChessColumnToBoardColumn(column string) int {
	possibleColumns := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for index, element := range possibleColumns {
		if element == column {
			return index
		}
	}
	return -1
}

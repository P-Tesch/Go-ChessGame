package main

type Piece interface {
	PossibleMoves() [8][8]bool
	MakeMove(moveTo Position) bool
	Position() Position
	Color() Color
	IsSameTeam(piece Piece) bool
	ToString() string
}

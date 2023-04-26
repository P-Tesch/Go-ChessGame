package main

type Color struct {
	white bool
}

func NewColor(isWhite bool) *Color {
	color := new(Color)
	color.white = isWhite
	return color
}

func (color Color) isWhite() bool {
	return color.white
}

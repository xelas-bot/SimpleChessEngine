package chess_utils

import (
	"strconv"
)

type Piece struct {
	Name  string
	Color int8
	X_Pos int8
	Y_Pos int8
}

func NewPiece(name string, color int8, x_pos int8, y_pos int8) *Piece {
	p := Piece{Name: name, Color: color, X_Pos: x_pos, Y_Pos: y_pos}
	return &p

}

func ReturnChessNot(piece Piece) string {
	ToReturn := ""
	ToReturn = ToReturn + string(piece.X_Pos+64)
	ToReturn = ToReturn + strconv.Itoa(int(piece.X_Pos))
	return ToReturn
}

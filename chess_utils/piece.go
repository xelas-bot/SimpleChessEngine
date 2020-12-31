package chess_utils

type Piece struct {
	name  string
	color int8
}

func NewPiece(name string, color int8) *Piece {
	p := Piece{name: name, color: color}
	return &p

}

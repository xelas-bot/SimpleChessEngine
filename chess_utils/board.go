package chess_utils

type Board struct {
	Pawn   Piece
	King   Piece
	Pieces []Piece
}

func NewBoard() *Board {
	board := Board{}

	piece_one := Piece{"KING", 0, 0, 0}
	piece_two := Piece{"PAWN", 0, 0, 0}

	board.King = piece_one
	board.Pawn = piece_two

	for i := 1; i <= 8; i++ {
		p := Piece{"PAWN", 0, int8(i), 2}
		board.Pieces = append(board.Pieces, p)
	}

	return &board

}

func PrintBoard() string {

	return ""

}

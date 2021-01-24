package chess_utils

import (
	"strconv"
)

type Piece struct {
	Name     rune
	Player   bool
	X_Pos    int8
	Y_Pos    int8
	HasMoved bool
}

type Move struct {
	Piece    *Piece
	CapPiece *Piece
	X_Pos    int8
	Y_Pos    int8
}

func IsEmpty(piece *Piece) bool {
	return piece.Name == 'E'
}
func IsPiece(piece *Piece) bool {
	return piece.Name != 'E'
}

func IsWhite(piece *Piece) bool {
	return !piece.Player
}

func IsBlack(piece *Piece) bool {
	return piece.Player
}

func NewPiece(name rune, player bool, x_pos int8, y_pos int8) *Piece {
	p := Piece{Name: name, Player: player, X_Pos: x_pos, Y_Pos: y_pos, HasMoved: false}
	return &p
}

func ReturnChessNot(piece Piece) string {
	ToReturn := ""
	ToReturn = ToReturn + string(piece.X_Pos+64)
	ToReturn = ToReturn + strconv.Itoa(int(piece.X_Pos+1))
	return ToReturn
}

/*

func Moves(piece *Piece) []Point {
	initial_moves := []Point{}
	if piece.Color == 0 {
		initial_moves = append(initial_moves, Point{-1, -1})
		return initial_moves
	}

	if piece.Name == 'P' {
		if piece.Y_Pos != 1 && piece.Y_Pos != 6 && piece.Y_Pos != 0 {
			if piece.Color == 1 {
				initial_moves = append(initial_moves, Point{piece.X_Pos, piece.Y_Pos + 1})
				return initial_moves
			}
			if piece.Color == 2 {
				initial_moves = append(initial_moves, Point{piece.X_Pos, piece.Y_Pos - 1})
				return initial_moves
			}
		}
		if piece.Y_Pos == 1 && piece.Color == 1 {
			initial_moves = append(initial_moves, Point{piece.X_Pos, piece.Y_Pos + 1})
			initial_moves = append(initial_moves, Point{piece.X_Pos, piece.Y_Pos + 2})
			return initial_moves
		}
		if piece.Y_Pos == 6 && piece.Color == 2 {
			initial_moves = append(initial_moves, Point{piece.X_Pos, piece.Y_Pos - 1})
			initial_moves = append(initial_moves, Point{piece.X_Pos, piece.Y_Pos - 2})
			return initial_moves
		}
	}

	return initial_moves
}
*/

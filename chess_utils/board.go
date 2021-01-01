package chess_utils

import (
	"fmt"
	"strconv"
)

type Board struct {
	Pawn   Piece
	King   Piece
	Pieces []*Piece
}

func NewBoard() *Board {
	board := Board{}

	piece_one := Piece{"KING", 0, 0, 0}
	piece_two := Piece{"PAWN", 0, 0, 0}

	board.King = piece_one
	board.Pawn = piece_two
	// white pawns
	for i := 0; i < 8; i++ {
		p := NewPiece("PAWN", 1, int8(i), 1)
		board.Pieces = append(board.Pieces, p)
	}
	//black pawns
	for i := 0; i < 8; i++ {
		p := NewPiece("PAWN", 2, int8(i), 6)
		board.Pieces = append(board.Pieces, p)
	}

	//black rook
	board.Pieces = append(board.Pieces, NewPiece("ROOK", 2, 0, 7))
	board.Pieces = append(board.Pieces, NewPiece("ROOK", 2, 7, 7))
	//black knights
	board.Pieces = append(board.Pieces, NewPiece("KNIGHT", 2, 1, 7))
	board.Pieces = append(board.Pieces, NewPiece("KNIGHT", 2, 6, 7))
	//black bishops
	board.Pieces = append(board.Pieces, NewPiece("BISHOP", 2, 2, 7))
	board.Pieces = append(board.Pieces, NewPiece("BISHOP", 2, 5, 7))

	//White rook
	board.Pieces = append(board.Pieces, NewPiece("ROOK", 1, 0, 0))
	board.Pieces = append(board.Pieces, NewPiece("ROOK", 1, 7, 0))
	//White knights
	board.Pieces = append(board.Pieces, NewPiece("KNIGHT", 1, 1, 0))
	board.Pieces = append(board.Pieces, NewPiece("KNIGHT", 1, 6, 0))
	//White bishops
	board.Pieces = append(board.Pieces, NewPiece("BISHOP", 1, 2, 0))
	board.Pieces = append(board.Pieces, NewPiece("BISHOP", 1, 5, 0))

	return &board

}

func PrintBoard(Pieces []*Piece) string {

	var board [64]int

	for _, element := range Pieces {
		board[((7-element.Y_Pos)*8)+element.X_Pos] = int(element.Color)

	}

	for i := 0; i < len(board); i++ {
		if i%8 == 0 {
			fmt.Println()
		}
		fmt.Print(strconv.Itoa(board[i]) + " ")

	}

	return ""

}
func GetPieceAt(board *Board, x int8, y int8) *Piece {

	for i := 0; i < len(board.Pieces); i++ {
		if board.Pieces[i].X_Pos == x && board.Pieces[i].Y_Pos == y {
			return board.Pieces[i]
		}
	}

	return NewPiece("PIECE NOT FOUND", 0, 0, 0)
}

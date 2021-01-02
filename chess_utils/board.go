package chess_utils

import (
	"fmt"
)

var KnightMoves = [16]int8{1, 2, 2, 1, 2, -1, 1, -2, -1, -2, -2, -1, -2, 1, -1, 2}

type Board struct {
	WhitePieces []*Piece
	BlackPieces []*Piece
	Turn   bool
}

func GetBoardMoves(board *Board) []Move{} {
	BetterCaptureList = []Move
	CaptureList = []Move
	MoveList = []Move

	// black turn
	if board.Turn {
		
	}

	for _, piece := range board.Pieces {
		if piece.Name == 'P' {
			if IsWhite(piece) {
				if IsEmpty(GetPieceAt(piece.X_Pos, piece.Y_Pos + 1)) {
					MoveList = append(MoveList, Move{piece, piece,X_Pos, piece.Y_Pos + 1})
					if IsEmpty(GetPieceAt(piece.X_Pos, piece.Y_Pos + 2)) {
						MoveList = append(MoveList, Move{piece, piece,X_Pos, piece.Y_Pos + 2})
					}
				}
			} else {
				if IsEmpty(GetPieceAt(piece.X_Pos, piece.Y_Pos - 1)) {
					MoveList = append(MoveList, Move{piece, piece,X_Pos, piece.Y_Pos - 1})
					if IsEmpty(GetPieceAt(piece.X_Pos, piece.Y_Pos - 2)) {
						MoveList = append(MoveList, Move{piece, piece,X_Pos, piece.Y_Pos - 2})
					}
				}
			}
		} else if piece.Name == 'N' {
			for i := 0; i < 8; i++ {
				if OutOfBounds(piece.X_Pos + KnightMoves[i], piece.Y_Pos + KnightMoves[i + 1]) {
					continue
				}
				tile := GetPieceAt(piece.X_Pos + KnightMoves[i], piece.Y_Pos + KnightMoves[i + 1])
				if IsEmpty(tile) {
					MoveList = append(MoveList, Move{piece, piece.X_Pos + KnightMoves[i], piece.Y_Pos + KnightMoves[i + 1]})
				} else if else if GetValue(tile) > GetValue(piece) {
					BetterCaptureList = append(BetterCaptureList, Move{piece, piece.X_Pos + KnightMoves[i], piece.Y_Pos + KnightMoves[i + 1]})
				} else {
					CaptureList = append(CaptureList, Move{piece, piece.X_Pos + KnightMoves[i], piece.Y_Pos + KnightMoves[i + 1]})
				}
			}
		} else if piece.Name == 'R'{
			
			// Add something to stop counting tiles if rook file is blocked by a piece

			// Vertical File
			for i:=0; i<8;i++{

			}

			// Horizontal File
			for i:=0; i<8;i++{

			}

			

		}
	}
}

func GetValue(piece *Piece) int8 {
	switch piece.Name {
		case 'P':
			return 100
		case 'N':
			return 320
		case 'B':
			return 330
		case 'R':
			return 500
		case 'Q':
			return 900
		case 'K':
			return 696969
	}
	return 0
}

func MovePiece(piece *Piece, move Move){

}

func NewBoard() *Board {
	board := Board

	// white pawns
	for i := 0; i < 8; i++ {
		board.Pieces = append(board.WhitePieces, NewPiece('P', false, int8(i), 1))
	}
	// white king
	board.Pieces = append(board.WhitePieces, NewPiece('K', false, 4, 0))
	// white queen
	board.Pieces = append(board.WhitePieces, NewPiece('Q', false, 3, 0))
	// white rook
	board.Pieces = append(board.WhitePieces, NewPiece('R', false, 0, 0))
	board.Pieces = append(board.WhitePieces, NewPiece('R', false, 7, 0))
	// white knights
	board.Pieces = append(board.WhitePieces, NewPiece('N', false, 1, 0))
	board.Pieces = append(board.WhitePieces, NewPiece('N', false, 6, 0))
	// white bishops
	board.Pieces = append(board.WhitePieces, NewPiece('B', false, 2, 0))
	board.Pieces = append(board.WhitePieces, NewPiece('B', false, 5, 0))

	// black pawns
	for i := 0; i < 8; i++ {
		board.Pieces = append(board.BlackPieces, NewPiece('P', true, int8(i), 6))
	}
	// black king
	board.Pieces = append(board.BlackPieces, NewPiece('K', true, 4, 7))
	// black queen
	board.Pieces = append(board.BlackPieces, NewPiece('Q', true, 3, 7))
	// black rook
	board.Pieces = append(board.BlackPieces, NewPiece('R', true, 0, 7))
	board.Pieces = append(board.BlackPieces, NewPiece('R', true, 7, 7))
	// black knight
	board.Pieces = append(board.BlackPieces, NewPiece('N', true, 1, 7))
	board.Pieces = append(board.BlackPieces, NewPiece('N', true, 6, 7))
	// black bishop
	board.Pieces = append(board.BlackPieces, NewPiece('B', true, 2, 7))
	board.Pieces = append(board.BlackPieces, NewPiece('B', true, 5, 7))

	return &board
}

func PrintBoard(Pieces []*Piece) string {

	var board [64]rune

	for _, element := range Pieces {
		board[((7-element.Y_Pos)*8)+element.X_Pos] = element.Name

	}

	for i := 0; i < len(board); i++ {
		if i%8 == 0 {
			fmt.Println()
		}
		fmt.Print(string(board[i]) + " ")

	}

	return ""

}
func OutOfBounds(x int8, y int8) bool {
	return x < 0 || y < 0 || x > 7 || y > 7
}
func GetPieceAt(board *Board, x int8, y int8) *Piece {
	for i := 0; i < len(board.Pieces); i++ {
		if board.Pieces[i].X_Pos == x && board.Pieces[i].Y_Pos == y {
			return board.Pieces[i]
		}
	}
	return NewPiece('E', false, 0, 0)
}

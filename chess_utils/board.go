package chess_utils

import (
	"fmt"
)

var KnightMoves = [16]int8{1, 2, 2, 1, 2, -1, 1, -2, -1, -2, -2, -1, -2, 1, -1, 2}

type Board struct {
	WhitePieces []*Piece
	BlackPieces []*Piece
	Turn        bool
}

func GetBoardMoves(board *Board) []Move {
	BetterCaptureList := []Move{}
	CaptureList := []Move{}
	MoveList := []Move{}
	//black is true
	// black turn
	if board.Turn {

	}

	for _, piece := range board.WhitePieces {
		if piece.Name == 'P' {
			if IsWhite(piece) {
				if IsEmpty(GetPieceAt(board, piece.X_Pos, piece.Y_Pos+1)) {
					// might need to use a pointer to move rather than creating move as value
					MoveList = append(MoveList, Move{piece, piece.X_Pos, piece.Y_Pos + 1})
					if IsEmpty(GetPieceAt(board, piece.X_Pos, piece.Y_Pos+2)) {
						MoveList = append(MoveList, Move{piece, piece.X_Pos, piece.Y_Pos + 2})
					}
				}
			} else {
				if IsEmpty(GetPieceAt(board, piece.X_Pos, piece.Y_Pos-1)) {
					MoveList = append(MoveList, Move{piece, piece.X_Pos, piece.Y_Pos - 1})
					if IsEmpty(GetPieceAt(board, piece.X_Pos, piece.Y_Pos-2)) {
						MoveList = append(MoveList, Move{piece, piece.X_Pos, piece.Y_Pos - 2})
					}
				}
			}
		} else if piece.Name == 'N' {
			for i := 0; i < 8; i++ {
				if OutOfBounds(piece.X_Pos+KnightMoves[i], piece.Y_Pos+KnightMoves[i+1]) {
					continue
				}
				tile := GetPieceAt(board, piece.X_Pos+KnightMoves[i], piece.Y_Pos+KnightMoves[i+1])
				if IsEmpty(tile) {
					MoveList = append(MoveList, Move{piece, piece.X_Pos + KnightMoves[i], piece.Y_Pos + KnightMoves[i+1]})
				} else if GetValue(tile) > GetValue(piece) {
					BetterCaptureList = append(BetterCaptureList, Move{piece, piece.X_Pos + KnightMoves[i], piece.Y_Pos + KnightMoves[i+1]})
				} else {
					CaptureList = append(CaptureList, Move{piece, piece.X_Pos + KnightMoves[i], piece.Y_Pos + KnightMoves[i+1]})
				}
			}
		} else if piece.Name == 'R' {

			// Add something to stop counting tiles if rook file is blocked by a piece
			counter := piece.X_Pos

			// Horizontal File

			for i := counter + 1; i < 8; i++ {
				tile := GetPieceAt(board, counter, piece.Y_Pos)
				if IsPiece(GetPieceAt(board, counter, piece.Y_Pos)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, counter, piece.Y_Pos})
					} else {
						CaptureList = append(CaptureList, Move{piece, counter, piece.Y_Pos})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, counter, piece.Y_Pos})
				}
			}

			for i := counter - 1; i >= 0; i-- {
				tile := GetPieceAt(board, counter, piece.Y_Pos)
				if IsPiece(GetPieceAt(board, counter, piece.Y_Pos)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, counter, piece.Y_Pos})
					} else {
						CaptureList = append(CaptureList, Move{piece, counter, piece.Y_Pos})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, counter, piece.Y_Pos})
				}
			}

			yCounter := piece.Y_Pos
			// Vertical File
			for i := yCounter + 1; i < 8; i++ {
				tile := GetPieceAt(board, piece.X_Pos, yCounter)
				if IsPiece(GetPieceAt(board, piece.X_Pos, yCounter)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, piece.X_Pos, yCounter})
					} else {
						CaptureList = append(CaptureList, Move{piece, piece.X_Pos, yCounter})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, piece.X_Pos, yCounter})
				}
			}
			for i := yCounter - 1; i >= 0; i-- {
				tile := GetPieceAt(board, piece.X_Pos, yCounter)
				if IsPiece(GetPieceAt(board, piece.X_Pos, yCounter)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, piece.X_Pos, yCounter})
					} else {
						CaptureList = append(CaptureList, Move{piece, piece.X_Pos, yCounter})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, piece.X_Pos, yCounter})
				}
			}
		} else if piece.Name == 'Q' {
			// Add something to stop counting tiles if rook file is blocked by a piece
			counter := piece.X_Pos

			// Horizontal File

			for i := counter + 1; i < 8; i++ {
				tile := GetPieceAt(board, counter, piece.Y_Pos)
				if IsPiece(GetPieceAt(board, counter, piece.Y_Pos)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, counter, piece.Y_Pos})
					} else {
						CaptureList = append(CaptureList, Move{piece, counter, piece.Y_Pos})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, counter, piece.Y_Pos})
				}
			}

			for i := counter - 1; i >= 0; i-- {
				tile := GetPieceAt(board, counter, piece.Y_Pos)
				if IsPiece(GetPieceAt(board, counter, piece.Y_Pos)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, counter, piece.Y_Pos})
					} else {
						CaptureList = append(CaptureList, Move{piece, counter, piece.Y_Pos})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, counter, piece.Y_Pos})
				}
			}

			yCounter := piece.Y_Pos
			// Vertical File
			for i := yCounter + 1; i < 8; i++ {
				tile := GetPieceAt(board, piece.X_Pos, yCounter)
				if IsPiece(GetPieceAt(board, piece.X_Pos, yCounter)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, piece.X_Pos, yCounter})
					} else {
						CaptureList = append(CaptureList, Move{piece, piece.X_Pos, yCounter})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, piece.X_Pos, yCounter})
				}
			}
			for i := yCounter - 1; i >= 0; i-- {
				tile := GetPieceAt(board, piece.X_Pos, yCounter)
				if IsPiece(GetPieceAt(board, piece.X_Pos, yCounter)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, piece.X_Pos, yCounter})
					} else {
						CaptureList = append(CaptureList, Move{piece, piece.X_Pos, yCounter})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, piece.X_Pos, yCounter})
				}
			}

		}
	}

	return nil
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

func MovePiece(piece *Piece, move Move) {

}

func NewBoard() *Board {
	board := Board{Turn: false}

	// white pawns
	for i := 0; i < 8; i++ {
		board.WhitePieces = append(board.WhitePieces, NewPiece('P', false, int8(i), 1))
	}
	// white king
	board.WhitePieces = append(board.WhitePieces, NewPiece('K', false, 4, 0))
	// white queen
	board.WhitePieces = append(board.WhitePieces, NewPiece('Q', false, 3, 0))
	// white rook
	board.WhitePieces = append(board.WhitePieces, NewPiece('R', false, 0, 0))
	board.WhitePieces = append(board.WhitePieces, NewPiece('R', false, 7, 0))
	// white knights
	board.WhitePieces = append(board.WhitePieces, NewPiece('N', false, 1, 0))
	board.WhitePieces = append(board.WhitePieces, NewPiece('N', false, 6, 0))
	// white bishops
	board.WhitePieces = append(board.WhitePieces, NewPiece('B', false, 2, 0))
	board.WhitePieces = append(board.WhitePieces, NewPiece('B', false, 5, 0))

	// black pawns
	for i := 0; i < 8; i++ {
		board.BlackPieces = append(board.BlackPieces, NewPiece('P', true, int8(i), 6))
	}
	// black king
	board.BlackPieces = append(board.BlackPieces, NewPiece('K', true, 4, 7))
	// black queen
	board.BlackPieces = append(board.BlackPieces, NewPiece('Q', true, 3, 7))
	// black rook
	board.BlackPieces = append(board.BlackPieces, NewPiece('R', true, 0, 7))
	board.BlackPieces = append(board.BlackPieces, NewPiece('R', true, 7, 7))
	// black knight
	board.BlackPieces = append(board.BlackPieces, NewPiece('N', true, 1, 7))
	board.BlackPieces = append(board.BlackPieces, NewPiece('N', true, 6, 7))
	// black bishop
	board.BlackPieces = append(board.BlackPieces, NewPiece('B', true, 2, 7))
	board.BlackPieces = append(board.BlackPieces, NewPiece('B', true, 5, 7))

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

	for i := 0; i < len(board.WhitePieces); i++ {
		if board.WhitePieces[i].X_Pos == x && board.WhitePieces[i].Y_Pos == y {
			return board.WhitePieces[i]
		}
	}

	for i := 0; i < len(board.BlackPieces); i++ {
		if board.BlackPieces[i].X_Pos == x && board.BlackPieces[i].Y_Pos == y {
			return board.BlackPieces[i]
		}
	}

	return NewPiece('E', false, 0, 0)
}

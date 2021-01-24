package chess_utils

import (
	"fmt"
)

var KnightMoves = [16]int8{1, 2, 2, 1, 2, -1, 1, -2, -1, -2, -2, -1, -2, 1, -1, 2}

var PawnTable = [64]int{0, 0, 0, 0, 0, 0, 0, 0,
	50, 50, 50, 50, 50, 50, 50, 50,
	10, 10, 20, 30, 30, 20, 10, 10,
	5, 5, 10, 25, 25, 10, 5, 5,
	0, 0, 0, 20, 20, 0, 0, 0,
	5, -5, -10, 0, 0, -10, -5, 5,
	5, 10, 10, -20, -20, 10, 10, 5,
	0, 0, 0, 0, 0, 0, 0, 0}

var KnightTable = [64]int{-50, -40, -30, -30, -30, -30, -40, -50,
	-40, -20, 0, 0, 0, 0, -20, -40,
	-30, 0, 10, 15, 15, 10, 0, -30,
	-30, 5, 15, 20, 20, 15, 5, -30,
	-30, 0, 15, 20, 20, 15, 0, -30,
	-30, 5, 10, 15, 15, 10, 5, -30,
	-40, -20, 0, 5, 5, 0, -20, -40,
	-50, -40, -30, -30, -30, -30, -40, -50}

var BishopTable = [64]int{-20, -10, -10, -10, -10, -10, -10, -20,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-10, 0, 5, 10, 10, 5, 0, -10,
	-10, 5, 5, 10, 10, 5, 5, -10,
	-10, 0, 10, 10, 10, 10, 0, -10,
	-10, 10, 10, 10, 10, 10, 10, -10,
	-10, 5, 0, 0, 0, 0, 5, -10,
	-20, -10, -10, -10, -10, -10, -10, -20}

var RookTable = [64]int{0, 0, 0, 0, 0, 0, 0, 0,
	5, 10, 10, 10, 10, 10, 10, 5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	0, 0, 0, 5, 5, 0, 0, 0}

var QueenTable = [64]int{-20, -10, -10, -5, -5, -10, -10, -20,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-10, 0, 5, 5, 5, 5, 0, -10,
	-5, 0, 5, 5, 5, 5, 0, -5,
	0, 0, 5, 5, 5, 5, 0, -5,
	-10, 5, 5, 5, 5, 5, 0, -10,
	-10, 0, 5, 0, 0, 0, 0, -10,
	-20, -10, -10, -5, -5, -10, -10, -20}

var KingTable = [64]int{-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-20, -30, -30, -40, -40, -30, -30, -20,
	-10, -20, -20, -20, -20, -20, -20, -10,
	20, 20, 0, 0, 0, 0, 20, 20,
	20, 30, 10, 0, 0, 10, 30, 20}

var KingEndTable = [64]int{-50, -40, -30, -20, -20, -30, -40, -50,
	-30, -20, -10, 0, 0, -10, -20, -30,
	-30, -10, 20, 30, 30, 20, -10, -30,
	-30, -10, 30, 40, 40, 30, -10, -30,
	-30, -10, 30, 40, 40, 30, -10, -30,
	-30, -10, 20, 30, 30, 20, -10, -30,
	-30, -30, 0, 0, 0, 0, -30, -30,
	-50, -30, -30, -30, -30, -30, -30, -50}

type Board struct {
	WhitePieces []*Piece
	BlackPieces []*Piece
	Turn        bool
}

func Evaluate(board *Board) int {
	Score := 0

	EndGame := IsEndgame(board)
	for _, piece := range board.WhitePieces {
		Score += GetValue(piece) + GetPieceTableValue(piece, EndGame)
	}
	for _, piece := range board.BlackPieces {
		Score -= GetValue(piece) + GetPieceTableValue(piece, EndGame)
	}
	return Score
}

func ExecuteMove(move Move, board *Board) {
	tempPiece := GetPieceAt(board, move.X_Pos, move.Y_Pos)
	index_ := 0

	pieceSet := &board.BlackPieces
	if board.Turn {
		pieceSet = &board.WhitePieces
	}

	if IsPiece(tempPiece) {
		for index, piece := range *pieceSet {
			if piece.X_Pos == move.CapPiece.X_Pos && piece.Y_Pos == move.CapPiece.Y_Pos {
				index_ = index
			}
		}

		move.Piece.X_Pos = move.X_Pos
		move.Piece.Y_Pos = move.Y_Pos
		//fmt.Print(len(pieceSet))
		(*pieceSet)[index_] = (*pieceSet)[len(*pieceSet)-1] // Copy last element to index i.
		*pieceSet = (*pieceSet)[:len(*pieceSet)-1]          // Truncate slice.

	} else {
		move.Piece.X_Pos = move.X_Pos
		move.Piece.Y_Pos = move.Y_Pos
	}

	//fmt.Print(len(pieceSet))
	board.Turn = !board.Turn

}

func UndoMove(move Move, board *Board) {
	tempPiece := GetPieceAt(board, move.X_Pos, move.Y_Pos)
	index_ := 0

	pieceSet := board.BlackPieces
	if board.Turn {
		pieceSet = board.WhitePieces
	}

	if IsPiece(tempPiece) {
		for index, piece := range pieceSet {
			if piece.X_Pos == move.CapPiece.X_Pos && piece.Y_Pos == move.CapPiece.Y_Pos {
				index_ = index
			}
		}

		move.Piece.X_Pos = move.X_Pos
		move.Piece.Y_Pos = move.Y_Pos

		pieceSet[index_] = pieceSet[len(pieceSet)-1] // Copy last element to index i.
		pieceSet = pieceSet[:len(pieceSet)-1]        // Truncate slice.
	} else {
		move.Piece.X_Pos = move.X_Pos
		move.Piece.Y_Pos = move.Y_Pos
	}

	board.Turn = !board.Turn
}

func GetBoardMoves(board *Board) []Move {
	BetterCaptureList := []Move{}
	CaptureList := []Move{}
	MoveList := []Move{}
	//black is true
	// black turn

	var boardToUse = board.WhitePieces
	if board.Turn {
		boardToUse = board.BlackPieces
	}

	for _, piece := range boardToUse {
		if piece.Name == 'P' {
			pieceToCap := NewPiece('E', false, -1, -1)

			if IsEmpty(GetPieceAt(board, piece.X_Pos, piece.Y_Pos+1)) {

				MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, piece.Y_Pos + 1})
				if IsEmpty(GetPieceAt(board, piece.X_Pos, piece.Y_Pos+2)) && piece.Y_Pos == 1 {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, piece.Y_Pos + 2})
				}
			}

		} else if piece.Name == 'N' {
			for i := 0; i < 8; i++ {
				pieceToCap := GetPieceAt(board, int8(i), piece.Y_Pos)
				if OutOfBounds(piece.X_Pos+KnightMoves[i], piece.Y_Pos+KnightMoves[i+1]) {
					continue
				}
				tile := GetPieceAt(board, piece.X_Pos+KnightMoves[i], piece.Y_Pos+KnightMoves[i+1])
				if IsEmpty(tile) {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos + KnightMoves[i], piece.Y_Pos + KnightMoves[i+1]})
				} else if GetValue(tile) > GetValue(piece) {
					BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos + KnightMoves[i], piece.Y_Pos + KnightMoves[i+1]})
				} else {
					CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos + KnightMoves[i], piece.Y_Pos + KnightMoves[i+1]})
				}
			}
		} else if piece.Name == 'R' {

			// Add something to stop counting tiles if rook file is blocked by a piece
			counter := piece.X_Pos

			// Horizontal File

			for i := counter + 1; i < 8; i++ {
				tile := GetPieceAt(board, i, piece.Y_Pos)
				pieceToCap := GetPieceAt(board, i, piece.Y_Pos)
				if IsPiece(GetPieceAt(board, i, piece.Y_Pos)) {

					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, i, piece.Y_Pos})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, i, piece.Y_Pos})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, i, piece.Y_Pos})
				}
			}

			for i := counter - 1; i >= 0; i-- {
				pieceToCap := GetPieceAt(board, i, piece.Y_Pos)
				tile := GetPieceAt(board, i, piece.Y_Pos)
				if IsPiece(GetPieceAt(board, i, piece.Y_Pos)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, i, piece.Y_Pos})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, i, piece.Y_Pos})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, i, piece.Y_Pos})
				}
			}

			yCounter := piece.Y_Pos
			// Vertical File
			for i := yCounter + 1; i < 8; i++ {
				pieceToCap := GetPieceAt(board, piece.X_Pos, i)
				tile := GetPieceAt(board, piece.X_Pos, i)
				if IsPiece(GetPieceAt(board, piece.X_Pos, i)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, i})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, i})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, i})
				}
			}
			for i := yCounter - 1; i >= 0; i-- {
				pieceToCap := GetPieceAt(board, piece.X_Pos, i)
				tile := GetPieceAt(board, piece.X_Pos, i)
				if IsPiece(GetPieceAt(board, piece.X_Pos, i)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, i})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, i})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, i})
				}
			}
		} else if piece.Name == 'Q' {
			// Add something to stop counting tiles if rook file is blocked by a piece
			counter := piece.X_Pos

			// Horizontal File

			for i := counter + 1; i < 8; i++ {
				pieceToCap := GetPieceAt(board, i, piece.Y_Pos)
				tile := GetPieceAt(board, i, piece.Y_Pos)
				if IsPiece(GetPieceAt(board, i, piece.Y_Pos)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, i, piece.Y_Pos})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, i, piece.Y_Pos})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, i, piece.Y_Pos})
				}
			}

			for i := counter - 1; i >= 0; i-- {
				pieceToCap := GetPieceAt(board, i, piece.Y_Pos)
				tile := GetPieceAt(board, i, piece.Y_Pos)
				if IsPiece(GetPieceAt(board, i, piece.Y_Pos)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, i, piece.Y_Pos})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, i, piece.Y_Pos})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, i, piece.Y_Pos})
				}
			}

			yCounter := piece.Y_Pos
			// Vertical File
			for i := yCounter + 1; i < 8; i++ {
				pieceToCap := GetPieceAt(board, piece.X_Pos, i)
				tile := GetPieceAt(board, piece.X_Pos, i)

				if IsPiece(GetPieceAt(board, piece.X_Pos, i)) {

					if tile.Name != 'E' && tile.Player == piece.Player {

						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, i})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, i})
					}
					break
				} else {

					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, i})
					//fmt.Println(MoveList)
				}
			}
			for i := yCounter - 1; i >= 0; i-- {
				pieceToCap := GetPieceAt(board, piece.X_Pos, i)
				tile := GetPieceAt(board, piece.X_Pos, i)
				if IsPiece(GetPieceAt(board, piece.X_Pos, i)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, i})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, i})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos, i})
				}
			}

			//Diagonals

			xCounter := piece.X_Pos + 1
			yCounter = piece.Y_Pos + 1

			for x, y := xCounter, yCounter; x < 8 && y < 8; x, y = x+1, y+1 {
				pieceToCap := GetPieceAt(board, x, y)
				tile := GetPieceAt(board, x, y)
				if IsPiece(GetPieceAt(board, x, y)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
				}

			}

			xCounter = piece.X_Pos - 1
			yCounter = piece.Y_Pos + 1

			for x, y := xCounter, yCounter; x >= 0 && y < 8; x, y = x-1, y+1 {
				pieceToCap := GetPieceAt(board, x, y)
				tile := GetPieceAt(board, x, y)
				if IsPiece(GetPieceAt(board, x, y)) {
					//fmt.Print("HERE BRUH WHY TF THIS NO WORK")
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					}
					break
				} else {

					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
				}

			}

			xCounter = piece.X_Pos - 1
			yCounter = piece.Y_Pos - 1

			for x, y := xCounter, yCounter; x >= 0 && y >= 0; x, y = x-1, y-1 {
				pieceToCap := GetPieceAt(board, x, y)
				tile := GetPieceAt(board, x, y)
				if IsPiece(GetPieceAt(board, x, y)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
				}

			}

			xCounter = piece.X_Pos + 1
			yCounter = piece.Y_Pos - 1

			for x, y := xCounter, yCounter; x < 8 && y >= 0; x, y = x+1, y-1 {
				pieceToCap := GetPieceAt(board, x, y)
				tile := GetPieceAt(board, x, y)
				if IsPiece(GetPieceAt(board, x, y)) {

					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					}
					break
				} else {

					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
				}

			}

		} else if piece.Name == 'B' {
			//Diagonals

			xCounter := piece.X_Pos + 1
			yCounter := piece.Y_Pos + 1

			for x, y := xCounter, yCounter; x < 8 && y < 8; x, y = x+1, y+1 {
				pieceToCap := GetPieceAt(board, x, y)
				tile := GetPieceAt(board, x, y)
				if IsPiece(GetPieceAt(board, x, y)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
				}

			}

			xCounter = piece.X_Pos - 1
			yCounter = piece.Y_Pos + 1

			for x, y := xCounter, yCounter; x >= 0 && y < 8; x, y = x-1, y+1 {
				pieceToCap := GetPieceAt(board, x, y)
				tile := GetPieceAt(board, x, y)
				if IsPiece(GetPieceAt(board, x, y)) {
					//fmt.Print("HERE BRUH WHY TF THIS NO WORK")
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					}
					break
				} else {

					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
				}

			}

			xCounter = piece.X_Pos - 1
			yCounter = piece.Y_Pos - 1

			for x, y := xCounter, yCounter; x >= 0 && y >= 0; x, y = x-1, y-1 {
				pieceToCap := GetPieceAt(board, x, y)
				tile := GetPieceAt(board, x, y)
				if IsPiece(GetPieceAt(board, x, y)) {
					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					}
					break
				} else {
					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
				}

			}

			xCounter = piece.X_Pos + 1
			yCounter = piece.Y_Pos - 1

			for x, y := xCounter, yCounter; x < 8 && y >= 0; x, y = x+1, y-1 {
				pieceToCap := GetPieceAt(board, x, y)
				tile := GetPieceAt(board, x, y)
				if IsPiece(GetPieceAt(board, x, y)) {

					if tile.Name != 'E' && tile.Player == piece.Player {
						break
					} else if GetValue(tile) > GetValue(piece) {
						BetterCaptureList = append(BetterCaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					} else {
						CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
					}
					break
				} else {

					MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, x, y})
				}

			}
		} else if piece.Name == 'K' {
			pieceToCap := GetPieceAt(board, piece.X_Pos+1, piece.Y_Pos)

			tile := GetPieceAt(board, piece.X_Pos+1, piece.Y_Pos)
			if IsPiece(tile) && tile.Player != piece.Player {
				CaptureList = append(CaptureList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos + 1, piece.Y_Pos})
			} else if !IsPiece(tile) {
				MoveList = append(MoveList, Move{piece, pieceToCap, piece.X_Pos, piece.Y_Pos, piece.X_Pos + 1, piece.Y_Pos})
			}

		}
	}

	BetterCaptureList = append(BetterCaptureList, CaptureList...)
	BetterCaptureList = append(BetterCaptureList, MoveList...)
	return BetterCaptureList
}

func IsEndgame(board *Board) bool {
	WhiteQueen := false
	WhiteMinor := false
	for _, piece := range board.WhitePieces {
		if piece.Name == 'Q' {
			WhiteQueen = true
		} else if piece.Name != 'P' || piece.Name != 'K' {
			WhiteMinor = true
		}
	}
	BlackQueen := false
	BlackMinor := false
	for _, piece := range board.BlackPieces {
		if piece.Name == 'Q' {
			BlackQueen = true
		} else if piece.Name != 'P' || piece.Name != 'K' {
			BlackMinor = true
		}
	}
	return !WhiteQueen && !BlackQueen || !WhiteMinor && !BlackMinor
}

func GetPieceTableValue(piece *Piece, EndGame bool) int {
	var pos int8
	if IsWhite(piece) {
		pos = 64 - 8*piece.Y_Pos + piece.X_Pos
	} else {
		pos = 8*piece.Y_Pos + piece.X_Pos
	}
	switch piece.Name {
	case 'P':
		return PawnTable[pos]
	case 'N':
		return KnightTable[pos]
	case 'B':
		return BishopTable[pos]
	case 'R':
		return RookTable[pos]
	case 'Q':
		return QueenTable[pos]
	case 'K':
		if EndGame {
			return KingEndTable[pos]
		} else {
			return KingTable[pos]
		}
	}
	return 0
}

func GetValue(piece *Piece) int {
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
	board.WhitePieces = append(board.WhitePieces, NewPiece('B', false, 3, 5))
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

func PrintBoard(board_ *Board) string {
	Pieces := board_.WhitePieces

	var board [64]rune

	for index, _ := range board {
		board[index] = '■'

	}

	for _, element := range Pieces {
		board[((7-element.Y_Pos)*8)+element.X_Pos] = element.Name

	}

	Pieces = board_.BlackPieces
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

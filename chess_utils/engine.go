package chess_utils

import (
	"math"
)

var Moves = 0

var pruning = true

func Engine(board *Board, depth int) (int, Move) {
	MoveList := GetBoardMoves(board)

	bestValue := int(math.Inf(-1))
	bestMove := 0
	for i, move := range MoveList {
		ExecuteMove(board, move)
		value := 0
		if pruning {
			value = negamax_pruning(board, depth, int(math.Inf(-1)), int(math.Inf(1)), board.Turn)
		} else {
			value = negamax(board, depth, board.Turn)
		}
		UndoMove(board, move)
		if value > bestValue {
			bestValue = value
			bestMove = i
		}
	}
	return bestValue, MoveList[bestMove]
}

func negamax(board *Board, depth int, player bool) int {
	Moves++
	if depth == 0 {
		if player {
			return -Evaluate(board)
		} else {
			return Evaluate(board)
		}
	}

	MoveList := GetBoardMoves(board)
	value := int(math.Inf(-1))
	for _, move := range MoveList {
		ExecuteMove(board, move)
		eval := -negamax(board, depth-1, !player)
		UndoMove(board, move)
		if eval > value {
			value = eval
		}
	}
	return value
}

func negamax_pruning(board *Board, depth int, a int, b int, player bool) int {
	Moves++
	if depth == 0 {
		if player {
			return -Evaluate(board)
		} else {
			return Evaluate(board)
		}
	}

	MoveList := GetBoardMoves(board)
	value := int(math.Inf(-1))
	for _, move := range MoveList {
		ExecuteMove(board, move)
		eval := -negamax_pruning(board, depth-1, -b, -a, !player)
		UndoMove(board, move)
		if eval > value {
			value = eval
		}
		if value > a {
			a = value
		}
		if a >= b {
			break
		}
	}
	return value
}

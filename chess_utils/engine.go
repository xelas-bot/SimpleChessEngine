package chess_utils

import "math"

func Engine(board *Board, depth int) Move {
	MoveList := GetBoardMoves(board)

	bestValue := int(math.Inf(-1))
	bestMove := 0
	for i, move := range MoveList {
		value := negamax(board, depth, int(math.Inf(-1)), int(math.Inf(1)), board.Turn)
		if value > bestValue {
			bestValue = value
			bestMove = i
		}
	}
	return MoveList[bestMove]
}

func negamax(board *Board, depth int, a int, b int, player bool) int {
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

		eval := negamax(board, depth-1, -b, -a, !player)
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

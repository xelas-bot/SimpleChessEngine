package main

import (
	"ChessEngine/chess_utils"
	"fmt"
)

func main() {
	board := chess_utils.NewBoard()

	moveList := chess_utils.GetBoardMoves(board)

	fmt.Println(len(moveList))

	for index := range moveList {
		fmt.Print(string(moveList[index].Piece.Name))
		fmt.Print(moveList[index].X_Pos)
		fmt.Print(", ")
		fmt.Print(moveList[index].Y_Pos)
		fmt.Println()
	}
	chess_utils.PrintBoard(board.WhitePieces)

	/*
		listeners := [5]int{11, 18, 20, 21, 23}
		towers := [4]int{1, 2, 13, 22}
		distances := []int{}
		boundary_tower := 0

		//distances = append(distances, listeners[0] - towers[0])
		//distances = append(distances, towers[len(towers)-1] - listeners[len(listeners)-1])

		//int(math.Min(float64(towers[boundary_tower]-listeners[i]), float64(listeners[i+1]-towers[boundary_tower])))

		for i := 0; i < len(listeners); i++ {
			dist_1 := 0
			dist_2 := 0

			for boundary_tower < len(towers) {

				if towers[boundary_tower] > listeners[i] {
					break
				}

				boundary_tower++

				if boundary_tower == len(towers) {
					boundary_tower = boundary_tower - 1
					break
				}

			}

			if boundary_tower-1 >= 0 {
				if i == len(listeners)-1 && towers[boundary_tower] < listeners[i] {
					dist_1 = listeners[i] - towers[boundary_tower]
				} else {
					dist_1 = listeners[i] - towers[boundary_tower-1]
				}

			}

			if towers[boundary_tower] > listeners[i] {
				dist_2 = towers[boundary_tower] - listeners[i]
			} else {
				dist_2 = listeners[i] - towers[boundary_tower]
			}

			distances = append(distances, int(math.Min(float64(dist_1), float64(dist_2))))

		}*/

}

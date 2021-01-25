package main

import (
	"ChessEngine/chess_utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	board := chess_utils.NewBoard()
	board.Turn = false
	fmt.Println("ORIG")
	chess_utils.PrintBoard(board)
	fmt.Println()

	//moveList := chess_utils.GetBoardMoves(board)

	/*
		for index := range moveList {
			if moveList[index].Piece.Name == 'P' {
				fmt.Print(moveList[index].Piece.X_Pos)
				fmt.Print(", ")
				fmt.Print(moveList[index].Piece.Y_Pos)
				fmt.Print(string(moveList[index].Piece.Name))
				fmt.Print(moveList[index].X_Pos)
				fmt.Print(", ")
				fmt.Print(moveList[index].Y_Pos)
				fmt.Println()
			}

		}*/

	reader := bufio.NewReader(os.Stdin)
	for true {

		fmt.Print("x1: ")
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			value, move := chess_utils.Engine(board, 5)
			chess_utils.ExecuteMove(board, move)
			fmt.Println("COMP: " + strconv.Itoa(value))
			fmt.Println(chess_utils.Moves)
		} else {
			x1, _ := strconv.Atoi(string(text[0]))

			fmt.Print("y1: ")
			text, _ = reader.ReadString('\n')
			y1, _ := strconv.Atoi(string(text[0]))

			fmt.Print("x2: ")
			text, _ = reader.ReadString('\n')
			x2, _ := strconv.Atoi(string(text[0]))

			fmt.Print("y2: ")
			text, _ = reader.ReadString('\n')
			y2, _ := strconv.Atoi(string(text[0]))
			move := chess_utils.TextToMove(board, int8(x1), int8(y1), int8(x2), int8(y2))
			chess_utils.ExecuteMove(board, move)
		}

		fmt.Println()
		chess_utils.PrintBoard(board)
		fmt.Println()

	}

	chess_utils.PrintPieces(board)

	/*


		move := chess_utils.Move{
			Piece:    chess_utils.GetPieceAt(board, 3, 5),
			CapPiece: chess_utils.GetPieceAt(board, 4, 6),
			X_Pos:    4,
			Y_Pos:    6,
		}
		chess_utils.ExecuteMove(move, board)

		moveList := chess_utils.GetBoardMoves(board)

		for index := range moveList {
			if moveList[index].Piece.Name == 'P' {
				fmt.Print(moveList[index].Piece.X_Pos)
				fmt.Print(", ")
				fmt.Print(moveList[index].Piece.Y_Pos)
				fmt.Print(string(moveList[index].Piece.Name))
				fmt.Print(moveList[index].X_Pos)
				fmt.Print(", ")
				fmt.Print(moveList[index].Y_Pos)
				fmt.Println()
			}

		}

		chess_utils.PrintBoard(board)
		fmt.Println(len(board.BlackPieces))
	*/
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

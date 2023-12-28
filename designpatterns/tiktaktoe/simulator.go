package main

import "fmt"

func main() {
	pieceX := &PieceX{}
	pieceO := &PieceO{}
	player1 := NewPlayer("ameya", pieceX)
	player2 := NewPlayer("soumya", pieceO)

	board := NewBoard()
	game := NewGame(board, player1, player2)

	for game.finished == false {
		turn := game.getTurn()
		if turn == player1 {
			fmt.Println("Player 1 tell which square you want to play at")
			input := make([]int, 2)
			fmt.Scan(&input[0])
			fmt.Scan(&input[1])
			result := game.makeMove(player1, input[0], input[1])
			fmt.Println("board is ", game.board)
			if result == false {
				fmt.Println("Wrong move made, try again")
			}
		} else {
			fmt.Println("Player 2 tell which square you want to play at")
			input := make([]int, 2)
			fmt.Scan(&input[0])
			fmt.Scan(&input[1])
			result := game.makeMove(player2, input[0], input[1])
			fmt.Println("board is ", game.board)
			if result == false {
				fmt.Println("Wrong move made, try again")
			}
		}
	}
	fmt.Println("game is finished, winner is ", game.winner.name)
}

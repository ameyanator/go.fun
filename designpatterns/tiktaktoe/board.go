package main

type Board struct {
	Squares [][]string
}

func NewBoard() *Board {
	squares := make([][]string, 3)
	for i, _ := range squares {
		squares[i] = make([]string, 3)
	}
	return &Board{
		Squares: squares,
	}
}

func (b *Board) validMove(x, y int) bool {
	n := len(b.Squares)
	if x < 0 || y < 0 || x >= n || y >= n || b.Squares[x][y] != "" {
		return false
	}
	return true
}

func (b *Board) makeMove(x, y int, piece Piece) {
	if !b.validMove(x, y) {
		return
	}
	b.Squares[x][y] = piece.GetSymbol()
}

func (b *Board) isGameFinished() bool {
	// fmt.Println("Checking for game finished with board ", b)
	// fmt.Println(b.Squares[0][0], b.Squares[0][1], b.Squares[0][2], (b.Squares[0][0] == b.Squares[0][1]), (b.Squares[0][0] == b.Squares[0][2]))
	if b.Squares[0][0] != "" && b.Squares[0][0] == b.Squares[0][1] && b.Squares[0][0] == b.Squares[0][2] {
		return true
	}
	if b.Squares[0][0] != "" && b.Squares[0][0] == b.Squares[1][0] && b.Squares[0][0] == b.Squares[2][0] {
		return true
	}
	if b.Squares[0][0] != "" && b.Squares[0][0] == b.Squares[1][1] && b.Squares[0][0] == b.Squares[2][2] {
		return true
	}
	if b.Squares[0][1] != "" && b.Squares[0][1] == b.Squares[1][1] && b.Squares[0][1] == b.Squares[2][1] {
		return true
	}
	if b.Squares[0][2] != "" && b.Squares[0][2] == b.Squares[1][2] && b.Squares[0][2] == b.Squares[2][2] {
		return true
	}
	if b.Squares[0][2] != "" && b.Squares[0][2] == b.Squares[1][1] && b.Squares[0][0] == b.Squares[2][0] {
		return true
	}
	if b.Squares[1][0] != "" && b.Squares[1][0] == b.Squares[1][1] && b.Squares[1][0] == b.Squares[1][2] {
		return true
	}
	if b.Squares[2][0] != "" && b.Squares[2][0] == b.Squares[2][1] && b.Squares[2][0] == b.Squares[2][2] {
		return true
	}
	return false
}

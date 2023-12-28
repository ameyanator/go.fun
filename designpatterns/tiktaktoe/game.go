package main

type Game struct {
	board    *Board
	player1  *Player
	player2  *Player
	turn     int
	finished bool
	winner   *Player
}

func NewGame(board *Board, player1, player2 *Player) *Game {
	return &Game{
		board:    board,
		player1:  player1,
		player2:  player2,
		turn:     0,
		finished: false,
		winner:   nil,
	}
}

func (g *Game) getTurn() *Player {
	if g.turn%2 == 0 {
		return g.player1
	} else {
		return g.player2
	}
}

func (g *Game) makeMove(player *Player, x, y int) bool {
	if g.board.validMove(x, y) == false {
		return false
	}
	g.board.makeMove(x, y, player.piece)
	g.incrementTurn()
	if g.gameHasWinner() {
		g.winner = player
		g.finished = true
	}
	return true
}

func (g *Game) incrementTurn() {
	g.turn++
}

func (g *Game) gameHasWinner() bool {
	return g.board.isGameFinished()
}

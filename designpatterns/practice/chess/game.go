package main

import "errors"

type Game struct {
	board   *Board
	player1 *Player
	player2 *Player

	turn         int
	gameFinished bool
	winner       []*Player
}

func NewGame(player1, player2 *Player) *Game {
	return &Game{
		board:        NewBoard(),
		player1:      player1,
		player2:      player2,
		turn:         0,
		gameFinished: false,
		winner:       nil,
	}
}

func (g *Game) makeMove(move *Move) error {
	g.incrementTurn()
	if g.board.isValidMove(move) == false {
		return errors.New("move is invalid")
	}
	g.board.makeMove(move)
	finished, winner := g.isGameFinished()
	if finished {
		return nil
	}
	g.winner = winner
	return nil
}

func (g *Game) incrementTurn() {
	g.turn++
}

func (g *Game) getCurrentPlayer() *Player {
	if g.turn%2 == 0 {
		return g.player1
	} else {
		return g.player2
	}
}

func (g *Game) isGameFinished() (bool, []*Player) {
	finished, result := g.board.gameOver()
	if finished && result == "tie" {
		return true, []*Player{g.player1, g.player2}
	}
	if finished {
		return true, []*Player{g.getCurrentPlayer()}
	}
	return false, nil
}

package main

import "errors"

type Board struct {
	grid [8][8]Piece
}

func NewBoard() *Board {
	grid := [8][8]Piece{}
	grid[0][1] = &Pawn{color: "white"}
	return &Board{grid: grid}
}

type Location struct {
	x, y int
}

type Move struct {
	p        Piece
	from, to *Location
}

func (b *Board) isValidMove(move *Move) bool {
	if move.p.getPieceName() == "pawn" {
		return true
	}
	return false
}

func (b *Board) makeMove(move *Move) error {
	if !b.isValidMove(move) {
		return errors.New("this is an illegal move, please try another move")
	}
	b.grid[move.to.x][move.to.y] = move.p
	return nil
}

func (b *Board) gameOver() (bool, string) {
	return true, "tie"
}

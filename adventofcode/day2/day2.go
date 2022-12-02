package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func getScore(player1, player2 uint8) int {
	score := 0
	if player2 == uint8('X') {
		score += 1
	} else if player2 == uint8('Y') {
		score += 2
	} else {
		score += 3
	}
	if (player1 == uint8('A') && player2 == uint8('X')) || (player1 == uint8('B') && player2 == uint8('Y')) || (player1 == uint8('C') && player2 == uint8('Z')) {
		score += 3
	} else if (player1 == uint8('A') && player2 == uint8('Y')) || (player1 == uint8('B') && player2 == uint8('Z')) || (player1 == uint8('C') && player2 == uint8('X')) {
		score += 6
	}
	return score
}

func getTotalScore() (int, error) {
	f, err := os.OpenFile("input.txt", os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return -1, errors.New("Error Opening File: " + err.Error())
	}
	defer f.Close()

	// Add new line to end of file
	_, err = f.WriteString("\n")
	if err != nil {
		return -1, errors.New("Error Writing Adding newline to end of file: " + err.Error())
	}

	f.Seek(0, io.SeekStart)

	totalScore := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		game := scanner.Text()
		player1, player2 := game[0], game[2]
		totalScore += getScore(player1, player2)
	}
	return totalScore, nil
}

// Rock 1 A X
// Paper 2 B Y
// Scissors 3 C Z

func getScore2(player1, outcome uint8) int {
	score := 0
	if outcome == uint8('Y') { //draw
		score += 3
		if player1 == uint8('A') {
			score += 1
		} else if player1 == uint8('B') {
			score += 2
		} else if player1 == uint8('C') {
			score += 3
		}
	} else if outcome == uint8('Z') { //win
		score += 6
		if player1 == uint8('A') {
			score += 2
		} else if player1 == uint8('B') {
			score += 3
		} else if player1 == uint8('C') {
			score += 1
		}
	} else if outcome == uint8('X') { //lose
		if player1 == uint8('A') {
			score += 3
		} else if player1 == uint8('B') {
			score += 1
		} else if player1 == uint8('C') {
			score += 2
		}
	}
	return score
}

func getTotalScore2() (int, error) {
	f, err := os.OpenFile("input.txt", os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return -1, errors.New("Error Opening File: " + err.Error())
	}
	defer f.Close()

	// Add new line to end of file
	_, err = f.WriteString("\n")
	if err != nil {
		return -1, errors.New("Error Writing Adding newline to end of file: " + err.Error())
	}

	f.Seek(0, io.SeekStart)

	totalScore := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		game := scanner.Text()
		player1, outcome := game[0], game[2]
		totalScore += getScore2(player1, outcome)
	}
	return totalScore, nil
}

func main() {
	if score, err := getTotalScore2(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(score)
	}
}

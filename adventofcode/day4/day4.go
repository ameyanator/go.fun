package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	x, y int
}

func getCompletelyOverlappedAssignments() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return -1, errors.New("Couldn't open file with error: " + err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	overlappedIntervals := 0

	for scanner.Scan() {
		assignment := scanner.Text()
		pairs := strings.Split(assignment, ",")
		pairvalue1, pairvalue2 := strings.Split(pairs[0], "-"), strings.Split(pairs[1], "-")
		var p1, p2 pair
		p1.x, err = strconv.Atoi(pairvalue1[0])
		if err != nil {
			return -1, errors.New("Couldn't parse assignment number got error: " + err.Error())
		}
		p1.y, err = strconv.Atoi(pairvalue1[1])
		if err != nil {
			return -1, errors.New("Couldn't parse assignment number got error: " + err.Error())
		}
		p2.x, err = strconv.Atoi(pairvalue2[0])
		if err != nil {
			return -1, errors.New("Couldn't parse assignment number got error: " + err.Error())
		}
		p2.y, err = strconv.Atoi(pairvalue2[1])
		if err != nil {
			return -1, errors.New("Couldn't parse assignment number got error: " + err.Error())
		}
		if (p1.x <= p2.x && p2.y <= p1.y) || (p2.x <= p1.x && p1.y <= p2.y) {
			overlappedIntervals++
		}
	}
	return overlappedIntervals, nil
}

func getOverlappedAssignments() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return -1, errors.New("Couldn't open file with error: " + err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	overlappedIntervals := 0

	for scanner.Scan() {
		assignment := scanner.Text()
		pairs := strings.Split(assignment, ",")
		pairvalue1, pairvalue2 := strings.Split(pairs[0], "-"), strings.Split(pairs[1], "-")
		var p1, p2 pair
		p1.x, err = strconv.Atoi(pairvalue1[0])
		if err != nil {
			return -1, errors.New("Couldn't parse assignment number got error: " + err.Error())
		}
		p1.y, err = strconv.Atoi(pairvalue1[1])
		if err != nil {
			return -1, errors.New("Couldn't parse assignment number got error: " + err.Error())
		}
		p2.x, err = strconv.Atoi(pairvalue2[0])
		if err != nil {
			return -1, errors.New("Couldn't parse assignment number got error: " + err.Error())
		}
		p2.y, err = strconv.Atoi(pairvalue2[1])
		if err != nil {
			return -1, errors.New("Couldn't parse assignment number got error: " + err.Error())
		}
		if (p1.x <= p2.x && p2.x <= p1.y) || (p1.x <= p2.y && p2.y <= p1.y) || (p2.x <= p1.x && p1.x <= p2.y) || (p2.x <= p1.y && p1.y <= p2.y) {
			overlappedIntervals++
		}
	}
	return overlappedIntervals, nil
}

func main() {
	if total, err := getOverlappedAssignments(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(total)
	}
}

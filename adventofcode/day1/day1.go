package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getMaxCalories2() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return -1, errors.New("error reading from file" + err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	currentCalories := 0
	h := &IntHeap{}
	heap.Init(h)

	for scanner.Scan() {
		number := scanner.Text()
		if len(number) == 0 {
			heap.Push(h, currentCalories)
			if h.Len() > 3 {
				heap.Pop(h)
			}
			currentCalories = 0
			continue
		}
		calories, err := strconv.Atoi(number)
		if err != nil {
			return -1, errors.New("error converting calorie " + number)
		}
		currentCalories += calories
	}
	heap.Push(h, currentCalories)
	if h.Len() > 3 {
		heap.Pop(h)
	}
	if h.Len() != 3 {
		return -1, errors.New("There should only be 3 values in min heap at the end")
	}
	topThreeCalories := 0
	for h.Len() > 0 {
		topThreeCalories += (*h)[0]
		heap.Pop(h)
	}
	return topThreeCalories, nil
}

func getMaxCalories1() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return -1, errors.New("error reading from file" + err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	maxCalories, currentCalories := 0, 0

	for scanner.Scan() {
		number := scanner.Text()
		if len(number) == 0 {
			maxCalories = int(math.Max(float64(currentCalories), float64(maxCalories)))
			currentCalories = 0
			continue
		}
		calories, err := strconv.Atoi(number)
		if err != nil {
			return -1, errors.New("error converting calorie " + number)
		}
		currentCalories += calories
	}

	return maxCalories, nil
}

func main() {
	if maxCalories, err := getMaxCalories2(); err != nil {
		fmt.Println("there was an error", err)
	} else {
		fmt.Println(maxCalories)
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func getPriority(x uint8) int {
	if 97 <= x && x <= 122 {
		return int(1 + x - 97)
	}
	return int(1 + x - 65 + 26)
}

func getSumOfPriorities1() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return -1, errors.New("Unable to open file with error: " + err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	totalSum := 0
	for scanner.Scan() {
		items := scanner.Text()
		s := NewSet[uint8]()
		for i := 0; i < len(items)/2; i++ {
			s.Add(items[i])
		}
		for i := len(items) / 2; i < len(items); i++ {
			if s.Contains(items[i]) {
				val := getPriority(items[i])
				totalSum += val
				break
			}
		}
	}
	return totalSum, nil
}

func getSumOfPriorities2() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return -1, errors.New("Unable to open file with error: " + err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	totalSum := 0

	for scanner.Scan() {
		item1 := scanner.Text()
		scanner.Scan()
		item2 := scanner.Text()
		scanner.Scan()
		item3 := scanner.Text()
		s1, s2, s3 := NewSet[uint8](), NewSet[uint8](), NewSet[uint8]()
		for i := 0; i < len(item1); i++ {
			s1.Add(item1[i])
		}
		for i := 0; i < len(item2); i++ {
			s2.Add(item2[i])
		}
		for i := 0; i < len(item3); i++ {
			s3.Add(item3[i])
		}
		found := false
		// fmt.Println(item1, item2, item3)
		for i := uint8(97); i <= uint8(122); i++ {
			// fmt.Println(string(i), s1.Contains(i), s2.Contains(i), s3.Contains(i))
			if s1.Contains(i) && s2.Contains(i) && s3.Contains(i) {
				found = true
				totalSum += getPriority(i)
				break
			}
		}
		if found {
			continue
		}
		for i := uint8(65); i <= uint8(90); i++ {
			if s1.Contains(i) && s2.Contains(i) && s3.Contains(i) {
				found = true
				totalSum += getPriority(i)
				break
			}
		}
		if !found {
			return -1, errors.New("Didnt find any common items between the 3 elves")
		}
	}
	return totalSum, nil
}

func main() {
	if ans, err := getSumOfPriorities2(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}

type set[T comparable] struct {
	m map[T]bool
}

func NewSet[T comparable]() *set[T] {
	return &set[T]{m: map[T]bool{}}
}

func (s *set[T]) Contains(x T) bool {
	_, present := s.m[x]
	return present
}

func (s *set[T]) Add(x T) bool {
	if s.Contains(x) {
		return false
	}
	s.m[x] = true
	return true
}

func (s *set[T]) Len() int {
	return len(s.m)
}

func (s *set[T]) Erase(x T) bool {
	if !s.Contains(x) {
		return false
	}
	delete(s.m, x)
	return true
}

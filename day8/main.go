package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Solver struct {
	Matrix [][]int
}

func (s *Solver) Prepare(input string) {
	s.Parse(input)
}

func (s *Solver) Parse(filename string) {
	readFile, err := os.Open(filename)
	defer func() {
		_ = readFile.Close()
	}()

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	s.Matrix = make([][]int, 0, 1)

	for fileScanner.Scan() {
		row := make([]int, 0, 1)
		line := fileScanner.Text()
		for _, symbol := range line {
			row = append(row, int(symbol-'0'))
		}
		s.Matrix = append(s.Matrix, row)
	}
}

func (s *Solver) CheckCoordinates(x, y int) bool {
	width := len(s.Matrix[0])
	height := len(s.Matrix[1])
	j := y
	mark := true
	for i := 0; i < x; i++ {
		if s.Matrix[i][j] >= s.Matrix[x][y] {
			mark = false
			break
		}
	}
	if mark {
		return true
	}
	mark = true
	for i := x + 1; i < height; i++ {
		if s.Matrix[i][j] >= s.Matrix[x][y] {
			mark = false
			break
		}
	}
	if mark {
		return true
	}
	i := x
	mark = true
	for j := 0; j < y; j++ {
		if s.Matrix[i][j] >= s.Matrix[x][y] {
			mark = false
			break
		}
	}
	if mark {
		return true
	}
	mark = true
	for j := y + 1; j < width; j++ {
		if s.Matrix[i][j] >= s.Matrix[x][y] {
			mark = false
			break
		}
	}
	return mark
}

func (s *Solver) First() int {
	width := len(s.Matrix[0])
	height := len(s.Matrix[1])
	visibleCount := 2*(width+height) - 4

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if s.CheckCoordinates(x, y) {
				visibleCount++
			}
		}
	}

	return visibleCount
}

func (s *Solver) CalculateScore(x, y int) int {
	width := len(s.Matrix[0])
	height := len(s.Matrix[1])
	j := y
	top := x
	if x == 1 && y == 2 {
		fmt.Println()
	}
	for i := x - 1; i >= 0; i-- {
		if s.Matrix[i][j] >= s.Matrix[x][y] {
			top = x - i
			break
		}
	}
	bottom := height - x - 1
	for i := x + 1; i < height; i++ {
		if s.Matrix[i][j] >= s.Matrix[x][y] {
			bottom = i - x
			break
		}
	}
	i := x
	left := y
	for j := y - 1; j >= 0; j-- {
		if s.Matrix[i][j] >= s.Matrix[x][y] {
			left = y - j
			break
		}
	}
	right := width - y - 1
	for j := y + 1; j < width; j++ {
		if s.Matrix[i][j] >= s.Matrix[x][y] {
			right = j - y
			break
		}
	}
	return top * bottom * left * right
}

func (s *Solver) Second() int {
	width := len(s.Matrix[0])
	height := len(s.Matrix[1])
	score := 0

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {

			if currentScore := s.CalculateScore(x, y); currentScore > score {
				score = currentScore
			}
		}
	}

	return score
}

func main() {
	solver := Solver{}
	solver.Prepare("input")
	fmt.Println(solver.First())
	fmt.Println(solver.Second())
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Shape [][2]int

var shapes = [5]Shape{
	{ // horizontal line
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
	},
	{ // cross
		{3, 2},
		{2, 1},
		{3, 1},
		{4, 1},
		{3, 0},
	},
	{ // L inverted and referenced FROM THE BOTTOM
		{4, 2},
		{4, 1},
		{4, 0},
		{2, 0},
		{3, 0},
	},
	{ // vertical line
		{2, 3},
		{2, 2},
		{2, 1},
		{2, 0},
	},
	{ // square
		{2, 1},
		{3, 1},
		{2, 0},
		{3, 0},
	},
}

var heights = map[int]int{
	0: 1,
	1: 3,
	2: 3,
	3: 4,
	4: 2,
}

type Solver struct {
	filename     string
	Cavern       [][7]bool // the whole thing is upside-down, 0 row is bottom
	Instructions []rune
	highestFloor int
}

func (s *Solver) Parse() {
	readFile, err := os.Open(s.filename)
	defer func() {
		_ = readFile.Close()
	}()
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Scan()
	s.Instructions = []rune(fileScanner.Text())
}

func (s *Solver) moveIfPossible(shape Shape, xOffset, yOffset int) bool {
	for i := range shape {
		newX := shape[i][0] + xOffset
		newY := shape[i][1] + yOffset
		if newX < 0 || newX >= 7 || newY < 0 || s.Cavern[newY][newX] {
			return false
		}
	}
	// can place
	for i := range shape {
		shape[i][0] += xOffset
		shape[i][1] += yOffset
	}
	return true
}

func (s *Solver) finalizeShape(shape Shape) {
	for i := range shape {
		x, y := shape[i][0], shape[i][1]
		s.Cavern[y][x] = true
	}
}

func (s *Solver) getShape(number, highest int) Shape {
	originalShape := shapes[number]
	shape := make(Shape, len(originalShape))
	for i := range originalShape {
		shape[i][0] = originalShape[i][0]
		shape[i][1] = highest + 4 + originalShape[i][1]
	}
	return shape
}

func (s *Solver) PrintCurrentState(shape Shape) string {
	cavern := make([][]rune, 0, len(s.Cavern))
	for y := range s.Cavern {
		row := make([]rune, 7)
		for x := range s.Cavern[y] {
			if s.Cavern[y][x] {
				row[x] = '#'
			} else {
				row[x] = '.'
			}
			for _, point := range shape {
				if point[0] == x && point[1] == y {
					row[x] = '@'
					break
				}
			}
		}
		cavern = append(cavern, row)
	}
	builder := strings.Builder{}
	builder.WriteString("-------\n")
	for i := len(cavern) - 1; i >= 0; i-- {
		builder.WriteString(string(cavern[i]))
		builder.WriteRune('\n')
	}
	builder.WriteString("-------\n")
	return builder.String()
}

func (s *Solver) findHighest() int {
	for y := len(s.Cavern) - 1; y >= 0; y-- {
		for x := range s.Cavern[y] {
			if s.Cavern[y][x] {
				return y
			}
		}
	}
	return -1
}

func (s *Solver) shrink() {
	found := -1
	for y := len(s.Cavern) - 1; y >= 0; y-- {
		fullRow := true
		for x := range s.Cavern[y] {
			if !s.Cavern[y][x] {
				fullRow = false
				break
			}
		}
		if fullRow {
			found = y
			break
		}
	}
	if found == -1 {
		return
	}
	s.highestFloor += found
	s.Cavern = s.Cavern[found:]
}

func (s *Solver) runRounds(n int) int {
	currentShape := 0
	shape := s.getShape(currentShape, -1)
	//fmt.Println(s.PrintCurrentState(shape))
	stopped := 0
	for i := 0; stopped < n; i = (i + 1) % len(s.Instructions) {
		var xDir int
		switch s.Instructions[i] {
		case '>':
			xDir = 1
		case '<':
			xDir = -1
		default:
			panic("oh noes")
		}
		_ = s.moveIfPossible(shape, xDir, 0)
		moved := s.moveIfPossible(shape, 0, -1)
		if moved {
			//fmt.Println(s.PrintCurrentState(shape))
			continue
		}
		stopped++
		s.finalizeShape(shape)
		s.shrink()
		currentShape = (currentShape + 1) % 5
		highest := s.findHighest()
		howManyRowsToAdd := highest + 4 + heights[currentShape] - len(s.Cavern)
		if howManyRowsToAdd > 0 {
			s.Cavern = append(s.Cavern, make([][7]bool, howManyRowsToAdd)...)
		}
		shape = s.getShape(currentShape, highest)
		//fmt.Println(s.PrintCurrentState(shape))
	}
	return s.findHighest() + 1
}

func (s *Solver) First() int {
	s.Cavern = make([][7]bool, 4)
	return s.highestFloor + s.runRounds(2022)
}

func (s *Solver) Second() int {
	s.Cavern = make([][7]bool, 4)
	return s.runRounds(1_000_000_000_000)
}

func main() {
	solver := Solver{filename: "input"}
	solver.Parse()
	fmt.Println(solver.First())
	fmt.Println(solver.Second())
}

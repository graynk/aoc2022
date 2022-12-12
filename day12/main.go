package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
	filename string
	Map      [][]*Cell
	Traveler Traveler
}

type Traveler struct {
	start  *Cell
	target *Cell
}

type Cell struct {
	options []*Cell
	visited bool
	value   rune
	minPath int
}

var directions = [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func (s *Solver) Parse() {
	readFile, err := os.Open(s.filename)
	defer func() {
		_ = readFile.Close()
	}()
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	s.Map = make([][]*Cell, 0, 1)
	s.Traveler = Traveler{}
	row := -1
	for fileScanner.Scan() {
		row++
		s.Map = append(s.Map, make([]*Cell, 0, 1))
		line := fileScanner.Text()
		for col, symbol := range line {
			cell := Cell{
				value:   symbol,
				minPath: 5000,
			}
			s.Map[row] = append(s.Map[row], &cell)
			switch symbol {
			case 'S':
				s.Map[row][col].value = 'a'
				s.Traveler.start = s.Map[row][col]
			case 'E':
				s.Map[row][col].value = 'z'
				s.Traveler.target = s.Map[row][col]
			}
		}
	}
	height := len(s.Map)
	width := len(s.Map[0])
	for row := range s.Map {
		for col := range s.Map[row] {
			symbol := s.Map[row][col].value
			options := make([]*Cell, 0, 4)
			for _, direction := range directions {
				x, y := col+direction[0], row+direction[1]
				if x < 0 || y < 0 || x >= width || y >= height || (s.Map[y][x].value-symbol) > 1 {
					continue
				}
				options = append(options, s.Map[y][x])
			}
			s.Map[row][col].options = options
		}
	}
}

func (c *Cell) hasRouteTo(other *Cell) bool {
	for _, option := range c.options {
		if option == other {
			return true
		}
	}
	return false
}

func (t *Traveler) Traverse(s *Solver) int {
	cell := t.start
	cell.minPath = 0
	cell.visited = true
	options := []*Cell{cell.options[0]}
	cell = cell.options[1]

	for len(options) != 0 || !cell.visited {
		minPathUpdated := false
		for _, neighbour := range cell.options {
			sumRisk := neighbour.minPath + 1
			if cell == s.Traveler.target && neighbour.hasRouteTo(cell) {
				fmt.Println()
			}
			if neighbour.visited && sumRisk < cell.minPath && neighbour.hasRouteTo(cell) {
				cell.minPath = sumRisk
				minPathUpdated = true
			}
		}
		if !cell.visited || minPathUpdated {
			options = append(options, cell.options...)
		}

		cell.visited = true

		cell = options[0]
		options = options[1:]
	}

	return t.target.minPath
}

func (s *Solver) String() string {
	builder := strings.Builder{}
	builder.WriteRune('\n')
	for row := range s.Map {
		for col := range s.Map[row] {
			cell := s.Map[row][col]
			//value := string(cell.value)
			//if cell == s.Traveler.start {
			//	value = "s"
			//} else if cell == s.Traveler.target {
			//	value = "e"
			//}
			//if cell.visited {
			//	value = strings.ToUpper(value)
			//}
			builder.WriteString(strconv.Itoa(cell.minPath))
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func (s *Solver) First() int {
	return s.Traveler.Traverse(s)
}

func (s *Solver) Second() int {
	return 0
}

func main() {
	solver := Solver{filename: "input"}
	solver.Parse()
	fmt.Println(solver.First())
	fmt.Println(solver.Second())
}

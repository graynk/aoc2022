package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Solver struct {
	filename       string
	Map            [][]rune
	Visited        [][]bool
	startX, startY int
	endX, endY     int
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
	s.Map = make([][]rune, 0, 1)
	s.Visited = make([][]bool, 0, 1)
	s.startX = -1
	s.endX = -1
	for fileScanner.Scan() {
		line := fileScanner.Text()
		s.Map = append(s.Map, []rune(line))
		s.Visited = append(s.Visited, make([]bool, len(line)))
		if s.startX == -1 {
			startIndex := strings.IndexRune(line, 'S')
			if startIndex != -1 {
				s.startX = startIndex
				s.startY = len(s.Map) - 1
				s.Map[s.startY][s.startX] = 'a'
				s.Visited[s.startY][s.startX] = true
			}
		}
		if s.endX == -1 {
			endIndex := strings.IndexRune(line, 'E')
			if endIndex != -1 {
				s.endX = endIndex
				s.endY = len(s.Map) - 1
				s.Map[s.endY][s.endX] = 'z'
			}
		}
	}
}

func clone(initial [][2]int) [][2]int {
	cloned := make([][2]int, len(initial))
	copy(cloned, initial)
	return cloned
}

func coordInPath(x, y int, path [][2]int) bool {
	for _, coord := range path {
		if coord[0] == x && coord[1] == y {
			return true
		}
	}
	return false
}

func (s *Solver) canGo(x, y, newX, newY int, path [][2]int) bool {
	if newX < 0 || newY < 0 || newY >= len(s.Map) || newX >= len(s.Map[0]) || coordInPath(newX, newY, path) {
		return false
	}
	return (s.Map[newY][newX] - s.Map[y][x]) <= 1
}

func (s *Solver) Search(x, y int, path [][2]int) [][2]int {
	cloned := clone(path)
	cloned = append(cloned, [2]int{x, y})

	if x == s.endX && y == s.endY {
		return cloned
	}
	//s.Visited[y][x] = true

	paths := make([][][2]int, 0, 1)
	for _, direction := range directions {
		dirX, dirY := direction[0], direction[1]
		newX, newY := x+dirX, y+dirY
		if !s.canGo(x, y, newX, newY, cloned) {
			continue
		}
		result := s.Search(newX, newY, cloned)
		last := result[len(result)-1]
		if last[0] == s.endX && last[1] == s.endY {
			paths = append(paths, result)
		}
	}

	if len(paths) == 0 {
		return cloned
	}
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
	return paths[0]
}

func (s *Solver) First() int {
	shortest := s.Search(s.startX, s.startY, [][2]int{})
	return len(shortest) - 1
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

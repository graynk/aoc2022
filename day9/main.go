package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	X, Y int
}

type Instruction struct {
	Direction rune
	Distance  int
}

type Solver struct {
	filename     string
	Instructions []Instruction
	Visited      map[int]any
}

func (s *Solver) Parse() {
	readFile, err := os.Open(s.filename)
	defer func() {
		_ = readFile.Close()
	}()

	if err != nil {
		log.Fatal(err)
	}

	s.Instructions = make([]Instruction, 0, 1)
	s.Visited = make(map[int]any)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		split := strings.Split(fileScanner.Text(), " ")
		direction := rune(split[0][0])
		value, _ := strconv.Atoi(split[1])
		s.Instructions = append(s.Instructions, Instruction{
			Direction: direction,
			Distance:  value,
		})
	}
}

func abshahahaha(value int) int {
	if value < 0 {
		return -1 * value
	}
	return value
}

func catchUp(hX, hY, tX, tY int) (int, int) {
	diffX, diffY := abshahahaha(hX-tX), abshahahaha(hY-tY)
	switch {
	case diffX <= 1 && diffY <= 1:
		return tX, tY
	case hX == tX:
		fallthrough
	case diffX < diffY:
		if hY > tY {
			return hX, hY - 1
		}
		return hX, hY + 1
	case hY == tY:
		fallthrough
	case diffX > diffY:
		if hX > tX {
			return hX - 1, hY
		}
		return hX + 1, hY
	}

	panic("shouldn't be here")
}

func hashCoords(x, y int) int {
	//return (x * 0x1f1f1f1f) ^ y
	//return fmt.Sprintf("%d,%d", x, y)
	return (y << 16) ^ x
}

func (s *Solver) First() int {
	hX, hY, tX, tY := 0, 0, 0, 0
	for _, instruction := range s.Instructions {
		switch instruction.Direction {
		case 'U':
			for i := 0; i < instruction.Distance; i++ {
				hY++
				tX, tY = catchUp(hX, hY, tX, tY)
				s.Visited[hashCoords(tX, tY)] = nil
			}
		case 'D':
			for i := 0; i < instruction.Distance; i++ {
				hY--
				tX, tY = catchUp(hX, hY, tX, tY)
				s.Visited[hashCoords(tX, tY)] = nil
			}
		case 'R':
			for i := 0; i < instruction.Distance; i++ {
				hX++
				tX, tY = catchUp(hX, hY, tX, tY)
				s.Visited[hashCoords(tX, tY)] = nil
			}
		case 'L':
			for i := 0; i < instruction.Distance; i++ {
				hX--
				tX, tY = catchUp(hX, hY, tX, tY)
				s.Visited[hashCoords(tX, tY)] = nil
			}
		}
	}

	return len(s.Visited)
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

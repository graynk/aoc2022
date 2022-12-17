package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func catchUp(head, tail [2]int) [2]int {
	hX, hY := head[0], head[1]
	tX, tY := tail[0], tail[1]
	diffX, diffY := abshahahaha(hX-tX), abshahahaha(hY-tY)
	switch {
	case diffX <= 1 && diffY <= 1:
		return tail
	case hX == tX:
		fallthrough
	case diffX < diffY:
		if hY > tY {
			return [2]int{hX, hY - 1}
		}
		return [2]int{hX, hY + 1}
	case hY == tY:
		fallthrough
	case diffX > diffY:
		if hX > tX {
			return [2]int{hX - 1, hY}
		}
		return [2]int{hX + 1, hY}
	}

	panic("shouldn't be here")
}

func hashCoords(coords [2]int) int {
	//return (x * 0x1f1f1f1f) ^ y
	//return fmt.Sprintf("%d,%d", x, y)
	return (coords[1] << 16) ^ coords[0]
}

func (s *Solver) runTheRope(n int) int {
	rope := make([][2]int, n)
	for _, instruction := range s.Instructions {
		for i := 0; i < instruction.Distance; i++ {
			switch instruction.Direction {
			case 'U':
				rope[0][1]++
			case 'D':
				rope[0][1]--
			case 'L':
				rope[0][0]--
			case 'R':
				rope[0][0]++
			}
			for knot := range rope {
				rope[knot] = catchUp(rope[0], rope[knot])
			}
			s.Visited[hashCoords(rope[n-1])] = nil
		}
	}

	return len(s.Visited)
}

func (s *Solver) First() int {
	return s.runTheRope(2)
}

func (s *Solver) Second() int {
	return s.runTheRope(10)
}

func main() {
	solver := Solver{filename: "input"}
	solver.Parse()
	fmt.Println(solver.First())
	fmt.Println(solver.Second())
}

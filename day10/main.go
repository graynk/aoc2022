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
	X []int
}

func (s *Solver) Prepare(input string) {
	s.X = []int{1, 1}
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
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		last := s.X[len(s.X)-1]
		s.X = append(s.X, last) // in-progress
		line := fileScanner.Text()
		split := strings.Split(line, " ")
		if split[0] == "noop" {
			s.X = append(s.X, last) // after
			continue
		}
		s.X = append(s.X, last) // after
		s.X = append(s.X, last) // // second cycle: in-progress
		value, _ := strconv.Atoi(split[1])
		s.X = append(s.X, last+value) // after
	}
}

func (s *Solver) First() int {
	sum := 0
	for i := 20; i*2-1 < len(s.X); i += 40 {
		sum += i * s.X[i*2-1]
	}
	return sum
}

func (s *Solver) pixelInSprite(i int) bool {
	pixel := (i/2 - 1) % 40
	spritePosition := s.X[i]
	return pixel >= spritePosition-1 && pixel <= spritePosition+1
}

func (s *Solver) Second() string {
	builder := strings.Builder{}
	for i := 2; i < len(s.X); i += 2 {
		if s.pixelInSprite(i) {
			builder.WriteRune('#')
		} else {
			builder.WriteRune('.')
		}
		if (i/2)%40 == 0 {
			builder.WriteRune('\n')
		}
	}
	return builder.String()
}

func main() {
	solver := Solver{}
	solver.Prepare("input")
	fmt.Println(solver.First())
	fmt.Println(solver.Second())
}

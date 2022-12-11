package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	StartingItems = "  Starting items: "
	OperationNew  = "  Operation: new = "
)

type Operation struct {
	First    string
	Operator rune
	Second   string
}

type ThrowTo struct {
	TrueMonkey  int
	FalseMonkey int
}

type MonkeTurn struct {
	Items       []int
	TestDivisor int
	Operation   Operation
	ThrowTo     ThrowTo
}

type Solver struct {
	filename    string
	modulo      int
	Turns       []MonkeTurn
	Inspections []int
}

func (s *Solver) Parse() {
	readFile, err := os.Open(s.filename)
	defer func() {
		_ = readFile.Close()
	}()

	if err != nil {
		log.Fatal(err)
	}

	s.Turns = make([]MonkeTurn, 0, 1)
	s.Inspections = make([]int, 0, 1)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() { // Monke X
		monkeTurn := MonkeTurn{Items: make([]int, 0, 1)}
		line := fileScanner.Text()
		fileScanner.Scan() // Starting items:
		line = fileScanner.Text()
		worries := strings.Split(line[len(StartingItems):], ", ")
		for _, worry := range worries {
			value, _ := strconv.Atoi(worry)
			monkeTurn.Items = append(monkeTurn.Items, value)
		}
		fileScanner.Scan() // OperationNew:
		monkeTurn.Operation = parseOperation(fileScanner.Text()[len(OperationNew):])
		fileScanner.Scan() // Test:
		line = fileScanner.Text()
		monkeTurn.TestDivisor = parseLastNumberInLine(line)
		fileScanner.Scan() // If true: throw to monke X
		trueMonke := parseLastNumberInLine(fileScanner.Text())
		fileScanner.Scan() // If false: throw to monke Y
		falseMonke := parseLastNumberInLine(fileScanner.Text())
		monkeTurn.ThrowTo = ThrowTo{
			TrueMonkey:  trueMonke,
			FalseMonkey: falseMonke,
		}
		fileScanner.Scan() // empty divisor
		s.Turns = append(s.Turns, monkeTurn)
	}
}

func parseLastNumberInLine(line string) int {
	valueIndex := strings.LastIndexByte(line, ' ') + 1
	value, _ := strconv.Atoi(line[valueIndex:])
	return value
}

func (s *Solver) increaseInspections(monke, inspections int) {
	if len(s.Inspections) <= monke {
		grownSlice := make([]int, monke+1, monke+1)
		copy(grownSlice, s.Inspections)
		s.Inspections = grownSlice
	}
	s.Inspections[monke] += inspections
}

func parseOperation(operation string) Operation {
	split := strings.Split(operation, " ")
	return Operation{
		First:    split[0],
		Operator: rune(split[1][0]),
		Second:   split[2],
	}
}

func parseOperand(operand string, old int) int {
	var value int
	if operand == "old" {
		value = old
	} else {
		value, _ = strconv.Atoi(operand)
	}
	return value
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(ints ...int) int {
	a, b := ints[0], ints[1]
	result := a * b / GCD(a, b)
	rest := ints[2:]
	for i := 0; i < len(rest); i++ {
		result = LCM(result, rest[i])
	}

	return result
}

func (o Operation) Execute(old int) int {
	first, second := parseOperand(o.First, old), parseOperand(o.Second, old)
	if o.Operator == '+' {
		return first + second
	}
	return first * second
}

func (s *Solver) throwToMonke(monke int, item int) {
	s.Turns[monke].Items = append(s.Turns[monke].Items, item)
}

func (s *Solver) executeRound(worried bool) {
	for monke := range s.Turns {
		s.increaseInspections(monke, len(s.Turns[monke].Items))
		for _, worry := range s.Turns[monke].Items {
			newWorry := s.Turns[monke].Operation.Execute(worry) // Inspect
			if !worried {
				newWorry /= 3 // calm down
			} else {
				newWorry %= s.modulo
			}
			// test worry
			if newWorry%s.Turns[monke].TestDivisor == 0 {
				s.throwToMonke(s.Turns[monke].ThrowTo.TrueMonkey, newWorry)
			} else {
				s.throwToMonke(s.Turns[monke].ThrowTo.FalseMonkey, newWorry)
			}
		}
		s.Turns[monke].Items = s.Turns[monke].Items[:0]
	}
}

func (s *Solver) runRounds(n int, worried bool) {
	for i := 0; i < n; i++ {
		s.executeRound(worried)
	}
	sort.Slice(s.Inspections, func(i, j int) bool {
		return s.Inspections[i] > s.Inspections[j]
	})
}

func (s *Solver) First() int {
	s.Parse()
	s.runRounds(20, false)
	return s.Inspections[0] * s.Inspections[1]
}

func (s *Solver) Second() int {
	s.Parse()
	divisors := make([]int, 0, 1)
	for _, turn := range s.Turns {
		divisors = append(divisors, turn.TestDivisor)
	}
	s.modulo = LCM(divisors...)
	s.runRounds(10_000, true)
	return s.Inspections[0] * s.Inspections[1]
}

func main() {
	solver := Solver{filename: "input"}
	fmt.Println(solver.First())
	fmt.Println(solver.Second())
}

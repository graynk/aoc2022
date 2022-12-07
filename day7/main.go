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

type Solver struct {
	Sizes []int
	Root  *Tree
}

type Tree struct {
	Files map[string]*Tree
	Size  int
}

func NewTree(parent *Tree, size int) *Tree {
	tree := &Tree{
		Files: make(map[string]*Tree),
		Size:  size,
	}
	tree.Files[".."] = parent
	return tree
}

func (s *Solver) Prepare(input string) {
	s.Root = Parse(input)
	s.Sizes = make([]int, 0, 1)
	rootSize := s.CalculateAndCollectSize(s.Root)
	s.Sizes = append(s.Sizes, rootSize)
	sort.Ints(s.Sizes)
}

func Parse(filename string) *Tree {
	readFile, err := os.Open(filename)
	defer func() {
		_ = readFile.Close()
	}()

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	root := &Tree{
		Files: make(map[string]*Tree),
	}
	root.Files["/"] = root
	root.Files[".."] = root
	currentDir := root

	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, " ")
		isCmd := split[0] == "$"

		if isCmd {
			command := split[1]
			if command == "cd" {
				target := split[2]
				currentDir = currentDir.Files[target]
			}
			continue
		}

		name, size := parseFileInfo(line)
		currentDir.Files[name] = NewTree(currentDir, size)
	}
	return root
}

func parseFileInfo(fileInfo string) (string, int) {
	info := strings.Split(fileInfo, " ")
	size, _ := strconv.ParseInt(info[0], 10, 64)
	return info[1], int(size)
}

func (s *Solver) CalculateAndCollectSize(t *Tree) int {
	size := t.Size
	for key, file := range t.Files {
		if key == ".." || key == "/" {
			continue
		}
		if file.Size == 0 {
			dirSize := s.CalculateAndCollectSize(file)
			s.Sizes = append(s.Sizes, dirSize)
			size += dirSize
			continue
		}
		size += file.Size
	}
	return size
}

func (s *Solver) First() int {
	sum := 0
	for _, size := range s.Sizes {
		if size < 100_000 {
			sum += size
		}
	}
	return sum
}

func (s *Solver) Second() int {
	rootSize := s.Sizes[len(s.Sizes)-1]
	target := 30_000_000 - (70_000_000 - rootSize)
	index := sort.SearchInts(s.Sizes, target)
	return s.Sizes[index]
}

func main() {
	solver := Solver{}
	solver.Prepare("input")
	fmt.Println(solver.First())
	fmt.Println(solver.Second())
}

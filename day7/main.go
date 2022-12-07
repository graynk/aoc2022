package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func parseFileInfo(fileInfo string) (string, int) {
	info := strings.Split(fileInfo, " ")
	size, _ := strconv.ParseInt(info[0], 10, 64)
	return info[1], int(size)
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
			if split[1] == "cd" {
				currentDir = currentDir.Files[split[2]]
			}
			continue
		}

		name, size := parseFileInfo(line)
		currentDir.Files[name] = NewTree(currentDir, size)
	}
	return root
}

func (t *Tree) CalculateSizeButOnlyIfItsCool() (int, int) {
	size := t.Size
	currentTaskSize := 0
	for key, file := range t.Files {
		if key == ".." || key == "/" {
			continue
		}
		if file.Size == 0 {
			dirSize, taskSize := file.CalculateSizeButOnlyIfItsCool()
			if dirSize < 100_000 {
				currentTaskSize += dirSize
			}
			currentTaskSize += taskSize
			size += dirSize
			continue
		}
		size += file.Size
	}

	return size, currentTaskSize
}

func (t *Tree) CalculateSizeOfBigEnoughEghWillDoDirectory(target int) (int, int) {
	size := t.Size
	currentTaskSize := math.MaxInt
	for key, file := range t.Files {
		if key == ".." || key == "/" {
			continue
		}
		if file.Size == 0 {
			dirSize, taskSize := file.CalculateSizeOfBigEnoughEghWillDoDirectory(target)
			if taskSize >= target && taskSize < currentTaskSize {
				currentTaskSize = taskSize
			} else if dirSize >= target && dirSize < currentTaskSize {
				currentTaskSize = dirSize
			}
			size += dirSize
			continue
		}
		size += file.Size
	}

	return size, currentTaskSize
}

func (t *Tree) First() int {
	_, taskSize := t.CalculateSizeButOnlyIfItsCool()
	return taskSize
}

func (t *Tree) Second(rootSize int) int {
	_, taskSize := t.CalculateSizeOfBigEnoughEghWillDoDirectory(30_000_000 - (70_000_000 - rootSize))
	return taskSize
}

func main() {
	tree := Parse("input")
	rootSize, first := tree.CalculateSizeButOnlyIfItsCool()
	fmt.Println(first)
	fmt.Println(tree.Second(rootSize))
}

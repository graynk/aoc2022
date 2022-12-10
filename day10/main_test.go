package main

import (
	"testing"
)

var solver = Solver{}

func TestFirst(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"part 1 test",
			"test",
			13140,
		},
		{
			"part 1 full",
			"input",
			13220,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver.Prepare(tt.filename)
			if got := solver.First(); got != tt.want {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecond(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			"part 2 test",
			"test",
			`##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
`,
		},
		{
			"part 2 full",
			"input",
			`###..#..#..##..#..#.#..#.###..####.#..#.
#..#.#..#.#..#.#.#..#..#.#..#.#....#.#..
#..#.#..#.#..#.##...####.###..###..##...
###..#..#.####.#.#..#..#.#..#.#....#.#..
#.#..#..#.#..#.#.#..#..#.#..#.#....#.#..
#..#..##..#..#.#..#.#..#.###..####.#..#.
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver.Prepare(tt.filename)
			if got := solver.Second(); got != tt.want {
				t.Errorf("Second() = %v, want %v", got, tt.want)
			}
		})
	}
}

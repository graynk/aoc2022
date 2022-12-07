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
			95437,
		},
		{
			"part 1 full",
			"input",
			1232307,
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

func TestParse(t *testing.T) {

	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"root test",
			"test",
			48381165,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver.Prepare(tt.filename)
			if got := solver.Sizes[len(solver.Sizes)-1]; got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecond(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"part 2 test",
			"test",
			24933642,
		},
		{
			"part 2 full",
			"test",
			24933642,
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

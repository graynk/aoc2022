package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"part 1 test",
			"test",
			13,
		},
		{
			"part 1 full",
			"input",
			6314,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := Solver{filename: tt.filename}
			solver.Parse()
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
		want     int
	}{
		{
			"part 2 test",
			"test",
			36,
		},
		{
			"part 2 full",
			"input",
			6314,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := Solver{filename: tt.filename}
			solver.Parse()
			if got := solver.Second(); got != tt.want {
				t.Errorf("Second() = %v, want %v", got, tt.want)
			}
		})
	}
}

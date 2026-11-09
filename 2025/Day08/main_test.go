package main

import "testing"

func TestRun(t *testing.T) {
	Tests.RunT(t, run)
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		run(Puzzle)
	}
}

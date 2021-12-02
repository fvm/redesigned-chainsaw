package day02

import "testing"

func BenchmarkAll(b *testing.B) {
	benchmarks := []struct {
		name string
		f    func([]displacement) (int, error)
	}{
		{name: "partOne", f: solvePartOne},
		{name: "partTwo", f: solvePartTwo},
	}

	data, err := readInput("input")
	if err != nil {
		b.Error(err)
	}

	for _, bm := range benchmarks {
		b.Run(
			bm.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = bm.f(data)
				}
			},
		)
	}
}

func Benchmark_solvePartOne(b *testing.B) {
	data, err := readInput("input")
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		_, _ = solvePartOne(data)
	}
}
func Benchmark_solvePartTwo(b *testing.B) {
	data, err := readInput("input")
	if err != nil {
		b.Errorf("Motherplucking error: %s", err)
	}
	for i := 0; i < b.N; i++ {
		_, _ = solvePartTwo(data)
	}
}

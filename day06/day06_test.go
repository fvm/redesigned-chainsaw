package day06

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkAllDay06(b *testing.B) {
	population, err := readAndParseInputFile("input")
	if err != nil {
		b.Errorf("Bloop, bloop, fishy bench fuckup %s", err)
	}
	type args struct {
		population Population
		days       uint64
	}
	benchmarks := []struct {
		name string
		f    func(population Population, days uint64) uint64
		args args
	}{
		{
			name: "partOne",
			f:    solvePartOne,
			args: args{
				population: population,
				days:       80,
			},
		},
		{
			name: "partTwo",
			f:    solvePartTwo,
			args: args{
				population: population,
				days:       256,
			},
		},
	}
	for _, bm := range benchmarks {
		b.Run(
			bm.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = bm.f(bm.args.population, bm.args.days)
				}
			},
		)
	}
}
func BenchmarkAllDayForNDays(b *testing.B) {
	population, err := readAndParseInputFile("input")
	if err != nil {
		b.Errorf("Bloop, bloop, fishy bench fuckup %s", err)
	}
	type args struct {
		population Population
		days       uint64
	}
	benchmarks := []struct {
		name string
		f    func(population Population, days uint64) uint64
		args args
	}{
		{
			name: fmt.Sprintf("N=%d", 80),
			f:    solvePartOne,
			args: args{
				population: population,
				days:       80,
			},
		},
		{
			name: fmt.Sprintf("N=%d", 256),
			f:    solvePartTwo,
			args: args{
				population: population,
				days:       256,
			},
		},
		{
			name: fmt.Sprintf("N=%d", 951),
			f:    solvePartTwo,
			args: args{
				population: population,
				days:       951,
			},
		}, {
			name: fmt.Sprintf("N=%d", 952),
			f:    solvePartTwo,
			args: args{
				population: population,
				days:       952,
			},
		}, {
			name: fmt.Sprintf("N=%d", 20*365),
			f:    solvePartTwo,
			args: args{
				population: population,
				days:       20 * 365,
			},
		}, {
			name: fmt.Sprintf("N=%d", 100*365),
			f:    solvePartTwo,
			args: args{
				population: population,
				days:       100 * 365,
			},
		},
	}
	for _, bm := range benchmarks {
		p := uint64(0)
		b.Run(
			bm.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					p = bm.f(bm.args.population, bm.args.days)
				}
			},
		)
		b.Logf("Population size: %d", p)
	}
}

func TestPopulation_Tick(t *testing.T) {
	tests := []struct {
		name               string
		startingPopulation *Population
		targetPopulation   *Population
	}{
		{
			name: testing.CoverMode(),
			startingPopulation: &Population{
				adult:    []uint64{2, 3, 2, 0, 1},
				maturing: []uint64{1, 2},
			},
			targetPopulation: &Population{
				adult:    []uint64{3, 2, 0, 1, 3},
				maturing: []uint64{2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := &Population{
					adult:    tt.startingPopulation.adult,
					maturing: tt.startingPopulation.maturing,
				}
				p.Tick()
				if !reflect.DeepEqual(p.adult, tt.targetPopulation.adult) {
					t.Errorf("initFromState() got=%v, targetPopulation=%v", p.adult, tt.targetPopulation.adult)
				}
				if !reflect.DeepEqual(p.maturing, tt.targetPopulation.maturing) {
					t.Errorf("initFromState() got=%v, targetPopulation=%v", p.maturing, tt.targetPopulation.maturing)
				}
			},
		)
	}
}

func TestPopulation_initFromState(t *testing.T) {
	type args struct {
		states []uint64
	}
	tests := []struct {
		name             string
		population       *Population
		args             args
		wantErr          bool
		targetPopulation Population
	}{
		{
			name:       testing.CoverMode(),
			population: NewPopulation(7, 2),
			args: args{
				states: []uint64{1, 2, 1, 6, 0},
			},
			targetPopulation: Population{
				adult:    []uint64{1, 2, 1, 0, 0, 0, 1},
				maturing: []uint64{0, 0},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := &Population{
					adult:    tt.population.adult,
					maturing: tt.population.maturing,
				}
				if err := p.initFromState(tt.args.states); (err != nil) != tt.wantErr {
					t.Errorf("initFromState() error = %v, wantErr %v", err, tt.wantErr)
				}
				// 	Compare the values
				if !reflect.DeepEqual(p.adult, tt.targetPopulation.adult) {
					t.Errorf("initFromState() got=%v, targetPopulation=%v", p.adult, tt.targetPopulation.adult)
				}
				if !reflect.DeepEqual(p.maturing, tt.targetPopulation.maturing) {
					t.Errorf("initFromState() got=%v, targetPopulation=%v", p.maturing, tt.targetPopulation.maturing)
				}
			},
		)
	}
}

func Test_solvePartOne(t *testing.T) {
	type args struct {
		population Population
		days       uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: testing.CoverMode(),
			args: args{
				population: Population{
					adult:    []uint64{0, 1, 1, 2, 1, 0, 0},
					maturing: []uint64{0, 0},
				},
				days: 80,
			}, want: 5934,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := solvePartOne(tt.args.population, tt.args.days); got != tt.want {
					t.Errorf("solvePartOne() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

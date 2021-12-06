package day06

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func BenchmarkAllDay06(b *testing.B) {
	t1 := time.Now()
	t2 := t1.AddDate(20, 0, 0)
	days := t2.Sub(t1).Hours() * 24
	population, err := readAndParseInputFile("input")
	if err != nil {
		b.Errorf("Bloop, bloop, fishy bench fuckup %s", err)
	}
	type args struct {
		population Population
		days       int
	}
	benchmarks := []struct {
		name string
		f    func(population Population, days int) int
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
		{
			name: fmt.Sprintf("%d days", int(days)),
			f:    solvePartTwo,
			args: args{
				population: population,
				days:       int(days),
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

func TestPopulation_Tick(t *testing.T) {
	tests := []struct {
		name               string
		startingPopulation *Population
		targetPopulation   *Population
	}{
		{
			name: testing.CoverMode(),
			startingPopulation: &Population{
				adult:    []int{2, 3, 2, 0, 1},
				maturing: []int{1, 2},
			},
			targetPopulation: &Population{
				adult:    []int{3, 2, 0, 1, 3},
				maturing: []int{2, 2},
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
		states []int
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
				states: []int{1, 2, 1, 6, 0},
			},
			targetPopulation: Population{
				adult:    []int{1, 2, 1, 0, 0, 0, 1},
				maturing: []int{0, 0},
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
		days       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: testing.CoverMode(),
			args: args{
				population: Population{
					adult:    []int{0, 1, 1, 2, 1, 0, 0},
					maturing: []int{0, 0},
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

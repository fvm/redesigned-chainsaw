package day07

import (
	"io"
	"math"
	"reflect"
	"strings"
	"testing"
)

func BenchmarkAllDay07(b *testing.B) {
	puzzleInput, err := readAndParseInputFile("input")
	if err != nil {
		b.Errorf("Bloop, bloop, crappy bench %s", err)
	}
	benchmarks := []struct {
		name string
		f    func(input PuzzleInput) (int, error)
	}{
		{name: "partOne", f: solvePartOne},
		{name: "partTwo", f: solvePartTwo},
	}
	for _, bm := range benchmarks {
		b.Run(
			bm.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = bm.f(puzzleInput)
				}
			},
		)
	}
}

func Test_solvePartOne(t *testing.T) {
	type args struct {
		input PuzzleInput
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: testing.CoverMode(),
			args: args{
				input: PuzzleInput{positions: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}},
			}, want: 37, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := solvePartOne(tt.args.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("solvePartOne() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("solvePartOne() got = %v, wantLo %v", got, tt.want)
				}
			},
		)
	}
}

func Test_solvePartTwo(t *testing.T) {
	type args struct {
		input PuzzleInput
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: testing.CoverMode(),
			args: args{
				input: PuzzleInput{positions: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}},
			}, want: 168, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := solvePartTwo(tt.args.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("solvePartTwo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("solvePartTwo() got = %v, wantLo %v", got, tt.want)
				}
			},
		)
	}
}

func Test_getRange(t *testing.T) {
	type args struct {
		domain []int
	}
	tests := []struct {
		name   string
		args   args
		wantLo int
		wantHi int
	}{
		{
			name: testing.CoverMode(),
			args: args{
				domain: []int{1, 2, 6, -1},
			},
			wantLo: -1,
			wantHi: 6,
		}, {
			name: testing.CoverMode(),
			args: args{
				domain: []int{0},
			},
			wantLo: 0,
			wantHi: 0,
		}, {
			name: testing.CoverMode(),
			args: args{
				domain: []int{math.MinInt, math.MaxInt},
			},
			wantLo: math.MinInt,
			wantHi: math.MaxInt,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, got1 := getRange(tt.args.domain)
				if got != tt.wantLo {
					t.Errorf("getRange() got = %v, wantLo %v", got, tt.wantLo)
				}
				if got1 != tt.wantHi {
					t.Errorf("getRange() got1 = %v, wantLo %v", got1, tt.wantHi)
				}
			},
		)
	}
}

func Test_parseInput(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    PuzzleInput
		wantErr bool
	}{
		{
			name: testing.CoverMode(),
			args: args{
				input: strings.NewReader("1,2,3,4,5,6"),
			},
			want: PuzzleInput{
				positions: []int{1, 2, 3, 4, 5, 6},
			},
			wantErr: false,
		}, {
			name: testing.CoverMode(),
			args: args{
				input: strings.NewReader(",;"),
			},
			want: PuzzleInput{
				positions: []int{},
			},
			wantErr: true,
		}, {
			name: testing.CoverMode(),
			args: args{
				input: strings.NewReader(","),
			},
			want: PuzzleInput{
				positions: []int{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := parseInput(tt.args.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr {
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("parseInput() got = %v, want %v", got, tt.want)
					}
				}
			},
		)
	}
}

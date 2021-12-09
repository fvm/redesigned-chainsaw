package day08

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_parseInput(t *testing.T) {

	type args struct {
		rc io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    PuzzleInput
		wantErr bool
	}{
		{
			name: testing.CoverMode(),
			args: args{rc: strings.NewReader("be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe")},
			want: PuzzleInput{
				displayLines: []DisplayLine{
					{
						inputFields: []string{
							"be",
							"cfbegad",
							"cbdgef",
							"fgaecd",
							"cgeb",
							"fdcge",
							"agebfd",
							"fecdb",
							"fabcd",
							"edb",
						},
						displayFields: []string{
							"fdgacbe",
							"cefdb",
							"cefbgd",
							"gcbe",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := parseInput(tt.args.rc)
				if (err != nil) != tt.wantErr {
					t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("parseInput() got = %v, want %v", got, tt.want)
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
				input: PuzzleInput{
					displayLines: []DisplayLine{
						{
							inputFields: []string{
								"be",
								"cfbegad",
								"cbdgef",
								"fgaecd",
								"cgeb",
								"fdcge",
								"agebfd",
								"fecdb",
								"fabcd",
								"edb",
							},
							displayFields: []string{
								"fdgacbe",
								"cefdb",
								"cefbgd",
								"gcbe",
							},
						},
					},
				},
			},
			want:    1,
			wantErr: false,
		}, {
			name: testing.CoverMode(),
			args: args{
				input: PuzzleInput{
					displayLines: []DisplayLine{
						{
							inputFields: []string{
								"be",
								"cfbegad",
								"cbdgef",
								"fgaecd",
								"cgeb",
								"fdcge",
								"agebfd",
								"fecdb",
								"fabcd",
								"edb",
							},
							displayFields: []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
						}, {
							inputFields: []string{
								"dbcfeag",
								"gecda",
								"bed",
								"eb",
								"dfcab",
								"abeg",
								"agfced",
								"dcebfg",
								"adgceb",
								"abecd",
							},
							displayFields: []string{"be", "gabe", "bed", "fbdac"},
						},
					},
				},
			},
			want:    4,
			wantErr: false,
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
					t.Errorf("solvePartOne() got = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
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
					t.Errorf("solvePartTwo() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

package day07

import "testing"

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
					t.Errorf("solvePartTwo() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

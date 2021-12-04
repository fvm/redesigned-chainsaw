package day04

import "testing"

func BenchmarkAllDay04(b *testing.B) {
	benchmarks := []struct {
		name string
		f    func([]Board, []int) (int, error)
	}{
		{name: "partOne", f: solvePartOne},
		{name: "partTwo", f: solvePartTwo},
	}
	boards, calledNumbers, err := readAndParseInputFile("input")
	if err != nil {
		b.Errorf("Bench the fuck out of here buddy! %s", err)
	}

	for _, bm := range benchmarks {
		b.Run(
			bm.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = bm.f(boards, calledNumbers)

				}
			},
		)
	}
}

func Test_solvePartTwo(t *testing.T) {
	tBoards, tNumbers, err := readAndParseInputFile("test_input")
	if err != nil {
		t.Error(err)
	}
	type args struct {
		boards        []Board
		calledNumbers []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "BingoBangoBongo",
			args: args{
				boards:        tBoards,
				calledNumbers: tNumbers,
			},
			want:    1924,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := solvePartTwo(tt.args.boards, tt.args.calledNumbers)
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

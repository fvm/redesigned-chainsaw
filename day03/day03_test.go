package day03

import (
	"reflect"
	"testing"
)

func BenchmarkAllDay03(b *testing.B) {
	benchmarks := []struct {
		name string
		f    func([]string) (int, error)
	}{
		{name: "partOne", f: solvePartOne},
		{name: "partTwo", f: solvePartTwo},
	}
	data, err := readInput("input")
	if err != nil {
		b.Errorf("Bench the fuck out of here buddy! %s", err)
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

func Test_getCommonBitsies(t *testing.T) {
	type args struct {
		d []string
	}
	tests := []struct {
		name      string
		args      args
		wantMost  string
		wantLeast string
		wantErr   bool
	}{
		{
			name: testing.CoverMode(),
			args: args{
				d: []string{"011110011100", "010001010101", "111111110000", "011101100011", "000111100100"},
			},
			wantMost:  "10110",
			wantLeast: "01001",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				gotMost, gotLeast, err := getCommonBitsies(tt.args.d)
				if (err != nil) != tt.wantErr {
					t.Errorf("getCommonBitsies() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotMost != tt.wantMost {
					t.Errorf("getCommonBitsies() gotMost = %v, wantMost %v", gotMost, tt.wantMost)
				}
				if gotLeast != tt.wantLeast {
					t.Errorf("getCommonBitsies() gotLeast = %v, wantMost %v", gotLeast, tt.wantLeast)
				}
			},
		)
	}
}

// func Test_doThatWeirdFilterShit(t *testing.T) {
// 	type args struct {
// 		f flippedFloppedStrings
// 	}
// 	tests := []struct {
// 		name      string
// 		args      args
// 		wantMost  string
// 		wantLeast string
// 		wantErr   bool
// 	}{
// 		{
// 			name: testing.CoverMode(), args: args{
// 			f: flippedFloppedStrings{
// 				transposed: []string{"011110011100", "010001010101", "111111110000", "011101100011", "000111100100"},
// 				original: []string{
// 					"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010",
// 				},
// 			},
// 		},
// 			wantMost:  "10111",
// 			wantLeast: "01010",
// 			wantErr:   false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(
// 			tt.name, func(t *testing.T) {
// 				gotMost, gotLeast, err := doThatWeirdFilterShit(tt.args.f)
// 				if (err != nil) != tt.wantErr {
// 					t.Errorf("doThatWeirdFilterShit() error = %v, wantErr %v", err, tt.wantErr)
// 					return
// 				}
// 				if gotMost != tt.wantMost {
// 					t.Errorf("doThatWeirdFilterShit() gotMost = %v, wantMost %v", gotMost, tt.wantMost)
// 				}
// 				if gotLeast != tt.wantLeast {
// 					t.Errorf("doThatWeirdFilterShit() gotLeast = %v, wantMost %v", gotLeast, tt.wantLeast)
// 				}
// 			},
// 		)
// 	}
// }

func Test_transpose(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: testing.CoverMode(), args: args{
				input: []string{"011110011100", "010001010101", "111111110000", "011101100011", "000111100100"},
			},
			want: []string{
				"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010",
			}, wantErr: false,
		}, {
			name: testing.CoverMode(), args: args{
				input: []string{
					"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010",
				},
			},
			want:    []string{"011110011100", "010001010101", "111111110000", "011101100011", "000111100100"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := transpose(tt.args.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("transpose() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("transpose() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_okayOnceMoreWithFeeling(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name     string
		args     args
		wantHigh string
		wantLow  string
		wantErr  bool
	}{
		{
			name: testing.CoverMode(),
			args: args{
				data: []string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
			},
			wantHigh: "10111",
			wantLow:  "01010",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				gotHigh, gotLow, err := okayOnceMoreWithFeeling(tt.args.data)
				if (err != nil) != tt.wantErr {
					t.Errorf("okayOnceMoreWithFeeling() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotHigh != tt.wantHigh {
					t.Errorf("okayOnceMoreWithFeeling() gotHigh = %v, wantHigh %v", gotHigh, tt.wantHigh)
				}
				if gotLow != tt.wantLow {
					t.Errorf("okayOnceMoreWithFeeling() gotLow = %v, wantLow %v", gotLow, tt.wantLow)
				}
			},
		)
	}
}

func Test_reduceDataByTarget(t *testing.T) {
	type args struct {
		input    []string
		target   string
		position int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: testing.CoverMode(), args: args{
				input:    []string{"10", "10", "01", "01"},
				target:   "1",
				position: 0,
			},
			want:    []string{"10", "10"},
			wantErr: false,
		}, {
			name: testing.CoverMode(), args: args{
				input:    []string{"00", "00", "00", "01"},
				target:   "1",
				position: 1,
			},
			want:    []string{"01"},
			wantErr: false,
		}, {
			name: testing.CoverMode(), args: args{
				input:    []string{"001", "011", "011", "011"},
				target:   "1",
				position: 1,
			},
			want:    []string{"011", "011", "011"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := reduceDataByTarget(tt.args.input, tt.args.target, tt.args.position)
				if (err != nil) != tt.wantErr {
					t.Errorf("reduceDataByTarget() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("reduceDataByTarget() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

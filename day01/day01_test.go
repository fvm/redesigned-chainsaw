package day01

import "testing"

func BenchmarkAll(b *testing.B) {
	benchmarks := []struct {
		name string
		f    func([]int) (int, error)
	}{
		{name: "partOne", f: solvePartOne},
		{name: "partTwo", f: solvePartTwo},
	}

	data, err := ReadInput("input")
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

func BenchmarkSolvePartOne(b *testing.B) {
	data, err := ReadInput("input")
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		_, _ = solvePartOne(data)
	}
}
func BenchmarkSolvePartTwo(b *testing.B) {
	data, err := ReadInput("input")
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		_, _ = solvePartTwo(data)
	}
}

func TestCountIncrements(t *testing.T) {
	type args struct {
		values []int
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: testing.CoverMode(), args: args{values: []int{1, 1, 2, 2}}, want: 1, wantErr: false},
		{name: testing.CoverMode(), args: args{values: []int{1, 1, 1, 1, 1}}, want: 0, wantErr: false},
		{name: testing.CoverMode(), args: args{values: []int{0, -1, -2, -3, -4}}, want: 0, wantErr: false},
		{name: testing.CoverMode(), args: args{values: []int{0, 1, 2, 3, 4}}, want: 4, wantErr: false},
		{
			name: testing.CoverMode(), args: args{values: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}}, want: 7,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := CountSingularIncrements(tt.args.values)
				if (err != nil) != tt.wantErr {
					t.Errorf("CountSingularIncrements() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("CountSingularIncrements() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestCountSubsliceIncrements(t *testing.T) {
	type args struct {
		input      []int
		windowsize int
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
				input:      nil,
				windowsize: 0,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: testing.CoverMode(),
			args: args{
				input:      []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
				windowsize: 3,
			},
			want:    5,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := CountSubsliceIncrements(tt.args.input, tt.args.windowsize)
				if (err != nil) != tt.wantErr {
					t.Errorf("CountSubsliceIncrements() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("CountSubsliceIncrements() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

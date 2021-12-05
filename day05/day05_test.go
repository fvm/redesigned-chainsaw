package day05

import (
	"fmt"
	"image"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "SolveDay05", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if err := Solve(); (err != nil) != tt.wantErr {
					t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

// .......1..
// ..1....1..
// ..1....1..
// .......1..
// .112111211
// ..........
// ..........
// ..........
// ..........
// 222111....
func Test_interpolate(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: testing.CoverMode(), args: args{
				min: 2,
				max: 5,
			},
			want: []int{2, 3, 4, 5},
		}, {
			name: testing.CoverMode(), args: args{
				min: 8,
				max: 2,
			}, want: []int{8, 7, 6, 5, 4, 3, 2},
		}, {
			name: testing.CoverMode(), args: args{
				min: 0,
				max: 0,
			}, want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := interpolate(tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("interpolate() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_drawLine(t *testing.T) {

	type args struct {
		l      ManhattanLine
		canvas [][]int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name: testing.CoverMode(), args: args{
				l:      NewManhattanLine(0, 1, 0, 3),
				canvas: createCanvas(5, 5),
			},
			want: [][]int{
				{0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			wantErr: false,
		}, {
			name: testing.CoverMode(), args: args{
				l: NewManhattanLine(0, 1, 3, 1),
				canvas: [][]int{
					{0, 0, 0, 0, 0},
					{1, 0, 0, 0, 0},
					{1, 0, 0, 0, 0},
					{1, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0, 0, 0},
				{2, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			wantErr: false,
		}, {
			name: testing.CoverMode(), args: args{
				l: NewManhattanLine(4, 0, 4, 4),
				canvas: [][]int{
					{0, 0, 0, 0, 0},
					{2, 1, 1, 1, 0},
					{1, 0, 0, 0, 0},
					{1, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0, 0, 1},
				{2, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 0, 0, 1},
				{0, 0, 0, 0, 1},
			},
			wantErr: false,
		}, {
			name: testing.CoverMode(), args: args{
				l: NewManhattanLine(0, 0, 0, 4),
				canvas: [][]int{
					{0, 0, 0, 0, 1},
					{2, 1, 1, 1, 1},
					{1, 0, 0, 0, 1},
					{1, 0, 0, 0, 1},
					{0, 0, 0, 0, 1},
				},
			},
			want: [][]int{
				{1, 0, 0, 0, 1},
				{3, 1, 1, 1, 1},
				{2, 0, 0, 0, 1},
				{2, 0, 0, 0, 1},
				{1, 0, 0, 0, 1},
			},
			wantErr: false,
		}, {
			name: testing.CoverMode(),
			args: args{
				l: NewManhattanLine(0, 0, 4, 4),
				canvas: [][]int{
					{0, 1, 1, 1, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
			},
			want: [][]int{
				{1, 1, 1, 1, 0},
				{0, 1, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 0},
				{0, 0, 0, 0, 1},
			},
			wantErr: false,
		}, {
			name: testing.CoverMode(),
			args: args{
				l: NewManhattanLine(4, 0, 0, 4),
				canvas: [][]int{
					{0, 1, 1, 1, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
			},
			want: [][]int{
				{0, 1, 1, 1, 1},
				{0, 0, 0, 1, 0},
				{0, 0, 1, 0, 0},
				{0, 1, 0, 0, 0},
				{1, 0, 0, 0, 0},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := drawLine(tt.args.l, tt.args.canvas)
				if (err != nil) != tt.wantErr {
					t.Errorf("drawLine() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					if !reflect.DeepEqual(got, tt.want) {
						bldGot := strings.Builder{}
						for _, l := range got {
							bldGot.WriteString(fmt.Sprintln(l))
						}
						bldWant := strings.Builder{}
						for _, l := range tt.want {
							bldWant.WriteString(fmt.Sprintln(l))
						}
						t.Errorf("drawLine()\ngot:\n%v\nwant\n%v\n", bldGot.String(), bldWant.String())
					}
				}
			},
		)
	}
}

func Test_countIntersections(t *testing.T) {
	type args struct {
		canvas [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: testing.CoverMode(),
			args: args{
				canvas: [][]int{
					{0, 2, 1, 1, 0},
					{0, 1, 0, 0, 0},
					{0, 1, 0, 0, 0},
					{0, 1, 0, 0, 0},
					{1, 1, 1, 1, 1},
				},
			},
			want: 1,
		}, {
			name: testing.CoverMode(),
			args: args{
				canvas: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
					{0, 1, 1, 2, 1, 1, 1, 2, 1, 1},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
				},
			},
			want: 5,
		}, {
			name: testing.CoverMode(),
			args: args{
				canvas: [][]int{
					{1, 0, 1, 0, 0, 0, 0, 1, 1, 0},
					{0, 1, 1, 1, 0, 0, 0, 2, 0, 0},
					{0, 0, 2, 0, 1, 0, 1, 1, 1, 0},
					{0, 0, 0, 1, 0, 2, 0, 2, 0, 0},
					{0, 1, 1, 2, 3, 1, 3, 2, 1, 1},
					{0, 0, 0, 1, 0, 2, 0, 0, 0, 0},
					{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
					{0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
					{1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
					{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
				},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := countIntersections(tt.args.canvas); got != tt.want {
					t.Errorf("countIntersections() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_parseInput(t *testing.T) {
	rdr, err := readerFromFileContents("test_input")
	if err != nil {
		t.Errorf("Can't open the bleeping test file: %s", err.Error())
	}

	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []ManhattanLine
		wantErr bool
	}{
		{
			name: testing.CoverMode(),
			args: args{input: rdr},
			want: []ManhattanLine{
				NewManhattanLine(0, 9, 5, 9),
				NewManhattanLine(8, 0, 0, 8),
				NewManhattanLine(9, 4, 3, 4),
				NewManhattanLine(2, 2, 2, 1),
				NewManhattanLine(7, 0, 7, 4),
				NewManhattanLine(6, 4, 2, 0),
				NewManhattanLine(0, 9, 2, 9),
				NewManhattanLine(3, 4, 1, 4),
				NewManhattanLine(0, 0, 8, 8),
				NewManhattanLine(5, 5, 8, 2),
			},
			wantErr: false,
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
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("parseInput() got =\n%v\nwant\n%v\n", got, tt.want)
				}
			},
		)
	}
}

func Test_isManhattanish(t *testing.T) {
	type args struct {
		x0 int
		y0 int
		x1 int
		y1 int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Diagonal Up", args: args{
				x0: 8,
				y0: 0,
				x1: 0,
				y1: 8,
			}, want: true,
		}, {
			name: "Diagonal Down", args: args{
				x0: 0,
				y0: 0,
				x1: 8,
				y1: 8,
			}, want: true,
		}, {
			name: "Point", args: args{
				x0: 8,
				y0: 0,
				x1: 8,
				y1: 0,
			}, want: true,
		}, {
			name: "Horizontal", args: args{
				x0: 3,
				y0: 1,
				x1: 3,
				y1: 8,
			}, want: true,
		}, {
			name: "Vertical", args: args{
				x0: 3,
				y0: 1,
				x1: 8,
				y1: 1,
			}, want: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := isManhattanish(tt.args.x0, tt.args.y0, tt.args.x1, tt.args.y1); got != tt.want {
					t.Errorf("isManhattan() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestNewManhattanLine(t *testing.T) {
	type args struct {
		x0 int
		y0 int
		x1 int
		y1 int
	}
	tests := []struct {
		name string
		args args
		want ManhattanLine
	}{
		{
			name: testing.CoverMode(), args: args{
				x0: 0,
				y0: 0,
				x1: 0,
				y1: 2,
			},
			want: ManhattanLine{
				ish:     false,
				Points:  []image.Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}},
				XValues: []int{0, 0, 0},
				YValues: []int{0, 1, 2},
			},
		}, {
			name: testing.CoverMode(), args: args{
				x0: 0,
				y0: 0,
				x1: 2,
				y1: 2,
			},
			want: ManhattanLine{
				ish:     true,
				Points:  []image.Point{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}},
				XValues: []int{0, 1, 2},
				YValues: []int{0, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NewManhattanLine(tt.args.x0, tt.args.y0, tt.args.x1, tt.args.y1); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewManhattanLine() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_drawManhattanishLines(t *testing.T) {
	type args struct {
		manhattanLines []ManhattanLine
		canvas         [][]int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name: testing.CoverMode(),
			args: args{
				manhattanLines: []ManhattanLine{
					NewManhattanLine(0, 9, 5, 9),
					NewManhattanLine(8, 0, 0, 8),
					NewManhattanLine(9, 4, 3, 4),
					NewManhattanLine(2, 2, 2, 1),
					NewManhattanLine(7, 0, 7, 4),
					NewManhattanLine(6, 4, 2, 0),
					NewManhattanLine(0, 9, 2, 9),
					NewManhattanLine(3, 4, 1, 4),
					NewManhattanLine(0, 0, 8, 8),
					NewManhattanLine(5, 5, 8, 2),
				},
				canvas: createCanvas(10, 10),
			},
			want: [][]int{
				{1, 0, 1, 0, 0, 0, 0, 1, 1, 0},
				{0, 1, 1, 1, 0, 0, 0, 2, 0, 0},
				{0, 0, 2, 0, 1, 0, 1, 1, 1, 0},
				{0, 0, 0, 1, 0, 2, 0, 2, 0, 0},
				{0, 1, 1, 2, 3, 1, 3, 2, 1, 1},
				{0, 0, 0, 1, 0, 2, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
				{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := drawManhattanishLines(tt.args.manhattanLines, tt.args.canvas)
				if (err != nil) != tt.wantErr {
					t.Errorf("drawManhattanishLines() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					bldGot := strings.Builder{}
					for _, l := range got {
						bldGot.WriteString(fmt.Sprintln(l))
					}
					bldWant := strings.Builder{}
					for _, l := range tt.want {
						bldWant.WriteString(fmt.Sprintln(l))
					}
					t.Errorf("drawManhattanishLines()\ngot:\n%v\nwant\n%v\n", bldGot.String(), bldWant.String())
				}
			},
		)
	}
}
func Test_drawManhattanLines(t *testing.T) {
	type args struct {
		manhattanLines []ManhattanLine
		canvas         [][]int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name: testing.CoverMode(),
			args: args{
				manhattanLines: []ManhattanLine{
					NewManhattanLine(0, 9, 5, 9),
					NewManhattanLine(8, 0, 0, 8),
					NewManhattanLine(9, 4, 3, 4),
					NewManhattanLine(2, 2, 2, 1),
					NewManhattanLine(7, 0, 7, 4),
					NewManhattanLine(6, 4, 2, 0),
					NewManhattanLine(0, 9, 2, 9),
					NewManhattanLine(3, 4, 1, 4),
					NewManhattanLine(0, 0, 8, 8),
					NewManhattanLine(5, 5, 8, 2),
				},
				canvas: createCanvas(10, 10),
			},
			want: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 1, 1, 2, 1, 1, 1, 2, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := drawManhattanLines(tt.args.manhattanLines, tt.args.canvas)
				if (err != nil) != tt.wantErr {
					t.Errorf("drawManhattanLines() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					bldGot := strings.Builder{}
					for _, l := range got {
						bldGot.WriteString(fmt.Sprintln(l))
					}
					bldWant := strings.Builder{}
					for _, l := range tt.want {
						bldWant.WriteString(fmt.Sprintln(l))
					}
					t.Errorf("drawManhattanLines()\ngot:\n%v\nwant\n%v\n", bldGot.String(), bldWant.String())
				}
			},
		)
	}
}

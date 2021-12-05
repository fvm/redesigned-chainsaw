package day05

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"reflect"
	"strings"
	"testing"
)

func BenchmarkDrawManhattanLines(b *testing.B) {
	manhattanLines, err := readAndParseInputFile("input")
	if err != nil {
		b.Errorf("Bench the fuck out of here motherbucker! %s", err)
	}

	benchmarks := []struct {
		name   string
		lines  []ManhattanLine
		canvas *Canvas
		f      func()
	}{
		{name: "Manhattan", lines: manhattanLines, canvas: newCanvas(1000)},
	}
	for _, bm := range benchmarks {
		b.Run(
			bm.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = bm.canvas.drawManhattanLines(bm.lines)
				}
			},
		)
	}
}
func BenchmarkDrawManhattanishLines(b *testing.B) {
	manhattanLines, err := readAndParseInputFile("input")
	if err != nil {
		b.Errorf("Bench the fuck out of here motherbucker! %s", err)
	}

	benchmarks := []struct {
		name   string
		lines  []ManhattanLine
		canvas *Canvas
		f      func()
	}{
		{name: "Manhattan", lines: manhattanLines, canvas: newCanvas(1000)},
	}
	for _, bm := range benchmarks {
		b.Run(
			bm.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = bm.canvas.drawManhattanishLines(bm.lines)
				}
			},
		)
	}
}

func BenchmarkAllDay05(b *testing.B) {
	lines, err := readAndParseInputFile("input")
	if err != nil {
		b.Errorf("Bench the fuck out of here motherbucker! %s", err)
	}
	benchmarks := []struct {
		name string
		f    func([]ManhattanLine) (int, error)
	}{
		{name: "partOne", f: solvePartOne},
		{name: "partTwo", f: solvePartTwo},
	}
	for _, bm := range benchmarks {
		b.Run(
			bm.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = bm.f(lines)
				}
			},
		)
	}
}

func BenchmarkNewManhattanLine(b *testing.B) {
	rdr, err := readerFromFileContents("input")
	if err != nil {
		b.Errorf("Bench the fuck out of here motherbucker! %s", err)
	}
	type arg struct {
		x0 int
		y0 int
		x1 int
		y1 int
	}
	var args []arg
	s := bufio.NewScanner(rdr)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		a := arg{
			x0: 0,
			y0: 0,
			x1: 0,
			y1: 0,
		}
		line := s.Text()
		_, err := fmt.Fscanf(strings.NewReader(line), "%d,%d -> %d,%d", &a.x0, &a.y0, &a.x1, &a.y1)
		if err != nil {
			b.Errorf("Bench the fuck out of here motherbucker! %s", err)
		}
	}
	benchmarks := []struct {
		name string
		args []arg
	}{
		{name: "NewManhattanLine", args: args},
	}
	for _, bm := range benchmarks {
		b.Run(
			bm.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, a := range bm.args {
						_ = NewManhattanLine(a.x0, a.y0, a.x1, a.y1)
					}
				}
			},
		)
	}
}

func BenchmarkParseInput(b *testing.B) {
	rdr, err := readerFromFileContents("input")
	if err != nil {
		b.Errorf("Bench the fuck out of here motherbucker! %s", err)
	}
	benchmarks := []struct {
		name string
	}{
		{name: "ParseInput"},
	}
	for _, bm := range benchmarks {
		b.Run(
			bm.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = parseInput(rdr)
				}
			},
		)
	}
}

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
		canvas *Canvas
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
				canvas: newCanvas(5),
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
				canvas: &Canvas{
					[][]int{
						{0, 0, 0, 0, 0},
						{1, 0, 0, 0, 0},
						{1, 0, 0, 0, 0},
						{1, 0, 0, 0, 0},
						{0, 0, 0, 0, 0},
					},
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
				canvas: &Canvas{
					[][]int{
						{0, 0, 0, 0, 0},
						{2, 1, 1, 1, 0},
						{1, 0, 0, 0, 0},
						{1, 0, 0, 0, 0},
						{0, 0, 0, 0, 0},
					},
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
				canvas: &Canvas{
					[][]int{
						{0, 0, 0, 0, 1},
						{2, 1, 1, 1, 1},
						{1, 0, 0, 0, 1},
						{1, 0, 0, 0, 1},
						{0, 0, 0, 0, 1},
					},
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
				canvas: &Canvas{
					[][]int{
						{0, 1, 1, 1, 0},
						{0, 0, 0, 0, 0},
						{0, 0, 0, 0, 0},
						{0, 0, 0, 0, 0},
						{0, 0, 0, 0, 0},
					},
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
				canvas: &Canvas{
					[][]int{
						{0, 1, 1, 1, 0},
						{0, 0, 0, 0, 0},
						{0, 0, 0, 0, 0},
						{0, 0, 0, 0, 0},
						{0, 0, 0, 0, 0},
					},
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
				cv := tt.args.canvas
				err := cv.drawLine(tt.args.l)
				got := cv._c
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
		canvas *Canvas
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: testing.CoverMode(),
			args: args{
				canvas: &Canvas{
					[][]int{
						{0, 2, 1, 1, 0},
						{0, 1, 0, 0, 0},
						{0, 1, 0, 0, 0},
						{0, 1, 0, 0, 0},
						{1, 1, 1, 1, 1},
					},
				},
			},
			want: 1,
		}, {
			name: testing.CoverMode(),
			args: args{
				canvas: &Canvas{
					[][]int{
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
			},
			want: 5,
		}, {
			name: testing.CoverMode(),
			args: args{
				canvas: &Canvas{
					_c: [][]int{
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
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				cv := tt.args.canvas
				if got := cv.countIntersections(); got != tt.want {
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
				ish:    false,
				Points: []image.Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}},
			},
		}, {
			name: testing.CoverMode(), args: args{
				x0: 0,
				y0: 0,
				x1: 2,
				y1: 2,
			},
			want: ManhattanLine{
				ish:    true,
				Points: []image.Point{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}},
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

func Test_drawManhattanLines(t *testing.T) {
	type args struct {
		manhattanLines []ManhattanLine
		canvas         *Canvas
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
				canvas: newCanvas(10),
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
				cv := tt.args.canvas
				err := cv.drawManhattanLines(tt.args.manhattanLines)
				got := cv._c
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

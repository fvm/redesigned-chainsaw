package day05

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"math"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ManhattanLine struct {
	ish    bool          // Is the line Manhattan'ish, a.k.a. diagonal
	Points []image.Point `json:"points,omitempty" yaml:"points"`
}

func NewManhattanLine(x0, y0, x1, y1 int) ManhattanLine {
	var xs, ys []int
	var ish bool
	if x0 == x1 && y0 == y1 { // Point
		xs = []int{x0}
		ys = []int{y0}
		ish = false
	} else if x0 == x1 && y0 != y1 { // Vertical
		ys = interpolate(y0, y1)
		xs = fill(x0, len(ys))
		ish = false
	} else if x0 != x1 && y0 == y1 { // Horizontal
		xs = interpolate(x0, x1)
		ys = fill(y0, len(xs))
		ish = false
	} else if math.Abs(float64(y1-y0)/float64(x1-x0)) == 1 { // Diagonal
		xs = interpolate(x0, x1)
		ys = interpolate(y0, y1)
		ish = true
	}

	// Now get all the points
	ps := make([]image.Point, len(xs), len(xs))
	for i := range ps {
		ps[i] = image.Point{
			X: xs[i],
			Y: ys[i],
		}
	}
	return ManhattanLine{
		ish:    ish,
		Points: ps,
	}
}

type Canvas struct {
	_c [][]int
}

func (canvas *Canvas) drawManhattanLines(manhattanLines []ManhattanLine) error {
	var err error
	for _, l := range manhattanLines {
		if l.ish {
			continue
		}
		// drawLine each line on the _c
		err = canvas.drawLine(l)
		if err != nil {
			return errors.Wrap(err, "Error while drawing manhattanLines on _c")
		}
	}
	return nil
}

func (canvas *Canvas) drawLine(l ManhattanLine) error {
	for _, p := range l.Points {
		if p.X > len(canvas._c) {
			return errors.Errorf("Pixel X value %d out of _c bounds", p.X)
		}
		if p.Y > len(canvas._c[p.X]) {
			return errors.Errorf("Pixel Y value %d out of _c bounds", p.Y)
		}
		canvas._c[p.Y][p.X]++
	}

	return nil
}

func (canvas *Canvas) countIntersections() int {
	count := 0
	for x := range canvas._c {
		for y := range canvas._c[x] {
			if canvas._c[x][y] > 1 {
				count++
			}
		}
	}
	return count
}

func (canvas *Canvas) drawManhattanishLines(manhattanLines []ManhattanLine) error {
	var err error
	for _, l := range manhattanLines {
		// drawLine each line on the _c
		err = canvas.drawLine(l)
		if err != nil {
			return errors.Wrap(err, "Error while drawing manhattanLines on _c")
		}
	}
	return nil
}

func Solve() error {

	manhattanLines, err := readAndParseInputFile("day05/input")
	if err != nil {
		return err
	}

	solutionPartOne, err := solvePartOne(manhattanLines)
	if err != nil {
		return err
	}

	zap.L().Info(
		"Solution",
		zap.Int("Day", 5),
		zap.Int("Part", 1),
		zap.Int("Solution (increments)", solutionPartOne),
	)

	solutionPartTwo, err := solvePartTwo(manhattanLines)
	if err != nil {
		return err
	}

	zap.L().Info(
		"Solution",
		zap.Int("Day", 5),
		zap.Int("Part", 2),
		zap.Int("Solution (increments)", solutionPartTwo),
	)

	return nil
}

func solvePartOne(manhattanLines []ManhattanLine) (int, error) {
	var err error
	canvas := newCanvas(1000)
	err = canvas.drawManhattanLines(manhattanLines)
	if err != nil {
		return 0, err
	}
	numIntersections := canvas.countIntersections()
	return numIntersections, nil
}

func solvePartTwo(manhattanLines []ManhattanLine) (int, error) {
	var err error
	canvas := newCanvas(1000)
	err = canvas.drawManhattanishLines(manhattanLines)
	if err != nil {
		return 0, err
	}
	numIntersections := canvas.countIntersections()
	return numIntersections, nil
}

func newCanvas(n int) *Canvas {
	_c := make([][]int, n)
	rows := make([]int, n*n)
	for i := 0; i < n; i++ {
		_c[i] = rows[i*n : (i+1)*n]
	}

	return &Canvas{_c: _c}
}

func interpolate(min, max int) []int {
	if min == max {
		return []int{min}
	}
	// We need to preserve order, but loop in order
	flip := false
	var sta, sto int
	if min > max {
		sto = min
		sta = max
		flip = true
	} else {
		sta = min
		sto = max
	}
	rng := make([]int, sto-sta+1, sto-sta+1)
	for i := range rng {
		rng[i] = sta + i
	}
	if flip { // Flip it and reverse it
		sort.Sort(sort.Reverse(sort.IntSlice(rng)))
	}
	return rng
}
func fill(value int, length int) []int {
	vs := make([]int, length)
	for i := range vs {
		vs[i] = value
	}
	return vs
}

func readAndParseInputFile(fname string) ([]ManhattanLine, error) {
	rc, err := readerFromFileContents(fname)
	if err != nil {
		return nil, err
	}
	return parseInput(rc)
}

func parseInput(input io.Reader) ([]ManhattanLine, error) {

	lineScanner := bufio.NewScanner(input)

	lineScanner.Split(bufio.ScanLines)
	var lines []ManhattanLine
	var x0, y0, x1, y1 int
	for lineScanner.Scan() {
		inputLine := lineScanner.Text()
		// Parse the ManhattanLine into points and a ManhattanLine
		_, err := fmt.Fscanf(strings.NewReader(inputLine), "%d,%d -> %d,%d", &x0, &y0, &x1, &y1)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Error during scanning of inputLine '%s'", inputLine))
		}
		// We know that the lines (for now) are horizontal or vertical
		lines = append(lines, NewManhattanLine(x0, y0, x1, y1))
	}

	return lines, nil
}

func isManhattanish(x0, y0, x1, y1 int) bool {
	if x0 == x1 && y0 == y1 {
		return true
	}
	if x0 == x1 && y0 != y1 { // Vertical
		return true
	}
	if x0 != x1 && y0 == y1 { // Horizontal
		return true
	}
	if math.Abs(float64(y1-y0)/float64(x1-x0)) == 1 { // Diagonal
		return true
	}
	return false
}

func readerFromFileContents(fname string) (io.Reader, error) {
	contents, err := ioutil.ReadFile(fname)
	if err != nil {

	}
	return strings.NewReader(string(contents)), nil
}

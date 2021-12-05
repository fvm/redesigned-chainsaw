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
	ish     bool
	Points  []image.Point `json:"points,omitempty" yaml:"points"`
	XValues []int         `json:"x_values,omitempty" yaml:"x_values"` // Easy access while we're dealing with horizontal and vertical lines
	YValues []int         `json:"y_values,omitempty" yaml:"y_values"` // see above...
}

func NewManhattanLine(x0, y0, x1, y1 int) ManhattanLine {
	var xs, ys []int
	var ish bool
	if x0 == x1 && y0 == y1 { // Point
		xs = []int{x0}
		ys = []int{y0}
		ish = false
	} else if math.Abs(float64(y1-y0)/float64(x1-x0)) == 1 { // Diagonal
		xs = interpolate(x0, x1)
		ys = interpolate(y0, y1)
		ish = true
	} else if y0 != y1 { // Vertical
		ys = interpolate(y0, y1)
		xs = fill(x0, len(ys))
		ish = false
	} else if x0 != x1 { // Horizontal
		xs = interpolate(x0, x1)
		ys = fill(y0, len(xs))
		ish = false
	}

	// Now get all the points
	ps := make([]image.Point, len(xs))
	for i := range ps {
		ps[i] = image.Point{
			X: xs[i],
			Y: ys[i],
		}
	}
	return ManhattanLine{
		ish:     ish,
		Points:  ps,
		XValues: xs,
		YValues: ys,
	}
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
	canvas := createCanvas(1000, 1000)
	canvas, err = drawManhattanLines(manhattanLines, canvas)
	if err != nil {
		return 0, err
	}
	numIntersections := countIntersections(canvas)
	return numIntersections, nil
}

func solvePartTwo(manhattanLines []ManhattanLine) (int, error) {
	var err error
	canvas := createCanvas(1000, 1000)
	canvas, err = drawManhattanishLines(manhattanLines, canvas)
	if err != nil {
		return 0, err
	}
	numIntersections := countIntersections(canvas)
	return numIntersections, nil
}

func drawManhattanishLines(manhattanLines []ManhattanLine, canvas [][]int) ([][]int, error) {
	var err error
	for _, l := range manhattanLines {
		// drawLine each line on the canvas
		canvas, err = drawLine(l, canvas)
		if err != nil {
			return nil, errors.Wrap(err, "Error while drawing manhattanLines on canvas")
		}
	}
	return canvas, nil
}

func drawManhattanLines(manhattanLines []ManhattanLine, canvas [][]int) ([][]int, error) {
	var err error
	for _, l := range manhattanLines {
		if l.ish {
			continue
		}
		// drawLine each line on the canvas
		canvas, err = drawLine(l, canvas)
		if err != nil {
			return nil, errors.Wrap(err, "Error while drawing manhattanLines on canvas")
		}
	}
	return canvas, nil
}

func drawLine(l ManhattanLine, canvas [][]int) ([][]int, error) {
	for _, p := range l.Points {
		if p.X > len(canvas) {
			return nil, errors.Errorf("Pixel X value %d out of canvas bounds", p.X)
		}
		if p.Y > len(canvas[p.X]) {
			return nil, errors.Errorf("Pixel Y value %d out of canvas bounds", p.Y)
		}
		canvas[p.Y][p.X]++
	}
	return canvas, nil
}

func countIntersections(canvas [][]int) int {
	count := 0
	for x := range canvas {
		for y := range canvas[x] {
			if canvas[x][y] > 1 {
				count++
			}
		}
	}
	return count
}

func createCanvas(n, m int) [][]int {
	canvas := make([][]int, n)
	rows := make([]int, n*m)
	for i := 0; i < n; i++ {
		canvas[i] = rows[i*m : (i+1)*m]
	}
	return canvas
}

func interpolate(min, max int) []int {
	if min == max {
		return []int{min}
	}
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
	axis := make([]int, sto-sta+1)
	for i := range axis {
		axis[i] = sta + i
	}
	if flip {
		sort.Sort(sort.Reverse(sort.IntSlice(axis)))
	}
	return axis
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
	for lineScanner.Scan() {
		inputLine := lineScanner.Text()
		var x0, y0, x1, y1 int
		// Parse the ManhattanLine into points and a ManhattanLine
		_, err := fmt.Fscanf(strings.NewReader(inputLine), "%d,%d -> %d,%d", &x0, &y0, &x1, &y1)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Error during scanning of inputLine '%s'", inputLine))
		}
		// We know that the lines (for now) are horizontal or vertical
		if isManhattanish(x0, y0, x1, y1) {
			lines = append(
				lines, NewManhattanLine(x0, y0, x1, y1),
			)
		}
	}
	return lines, nil
}

func isManhattanish(x0, y0, x1, y1 int) bool {
	if x0 == x1 && y0 == y1 {
		return true
	}
	if math.Abs(float64(y1-y0)/float64(x1-x0)) == 1 { // Diagonal
		return true
	}
	if x0 == x1 && y0 != y1 { // Vertical
		return true
	}
	if x0 != x1 && y0 == y1 { // Horizontal
		return true
	}
	return false
}

func fill(value int, length int) []int {
	vs := make([]int, length)
	for i := range vs {
		vs[i] = value
	}
	return vs
}
func readerFromFileContents(fname string) (io.Reader, error) {
	contents, err := ioutil.ReadFile(fname)
	if err != nil {

	}
	return strings.NewReader(string(contents)), nil
}

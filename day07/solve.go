package day07

import (
	"encoding/csv"
	"io"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

type PuzzleInput struct {
	positions []int
}

type PuzzleOutput struct {
	fuel     int
	position int
}

type DistanceFunc func(int, []int) (int, error)

func getRange(s []int) (int, int) {
	currentMaximum := math.MinInt
	currentMinimum := math.MaxInt
	for _, v := range s {
		if v <= currentMinimum {
			currentMinimum = v
		}
		if v >= currentMaximum {
			currentMaximum = v
		}
	}
	return currentMinimum, currentMaximum
}

func solvePartOne(input PuzzleInput) (int, error) {
	fuel, _, err := leastDistance(input.positions, linear)
	if err != nil {
		return 0, err
	}
	return fuel, err
}

func solvePartTwo(input PuzzleInput) (int, error) {
	fuel, _, err := leastDistance(input.positions, cumulative)
	if err != nil {
		return 0, err
	}
	return fuel, err
}

// leastDistance gets the value for a range of inputs which minimises the distance using a distance function
func leastDistance(values []int, f DistanceFunc) (int, int, error) {
	minDistance := math.MaxInt
	centerValue := math.MinInt
	lo, hi := getRange(values)
	for i := lo; i <= hi; i++ {
		d, err := f(i, values)
		if err != nil {
			return 0, 0, err
		}
		if d <= minDistance {
			minDistance = d
			centerValue = i
		}
	}
	return minDistance, centerValue, nil
}

// linear gets the total distance from a proposed center to all values in a range using a linear sum
func linear(y int, values []int) (int, error) {
	sumOfDistances := 0
	for _, position := range values {
		// Linear sum, just add
		sumOfDistances += int(math.Abs(float64(position - y)))
	}
	return sumOfDistances, nil
}

func cumulative(y int, values []int) (int, error) {
	sumOfDistances := 0
	for _, position := range values {
		// Cumulative sum, 0+1+2+3+4+...+baseDistance
		baseDistance := int(math.Abs(float64(position - y)))
		for i := 0; i <= baseDistance; i++ {
			sumOfDistances += i
		}
	}
	return sumOfDistances, nil
}

func parseInput(input io.Reader) (PuzzleInput, error) {
	r := csv.NewReader(input)

	record, err := r.Read()
	if err != nil {
		return PuzzleInput{}, errors.Wrap(err, "CSV Bitches")
	}
	p := PuzzleInput{positions: make([]int, len(record))} // We know it's just one line
	for i, f := range record {
		val, err := strconv.Atoi(f)
		if err != nil {
			return p, errors.Wrap(err, "Beep! Boop! Your crabby string was very crappy")
		}
		p.positions[i] = val
	}

	return p, nil
}

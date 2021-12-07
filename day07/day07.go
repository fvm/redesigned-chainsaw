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
	// It would be smart if we did a gradient descent, but let's just brute force it
	type PuzzleOutput struct {
		fuel     int
		position int
	}
	puzzleOutput := PuzzleOutput{
		fuel:     math.MaxInt,
		position: math.MaxInt,
	}
	lo, hi := getRange(input.positions)
	for i := lo; i <= hi; i++ {
		f, err := input.linearGoodNessOfFit(i)
		if err != nil {
			return 0, err
		}
		if f <= puzzleOutput.fuel {
			puzzleOutput.position = i
			puzzleOutput.fuel = f
		}

	}

	return puzzleOutput.fuel, nil
}

func solvePartTwo(input PuzzleInput) (int, error) {
	fuel := math.MaxInt
	lo, hi := getRange(input.positions)
	for i := lo; i <= hi; i++ {
		f, err := input.cumulativeGoodnessOfFit(i)
		if err != nil {
			return 0, err
		}
		if f <= fuel {
			fuel = f
		}
	}
	return fuel, nil
}

func (p PuzzleInput) linearGoodNessOfFit(y int) (int, error) {
	sumOfDistances := 0
	for _, position := range p.positions {
		sumOfDistances += int(math.Abs(float64(position - y)))
	}
	return sumOfDistances, nil
}

func (p PuzzleInput) cumulativeGoodnessOfFit(y int) (int, error) {
	sumOfDistances := 0
	for _, position := range p.positions {
		baseDistance := int(math.Abs(float64(position - y)))
		for i := 1; i <= baseDistance; i++ {
			sumOfDistances += i
		}
	}
	return sumOfDistances, nil
}

func parseInput(input io.Reader) (PuzzleInput, error) {
	r := csv.NewReader(input)

	record, err := r.Read()
	if err != nil {
		return PuzzleInput{}, err
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

package day01

import (
	"bufio"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Solve() error {
	data, err := ReadInput("day01/input")
	if err != nil {
		return err
	}

	solutionPartOne, err := solvePartOne(data)
	if err != nil {
		return err
	}

	zap.L().Info(
		"Solution",
		zap.Int("Day", 1),
		zap.Int("Part", 1),
		zap.Int("Solution (increments)", solutionPartOne),
	)

	solutionPartTwo, err := solvePartTwo(data)
	if err != nil {
		return err
	}

	zap.L().Info(
		"Solution",
		zap.Int("Day", 1),
		zap.Int("Part", 2),
		zap.Int("Solution (increments)", solutionPartTwo),
	)

	return nil
}

func CountSingularIncrements(values []int) (int, error) {
	if len(values) <= 1 {
		return 0, nil
	}

	increments := 0

	for i := len(values) - 1; i > 0; i-- {
		if values[i] > values[i-1] {
			increments++
		}
	}

	return increments, nil
}

func solvePartOne(input []int) (int, error) {
	return CountSingularIncrements(input)
}
func solvePartTwo(input []int) (int, error) {
	return CountSubsliceIncrements(input, 3)
}

func CountSubsliceIncrements(input []int, windowsize int) (int, error) {
	var increments = 0

	if len(input) < windowsize {
		return 0, errors.Errorf("Window size (%d) exceeds input length (%d)", windowsize, len(input))
	}

	for i := len(input) - windowsize; i > 0; i-- {
		offset := i + windowsize
		sub01 := input[i:offset]
		sub02 := input[i-1 : offset-1]

		if intsSum(sub01) > intsSum(sub02) {
			increments++
		}
	}

	return increments, nil
}

func intsSum(values []int) (n int) {
	for _, v := range values {
		n += v
	}

	return n
}
func ReadInput(fname string) ([]int, error) {
	fptr, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(fptr)
	s.Split(bufio.ScanLines)

	var out []int

	for s.Scan() {
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}

		out = append(out, val)
	}

	return out, nil
}

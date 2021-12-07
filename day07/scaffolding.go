package day07

import (
	"io"
	"io/ioutil"
	"strings"

	"go.uber.org/zap"
)

func Solve() error {

	puzzleInput, err := readAndParseInputFile("day07/input")
	if err != nil {
		return err
	}

	solutionPartOne, err := solvePartOne(puzzleInput)
	if err != nil {
		return err
	}

	zap.L().Info(
		"Solution",
		zap.Int("Day", 7),
		zap.Int("Part", 1),
		zap.Int("Solution (increments)", solutionPartOne),
	)

	solutionPartTwo, err := solvePartTwo(puzzleInput)
	if err != nil {
		return err
	}

	zap.L().Info(
		"Solution",
		zap.Int("Day", 7),
		zap.Int("Part", 2),
		zap.Int("Solution (increments)", solutionPartTwo),
	)

	return nil
}
func readAndParseInputFile(fname string) (PuzzleInput, error) {
	rc, err := readerFromFileContents(fname)
	if err != nil {
		return PuzzleInput{}, err
	}
	return parseInput(rc)
}
func readerFromFileContents(fname string) (io.Reader, error) {
	contents, err := ioutil.ReadFile(fname)
	if err != nil {

	}
	return strings.NewReader(string(contents)), nil
}

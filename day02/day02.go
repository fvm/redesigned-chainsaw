package day02

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type displacement struct {
	direction string
	distance  int
}

func Solve() error {

	data, err := readInput("day02/input")
	if err != nil {
		return err
	}
	solutionPartOne, err := solvePartOne(data)
	if err != nil {
		return err
	}
	zap.L().Info(
		"Solution",
		zap.Int("Day", 2),
		zap.Int("Part", 1),
		zap.Int("Solution (increments)", solutionPartOne),
	)
	solutionPartTwo, err := solvePartTwo(data)
	if err != nil {
		return err
	}
	zap.L().Info(
		"Solution",
		zap.Int("Day", 2),
		zap.Int("Part", 2),
		zap.Int("Solution (increments)", solutionPartTwo),
	)
	return nil
}

func solvePartOne(displacements []displacement) (int, error) {
	var h, v int
	for _, d := range displacements {
		switch d.direction {
		case "up":
			h -= d.distance
		case "down":
			h += d.distance
		case "forward":
			v += d.distance
		default:
			return 0, errors.Errorf("WTF you went %s!", d.direction)
		}
	}
	return h * v, nil
}

func solvePartTwo(displacements []displacement) (int, error) {
	var h, v, a int
	for _, d := range displacements {
		switch d.direction {
		case "up":
			a -= d.distance
		case "down":
			a += d.distance
		case "forward":
			v += d.distance
			h += d.distance * a
		default:
			return 0, errors.Errorf("WTF you went %s!", d.direction)
		}
	}
	return h * v, nil
}

func readInput(fname string) ([]displacement, error) {
	fptr, err := os.Open(fname)
	if err != nil {
		return []displacement{}, err
	}
	defer func(fptr *os.File) {
		err := fptr.Close()
		if err != nil {
			zap.L().Error("Error closing file", zap.String("name", fptr.Name()), zap.Error(err))
		}
	}(fptr)
	// We know the file is 1000 lines
	ds := make([]displacement, 1000)
	s := bufio.NewScanner(fptr)
	s.Split(bufio.ScanLines)
	var i uint16
	for s.Scan() {
		n, err := fmt.Sscanf(s.Text(), "%s\t%d\n", &ds[i].direction, &ds[i].distance)
		if err != nil {
			return ds, err
		}
		if n != 2 {
			return ds, errors.Errorf("Somehow read %d inputs while scanning an input line... Qeu le fucque?", n)
		}
		i++
	}
	return ds, nil
}
